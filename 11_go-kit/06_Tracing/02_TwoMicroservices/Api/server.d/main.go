package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	api "github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Api/service"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"golang.org/x/time/rate"
)

const (
	defaultHTTPPort = "44513"
	//defaultGRPCPort = "44523"
	defaultMqttUrl            = "localhost"
	defaultMongoUrl           = "localhost"
	defaultIotMicroserviceUrl = "localhost:44413"
)

func init() {
	// PARALELISM ----------------------
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	errs := make(chan error, 2)

	var (
		httpaddr = envString("HTTP_PORT", defaultHTTPPort)
		//rpcpaddr = envString("GRPC_PORT", defaultGRPCPort)
		mqtturl            = envString("MQTT_URL", defaultMqttUrl)
		mongourl           = envString("MONGO_URL", defaultMongoUrl)
		iotmicroserviceurl = envString("CELLAR_IOT_URL", defaultIotMicroserviceUrl)

		httpAddr = flag.String("http.addr", ":"+httpaddr, "HTTP listen address")
		//gRPCAddr = flag.String("grpc.addr", ":"+rpcpaddr, "gRPC listen address")

		//ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	fieldKeys := []string{"method"}

	//TRACING ----------------------------------

	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New("officeapi", config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	stdopentracing.SetGlobalTracer(tracer)

	//-----------------------------------------
	// gRPC server
	//-----------------------------------------

	// // init lorem service
	// var bs0 iot.Service
	// bs0 = iot.NewService(mongourl, mqtturl)

	// // creating Endpoints struct
	// endpoints1 := iot.Endpoints{
	// 	GetAllSpacesEndpoint: iot.MakeGetAllSpacesEndpoint(bs0),
	// 	GetSenzorEndpoint:    iot.MakeGetSenzorEndpoint(bs0),
	// }

	// //execute grpc server
	// go func() {
	// 	listener, err := net.Listen("tcp", *gRPCAddr)
	// 	if err != nil {
	// 		errs <- err
	// 		return
	// 	}
	// 	handler := iot.MakeGrpcHandler(ctx, endpoints1)
	// 	gRPCServer := grpc.NewServer()
	// 	pb.RegisterIoTServiceServer(gRPCServer, handler)
	// 	errs <- gRPCServer.Serve(listener)
	// }()

	//-----------------------------------------
	// HTTP server
	//-----------------------------------------

	// API --------------------------
	var bs1 api.Service
	bs1 = api.NewService(mongourl, mqtturl, iotmicroserviceurl)
	bs1 = api.NewLoggingMiddleware(log.With(logger, "component", "officeapi"), bs1)
	bs1 = api.NewMetricsMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "officeapi",
			Subsystem: "officeApiService",
			Name:      "http_request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "officeapi",
			Subsystem: "officeApiService",
			Name:      "http_request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		bs1,
	)

	var getAllSpacesEndpoint endpoint.Endpoint
	{
		getAllSpacesEndpoint = api.MakeGetAllSpacesEndpoint(bs1)
		getAllSpacesEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getAllSpacesEndpoint)
		getAllSpacesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getAllSpacesEndpoint)
		getAllSpacesEndpoint = opentracing.TraceServer(tracer, "GetAllSpaces")(getAllSpacesEndpoint)
		getAllSpacesEndpoint = LoggingMiddleware(log.With(logger, "method", "getAllSpaces"))(getAllSpacesEndpoint)
	}

	apiendpoints := api.Endpoints{
		GetAllSpacesEndpoint: getAllSpacesEndpoint,
	}

	// SERVER -------------------------------------------
	httpLogger := log.With(logger, "component", "http")

	// headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})

	// http.Handle("/", handlers.CORS(headersOk, originsOk, methodsOk)(api.MakeHttpHandler(ctx, apiendpoints, httpLogger)))
	// http.Handle("/metrics", promhttp.Handler())

	logger.Log("API IS RUNNING")

	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, api.MakeHttpHandler(apiendpoints, tracer, httpLogger))
	}()

	//-----------------------------------------
	// END SIGNAL
	//-----------------------------------------

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

//---------------------------------------------------------
//HELPERS -------------------------------------------------
//---------------------------------------------------------

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			defer func(begin time.Time) {
				logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())
			return next(ctx, request)

		}
	}
}

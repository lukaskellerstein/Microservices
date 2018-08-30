package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/gorilla/handlers"
	"github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Iot/pb"
	iot "github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Iot/service"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/gobreaker"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
)

const (
	defaultHTTPPort = "44403"
	defaultGRPCPort = "44413"
	defaultMqttUrl  = "localhost"
	defaultMongoUrl = "localhost"
)

func main() {
	errs := make(chan error, 2)

	var (
		httpaddr = envString("HTTP_PORT", defaultHTTPPort)
		rpcpaddr = envString("GRPC_PORT", defaultGRPCPort)
		mqtturl  = envString("MQTT_URL", defaultMqttUrl)
		mongourl = envString("MONGO_URL", defaultMongoUrl)

		httpAddr = flag.String("http.addr", ":"+httpaddr, "HTTP listen address")
		gRPCAddr = flag.String("grpc.addr", ":"+rpcpaddr, "gRPC listen address")

		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

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
	tracer, closer, err := cfg.New("iotmicroservice", config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	//-----------------------------------------
	// gRPC server
	//-----------------------------------------

	// init service
	var bs0 iot.Service
	bs0 = iot.NewService(mongourl, mqtturl)
	bs0 = iot.NewMetricsMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "iot",
			Subsystem: "iotService",
			Name:      "grpc_request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "iot",
			Subsystem: "iotService",
			Name:      "grpc_request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		bs0,
	)

	var getAllSpacesEndpoint endpoint.Endpoint
	{
		getAllSpacesEndpoint = iot.MakeGetAllSpacesEndpoint(bs0)
		getAllSpacesEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getAllSpacesEndpoint)
		getAllSpacesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getAllSpacesEndpoint)
		getAllSpacesEndpoint = opentracing.TraceServer(tracer, "GetAllSpaces")(getAllSpacesEndpoint)
		// getAllSpacesEndpoint = LoggingMiddleware(log.With(logger, "method", "getAllSpaces"))(getAllSpacesEndpoint)
	}

	var getSenzorEndpoint endpoint.Endpoint
	{
		getSenzorEndpoint = iot.MakeGetAllSpacesEndpoint(bs0)
		getSenzorEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(getSenzorEndpoint)
		getSenzorEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(getSenzorEndpoint)
		getSenzorEndpoint = opentracing.TraceServer(tracer, "GetSenzor")(getSenzorEndpoint)
		// getSenzorEndpoint = LoggingMiddleware(log.With(logger, "method", "GetSenzor"))(getSenzorEndpoint)
	}

	// creating Endpoints struct
	endpoints1 := iot.Endpoints{
		GetAllSpacesEndpoint: getAllSpacesEndpoint,
		GetSenzorEndpoint:    getSenzorEndpoint,
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errs <- err
			return
		}
		handler := iot.MakeGrpcHandler(ctx, endpoints1)
		gRPCServer := grpc.NewServer()
		pb.RegisterIoTServiceServer(gRPCServer, handler)
		errs <- gRPCServer.Serve(listener)
	}()

	//-----------------------------------------
	// HTTP server
	//-----------------------------------------

	// IOT --------------------------
	var bs1 iot.Service
	bs1 = iot.NewService(mongourl, mqtturl)
	bs1 = iot.NewLoggingMiddleware(log.With(logger, "component", "iot"), bs1)
	bs1 = iot.NewMetricsMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "iot",
			Subsystem: "iotService",
			Name:      "http_request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "iot",
			Subsystem: "iotService",
			Name:      "http_request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		bs1,
	)

	iotendpoints := iot.Endpoints{
		GetAllSpacesEndpoint:  iot.MakeGetAllSpacesEndpoint(bs1),
		GetRootSpacesEndpoint: iot.MakeGetRootSpacesEndpoint(bs1),
		GetSpacesEndpoint:     iot.MakeGetSpacesEndpoint(bs1),
		RemoveSpacesEndpoint:  iot.MakeRemoveSpacesEndpoint(bs1),
		GetSpaceEndpoint:      iot.MakeGetSpaceEndpoint(bs1),
		AddSpaceEndpoint:      iot.MakeAddSpaceEndpoint(bs1),
		RemoveSpaceEndpoint:   iot.MakeRemoveSpaceEndpoint(bs1),
		UpdateSpaceEndpoint:   iot.MakeUpdateSpaceEndpoint(bs1),
		GetAllSenzorsEndpoint: iot.MakeGetAllSenzorsEndpoint(bs1),
		GetSenzorsEndpoint:    iot.MakeGetSenzorsEndpoint(bs1),
		RemoveSenzorsEndpoint: iot.MakeRemoveSenzorsEndpoint(bs1),
		GetSenzorEndpoint:     iot.MakeGetSenzorEndpoint(bs1),
		AddSenzorEndpoint:     iot.MakeAddSenzorEndpoint(bs1),
		RemoveSenzorEndpoint:  iot.MakeRemoveSenzorEndpoint(bs1),
		UpdateSenzorEndpoint:  iot.MakeUpdateSenzorEndpoint(bs1),
		GetAllPlacesEndpoint:  iot.MakeGetAllPlacesEndpoint(bs1),
		GetPlaceEndpoint:      iot.MakeGetPlaceEndpoint(bs1),
		AddPlaceEndpoint:      iot.MakeAddPlaceEndpoint(bs1),
		RemovePlaceEndpoint:   iot.MakeRemovePlaceEndpoint(bs1),
		UpdatePlaceEndpoint:   iot.MakeUpdatePlaceEndpoint(bs1),
		PublishToMqttEndpoint: iot.MakePublishToMqttEndpoint(bs1),
	}

	// SERVER -------------------------------------------
	httpLogger := log.With(logger, "component", "http")

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})

	http.Handle("/iot/", handlers.CORS(headersOk, originsOk, methodsOk)(iot.MakeHttpHandler(ctx, iotendpoints, httpLogger)))
	// http.Handle("/mqtt/", handlers.CORS(headersOk, originsOk, methodsOk)(mqtt.MakeHttpHandler(ctx, mqttendpoints, httpLogger)))
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("API IS RUNNING")

	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
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

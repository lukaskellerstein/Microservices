package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/time/rate"

	calc "github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/calculator"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/sony/gobreaker"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//TRACING
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.New("calculator", config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()

	//SERVICE CREATE
	var svc calc.Service
	svc = calc.CalculatorService{}

	//ENDPOINTS CREATE
	var plusEndpoint endpoint.Endpoint
	{
		plusEndpoint = calc.MakePlusEndpoint(svc)
		plusEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(plusEndpoint)
		plusEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(plusEndpoint)
		plusEndpoint = opentracing.TraceServer(tracer, "Plus")(plusEndpoint)
		plusEndpoint = LoggingMiddleware(log.With(logger, "method", "Plus"))(plusEndpoint)
	}

	var minusEndpoint endpoint.Endpoint
	{
		minusEndpoint = calc.MakeMinusEndpoint(svc)
		minusEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(minusEndpoint)
		minusEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(minusEndpoint)
		minusEndpoint = opentracing.TraceServer(tracer, "Minus")(minusEndpoint)
		minusEndpoint = LoggingMiddleware(log.With(logger, "method", "Minus"))(minusEndpoint)
	}

	var multiplyEndpoint endpoint.Endpoint
	{
		multiplyEndpoint = calc.MakeMultiEndpoint(svc)
		multiplyEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(multiplyEndpoint)
		multiplyEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(multiplyEndpoint)
		multiplyEndpoint = opentracing.TraceServer(tracer, "Multiply")(multiplyEndpoint)
		multiplyEndpoint = LoggingMiddleware(log.With(logger, "method", "Multiply"))(multiplyEndpoint)
	}

	var divideEndpoint endpoint.Endpoint
	{
		divideEndpoint = calc.MakeDivideEndpoint(svc)
		divideEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(divideEndpoint)
		divideEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(divideEndpoint)
		divideEndpoint = opentracing.TraceServer(tracer, "Divide")(divideEndpoint)
		divideEndpoint = LoggingMiddleware(log.With(logger, "method", "Divide"))(divideEndpoint)
	}

	endpoint := calc.Endpoints{
		PlusEndpoint:   plusEndpoint,
		MinusEndpoint:  minusEndpoint,
		MultiEndpoint:  multiplyEndpoint,
		DivideEndpoint: divideEndpoint,
	}

	// GOROUTINE - HTTP transport
	r := calc.MakeHttpHandler(ctx, endpoint, logger)
	go func() {
		fmt.Println("Starting server at port 8080")
		handler := r
		errChan <- http.ListenAndServe(":8080", handler)
	}()

	// DAEMON, run until ERR or EXIT signal
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}

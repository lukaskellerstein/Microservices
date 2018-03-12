package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	calc "../calculator"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//declare metrics
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "calculator",
		Subsystem: "calculatorService",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "calculator",
		Subsystem: "calculatorService",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	//SERVICE CREATE
	var svc calc.Service
	svc = calc.Calculator{}
	// svc = calc.LoggingMiddleware(logger)(svc)
	svc = calc.Metrics(requestCount, requestLatency)(svc)

	// rlbucket := ratelimit.NewBucket(1*time.Second, 5)
	// e := calc.MakePlusEndpoint(svc)
	// e = ratelimitkit.NewTokenBucketThrottler(rlbucket, time.Sleep)(e)
	endpoint := calc.Endpoints{
		PlusEndpoint:   calc.MakePlusEndpoint(svc),
		MinusEndpoint:  calc.MakeMinusEndpoint(svc),
		MultiEndpoint:  calc.MakeMultiplyEndpoint(svc),
		DivideEndpoint: calc.MakeDivideEndpoint(svc),
	}
	
	// GOROUTINE - HTTP transport
	r := calc.MakeHttpHandler(ctx, endpoint)

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

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	calc "../calculator"
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	//SERVICE CREATE
	var svc calc.Service
	svc = calc.CalculatorService{}
	endpoint := calc.Endpoints{
		PlusEndpoint:   calc.MakePlusEndpoint(svc),
		MinusEndpoint:  calc.MakeMinusEndpoint(svc),
		MultiEndpoint:  calc.MakeMultiEndpoint(svc),
		DivideEndpoint: calc.MakeDivideEndpoint(svc),
	}

	// Logging domain.
	// var logger log.Logger
	// {
	// 	logger = log.NewLogfmtLogger(os.Stderr)
	// 	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	// 	logger = log.With(logger, "caller", log.DefaultCaller)
	// }

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

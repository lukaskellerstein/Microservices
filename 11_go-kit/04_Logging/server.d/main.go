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
)

func main() {
	ctx := context.Background()
	errChan := make(chan error)

	logfile, err := os.OpenFile("./microservice.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()

	// Logging domain.
	var logger log.Logger
	{
		//File
		w := log.NewSyncWriter(logfile)
		logger = log.NewLogfmtLogger(w)

		//OR

		//Console
		// w := log.NewSyncWriter(os.Stderr)
		// logger := log.NewLogfmtLogger(w)

		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//SERVICE CREATE
	var svc calc.Service
	svc = calc.Calculator{}
	svc = calc.LoggingMiddleware(logger)(svc)
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

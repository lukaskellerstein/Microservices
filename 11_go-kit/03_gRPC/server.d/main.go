package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	calc "../calculator"
	pb "../pb"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var (
		gRPCAddr = flag.String("grpc", ":8081",
			"gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// init lorem service
	var svc calc.Service
	svc = calc.CalculatorService{}
	errChan := make(chan error)

	// creating Endpoints struct
	endpoints := calc.Endpoints{
		PlusEndpoint: calc.MakePlusEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := calc.MakeGrpcHandler(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterCalculatorServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println(<-errChan)
}

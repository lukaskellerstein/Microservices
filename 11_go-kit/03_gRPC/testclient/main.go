package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "../pb"
)

const (
	address = "localhost:8081"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func Plus(client pb.CalculatorClient, req *pb.PlusRequest) {
	resp, err := client.Plus(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	log.Println(resp.Result)
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCalculatorClient(conn)

	customer := &pb.PlusRequest{
		A: 10,
		B: 15,
	}

	// Create a new customer
	Plus(client, customer)

}

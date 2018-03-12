package calculator

import (
	pb "../pb"
	"golang.org/x/net/context"
)

//Encode and Decode Lorem Request
func EncodeGRPCPlusRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(PlusRequest)
	return &pb.PlusRequest{
		A: req.A,
		B: req.B,
	}, nil
}

func DecodeGRPCPlusRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.PlusRequest)
	return PlusRequest{
		A: req.A,
		B: req.B,
	}, nil
}

// Encode and Decode Lorem Response
func EncodeGRPCPlusResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(PlusResponse)
	return &pb.PlusResponse{
		Result: resp.Result,
	}, nil

}

func DecodeGRPCPlusResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.PlusResponse)
	return PlusResponse{
		Result: resp.Result,
	}, nil
}

package calculator

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/net/context"

	pb "../pb"
)

type grpcServer struct {
	plus     grpctransport.Handler
	minus    grpctransport.Handler
	multiple grpctransport.Handler
	divide   grpctransport.Handler
}

// implement LoremServer Interface in lorem.pb.go
func (s *grpcServer) Plus(ctx context.Context, r *pb.PlusRequest) (*pb.PlusResponse, error) {
	_, resp, err := s.plus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PlusResponse), nil
}

func MakeGrpcHandler(ctx context.Context, endpoint Endpoints) pb.CalculatorServer {
	return &grpcServer{
		plus: grpctransport.NewServer(
			endpoint.PlusEndpoint,
			DecodeGRPCPlusRequest,
			EncodeGRPCPlusResponse,
		),
	}
}

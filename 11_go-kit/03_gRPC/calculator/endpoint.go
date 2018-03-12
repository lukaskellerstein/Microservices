package calculator

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	PlusEndpoint   endpoint.Endpoint
	MinusEndpoint  endpoint.Endpoint
	MultiEndpoint  endpoint.Endpoint
	DivideEndpoint endpoint.Endpoint
}

type PlusRequest struct {
	A int32
	B int32
}

type PlusResponse struct {
	Result int32
}

func MakePlusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PlusRequest)

		//call service
		result2, err := svc.Plus(ctx, req.A, req.B)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(result2)

		return PlusResponse{Result: result2}, nil
	}
}

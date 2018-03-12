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

type CalculatorRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalculatorResponse struct {
	Result int `json:"result"`
}

func MakePlusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculatorRequest)

		//call service
		result2 := svc.Plus(req.A, req.B)
		fmt.Println(result2)

		return CalculatorResponse{Result: result2}, nil
	}
}

func MakeMinusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculatorRequest)

		//call service
		result2 := svc.Minus(req.A, req.B)
		fmt.Println(result2)

		return CalculatorResponse{Result: result2}, nil
	}
}

func MakeMultiplyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculatorRequest)

		//call service
		result2 := svc.Multiply(req.A, req.B)
		fmt.Println(result2)

		return CalculatorResponse{Result: result2}, nil
	}
}

func MakeDivideEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CalculatorRequest)

		//call service
		result2 := svc.Divide(req.A, req.B)
		fmt.Println(result2)

		return CalculatorResponse{Result: result2}, nil
	}
}

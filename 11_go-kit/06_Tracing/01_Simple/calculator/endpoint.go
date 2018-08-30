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

//*************************
// PLUS
//*************************

type plusRequest struct {
	a int
	b int
}

type PlusResponse struct {
	Result int `json: "result"`
}

func MakePlusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(plusRequest)

		//call service
		result2 := svc.plus(req.a, req.b)
		fmt.Println(result2)

		return PlusResponse{Result: result2}, nil
	}
}

//*************************
// MINUS
//*************************

type minusRequest struct {
	a int
	b int
}

type MinusResponse struct {
	Result int `json: "result"`
}

func MakeMinusEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(minusRequest)

		//call service
		result := svc.minus(req.a, req.b)

		return MinusResponse{Result: result}, nil
	}
}

//*************************
// MULTI
//*************************

type multiRequest struct {
	a int
	b int
}

type MultiResponse struct {
	Result int `json: "result"`
}

func MakeMultiEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(multiRequest)

		//call service
		result := svc.multiple(req.a, req.b)

		return MultiResponse{Result: result}, nil
	}
}

//*************************
// DIVIDE
//*************************

type divideRequest struct {
	a int
	b int
}

type DivideResponse struct {
	Result int `json: "result"`
}

func MakeDivideEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(divideRequest)

		//call service
		result := svc.divide(req.a, req.b)

		return DivideResponse{Result: result}, nil
	}
}

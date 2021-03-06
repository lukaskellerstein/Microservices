package calculator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdprometheus "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHttpHandler(ctx context.Context, endpoint Endpoints) http.Handler {

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		// httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/plus/{a}/{b}").Handler(httptransport.NewServer(
		endpoint.PlusEndpoint,
		decodeCalculatorRequest,
		encodeCalculatorResponse,
		options...,
	))

	r.Methods("POST").Path("/minus/{a}/{b}").Handler(httptransport.NewServer(
		endpoint.MinusEndpoint,
		decodeCalculatorRequest,
		encodeCalculatorResponse,
		options...,
	))

	r.Methods("POST").Path("/multi/{a}/{b}").Handler(httptransport.NewServer(
		endpoint.MultiEndpoint,
		decodeCalculatorRequest,
		encodeCalculatorResponse,
		options...,
	))

	r.Methods("POST").Path("/divide/{a}/{b}").Handler(httptransport.NewServer(
		endpoint.DivideEndpoint,
		decodeCalculatorRequest,
		encodeCalculatorResponse,
		options...,
	))

	// GET /metrics
	r.Path("/metrics").Handler(stdprometheus.Handler())

	return r
}

//*************************
// PLUS
//*************************

// decode url path variables into request
func decodeCalculatorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	a, ok := vars["a"]
	if !ok {
		return nil, ErrBadRouting
	}

	b, ok := vars["b"]
	if !ok {
		return nil, ErrBadRouting
	}

	aint, _ := strconv.Atoi(a)
	bint, _ := strconv.Atoi(b)
	return CalculatorRequest{
		A: aint,
		B: bint,
	}, nil
}

// encode response from endpoint
func encodeCalculatorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}

	// fmt.Println(ctx)
	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

//*************************
// ERROR
//*************************

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encode error
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

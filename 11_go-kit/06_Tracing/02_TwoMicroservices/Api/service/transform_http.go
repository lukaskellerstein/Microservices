package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHttpHandler(endpoint Endpoints, tracer stdopentracing.Tracer, logger kitlog.Logger) http.Handler {

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	//-----------------------------
	// SPACE
	//-----------------------------

	r.Path("/spaces").Handler(httptransport.NewServer(
		endpoint.GetAllSpacesEndpoint,
		decodeGetAllSpacesRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.HTTPToContext(tracer, "GetAllSpaces", logger)))...,
	))

	// r.Methods("GET").Path("/iot/getrootspaces").Handler(httptransport.NewServer(
	// 	endpoint.GetRootSpacesEndpoint,
	// 	decodeGetRootSpacesRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/getspaces").Handler(httptransport.NewServer(
	// 	endpoint.GetSpacesEndpoint,
	// 	decodeGetSpacesRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/removespaces").Handler(httptransport.NewServer(
	// 	endpoint.RemoveSpacesEndpoint,
	// 	decodeRemoveSpacesRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Path("/iot/getspace").Handler(httptransport.NewServer(
	// 	endpoint.GetSpaceEndpoint,
	// 	decodeGetSpaceRequest,
	// 	encodeGetSpaceResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/addspace").Handler(httptransport.NewServer(
	// 	endpoint.AddSpaceEndpoint,
	// 	decodeAddSpaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/removespace").Handler(httptransport.NewServer(
	// 	endpoint.RemoveSpaceEndpoint,
	// 	decodeRemoveSpaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/updatespace").Handler(httptransport.NewServer(
	// 	endpoint.UpdateSpaceEndpoint,
	// 	decodeUpdateSpaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// //-----------------------------
	// // SENZOR
	// //-----------------------------

	// r.Methods("GET").Path("/iot/getallsenzors").Handler(httptransport.NewServer(
	// 	endpoint.GetAllSenzorsEndpoint,
	// 	decodeGetAllSenzorsRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/getsenzors").Handler(httptransport.NewServer(
	// 	endpoint.GetSenzorsEndpoint,
	// 	decodeGetSenzorsRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/removesenzors").Handler(httptransport.NewServer(
	// 	endpoint.RemoveSenzorsEndpoint,
	// 	decodeRemoveSenzorsRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/getsenzor").Handler(httptransport.NewServer(
	// 	endpoint.GetSenzorEndpoint,
	// 	decodeGetSenzorRequest,
	// 	encodeGetSenzorResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/addsenzor").Handler(httptransport.NewServer(
	// 	endpoint.AddSenzorEndpoint,
	// 	decodeAddSenzorRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/removesenzor").Handler(httptransport.NewServer(
	// 	endpoint.RemoveSenzorEndpoint,
	// 	decodeRemoveSenzorRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/updatesenzor").Handler(httptransport.NewServer(
	// 	endpoint.UpdateSenzorEndpoint,
	// 	decodeUpdateSenzorRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// //-----------------------------
	// // PLACE
	// //-----------------------------

	// r.Methods("GET").Path("/iot/getallplaces").Handler(httptransport.NewServer(
	// 	endpoint.GetAllPlacesEndpoint,
	// 	decodeGetAllPlacesRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/getplace").Handler(httptransport.NewServer(
	// 	endpoint.GetPlaceEndpoint,
	// 	decodeGetPlaceRequest,
	// 	encodeGetPlaceResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/addplace").Handler(httptransport.NewServer(
	// 	endpoint.AddPlaceEndpoint,
	// 	decodeAddPlaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/removeplace").Handler(httptransport.NewServer(
	// 	endpoint.RemovePlaceEndpoint,
	// 	decodeRemovePlaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// r.Methods("POST").Path("/iot/updateplace").Handler(httptransport.NewServer(
	// 	endpoint.UpdatePlaceEndpoint,
	// 	decodeUpdatePlaceRequest,
	// 	encodeResponse,
	// 	options...,
	// ))

	// //-----------------------------
	// // MQTT
	// //-----------------------------

	// r.Path("/iot/publishtomqtt").Handler(httptransport.NewServer(
	// 	endpoint.PublishToMqttEndpoint,
	// 	decodePublishToMqttRequest,
	// 	encodePublishToMqttResponse,
	// 	options...,
	// ))

	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

//-----------------------------
// SPACE
//-----------------------------

func decodeGetAllSpacesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllSpacesRequest{}, nil
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

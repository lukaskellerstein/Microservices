package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

func MakeHttpHandler(ctx context.Context, endpoint Endpoints, logger kitlog.Logger) http.Handler {

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	//-----------------------------
	// SPACE
	//-----------------------------

	r.Path("/iot/getallspaces").Handler(httptransport.NewServer(
		endpoint.GetAllSpacesEndpoint,
		decodeGetAllSpacesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/iot/getrootspaces").Handler(httptransport.NewServer(
		endpoint.GetRootSpacesEndpoint,
		decodeGetRootSpacesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/getspaces").Handler(httptransport.NewServer(
		endpoint.GetSpacesEndpoint,
		decodeGetSpacesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/removespaces").Handler(httptransport.NewServer(
		endpoint.RemoveSpacesEndpoint,
		decodeRemoveSpacesRequest,
		encodeResponse,
		options...,
	))

	r.Path("/iot/getspace").Handler(httptransport.NewServer(
		endpoint.GetSpaceEndpoint,
		decodeGetSpaceRequest,
		encodeGetSpaceResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/addspace").Handler(httptransport.NewServer(
		endpoint.AddSpaceEndpoint,
		decodeAddSpaceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/removespace").Handler(httptransport.NewServer(
		endpoint.RemoveSpaceEndpoint,
		decodeRemoveSpaceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/updatespace").Handler(httptransport.NewServer(
		endpoint.UpdateSpaceEndpoint,
		decodeUpdateSpaceRequest,
		encodeResponse,
		options...,
	))

	//-----------------------------
	// SENZOR
	//-----------------------------

	r.Methods("GET").Path("/iot/getallsenzors").Handler(httptransport.NewServer(
		endpoint.GetAllSenzorsEndpoint,
		decodeGetAllSenzorsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/getsenzors").Handler(httptransport.NewServer(
		endpoint.GetSenzorsEndpoint,
		decodeGetSenzorsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/removesenzors").Handler(httptransport.NewServer(
		endpoint.RemoveSenzorsEndpoint,
		decodeRemoveSenzorsRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/getsenzor").Handler(httptransport.NewServer(
		endpoint.GetSenzorEndpoint,
		decodeGetSenzorRequest,
		encodeGetSenzorResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/addsenzor").Handler(httptransport.NewServer(
		endpoint.AddSenzorEndpoint,
		decodeAddSenzorRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/removesenzor").Handler(httptransport.NewServer(
		endpoint.RemoveSenzorEndpoint,
		decodeRemoveSenzorRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/updatesenzor").Handler(httptransport.NewServer(
		endpoint.UpdateSenzorEndpoint,
		decodeUpdateSenzorRequest,
		encodeResponse,
		options...,
	))

	//-----------------------------
	// PLACE
	//-----------------------------

	r.Methods("GET").Path("/iot/getallplaces").Handler(httptransport.NewServer(
		endpoint.GetAllPlacesEndpoint,
		decodeGetAllPlacesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/getplace").Handler(httptransport.NewServer(
		endpoint.GetPlaceEndpoint,
		decodeGetPlaceRequest,
		encodeGetPlaceResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/addplace").Handler(httptransport.NewServer(
		endpoint.AddPlaceEndpoint,
		decodeAddPlaceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/removeplace").Handler(httptransport.NewServer(
		endpoint.RemovePlaceEndpoint,
		decodeRemovePlaceRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/iot/updateplace").Handler(httptransport.NewServer(
		endpoint.UpdatePlaceEndpoint,
		decodeUpdatePlaceRequest,
		encodeResponse,
		options...,
	))

	//-----------------------------
	// MQTT
	//-----------------------------

	r.Path("/iot/publishtomqtt").Handler(httptransport.NewServer(
		endpoint.PublishToMqttEndpoint,
		decodePublishToMqttRequest,
		encodePublishToMqttResponse,
		options...,
	))

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

// func encodeGetAllSpacesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

// 	if e, ok := response.(errorer); ok && e.error() != nil {
// 		encodeError(ctx, e.error(), w)
// 		return nil
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(response)
// }

func decodeGetRootSpacesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetRootSpacesRequest{}, nil
}

// func encodeGetRootSpacesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

// 	if e, ok := response.(errorer); ok && e.error() != nil {
// 		encodeError(ctx, e.error(), w)
// 		return nil
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(response)
// }

func decodeGetSpacesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := GetSpacesRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

// func encodeGetSpacesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

// 	if e, ok := response.(errorer); ok && e.error() != nil {
// 		encodeError(ctx, e.error(), w)
// 		return nil
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	return json.NewEncoder(w).Encode(response)
// }

func decodeRemoveSpacesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := RemoveSpacesRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeGetSpaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := GetSpaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func encodeGetSpaceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	result := GetSpaceResponse{
		Data: response.(CellarSpace),
	}

	return json.NewEncoder(w).Encode(result)
}

func decodeAddSpaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := AddSpaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeRemoveSpaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := RemoveSpaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeUpdateSpaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := UpdateSpaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

//-----------------------------
// SENZOR
//-----------------------------

func decodeGetAllSenzorsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllSenzorsRequest{}, nil
}

func decodeGetSenzorsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := GetSenzorsRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeRemoveSenzorsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := RemoveSenzorsRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeGetSenzorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := GetSenzorRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func encodeGetSenzorResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	result := GetSenzorResponse{
		Data: response.(CellarSenzor),
	}

	return json.NewEncoder(w).Encode(result)
}

func decodeAddSenzorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := AddSenzorRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeRemoveSenzorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := RemoveSenzorRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeUpdateSenzorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := UpdateSenzorRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

//-----------------------------
// PLACE
//-----------------------------

func decodeGetAllPlacesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllPlacesRequest{}, nil
}

func decodeGetPlaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := GetPlaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func encodeGetPlaceResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	result := GetPlaceResponse{
		Data: response.(CellarPlace),
	}

	return json.NewEncoder(w).Encode(result)
}

func decodeAddPlaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := AddPlaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeRemovePlaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := RemovePlaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

func decodeUpdatePlaceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	body := UpdatePlaceRequest{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return body, nil
}

//*************************
// MQTT
//*************************

// decode url path variables into request
func decodePublishToMqttRequest(_ context.Context, r *http.Request) (interface{}, error) {

	//PARSING JSON
	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body) //<--- here!
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m, ok := gjson.Parse(string(htmlData)).Value().(map[string]interface{})
	if !ok {
		fmt.Println("Error")
	}

	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	item := &PublishToMqttRequest{}
	parseErr := json.Unmarshal(jsonBytes, item)
	if parseErr != nil {
		fmt.Println("JSON Error")
		fmt.Println(parseErr)
	}

	return *item, nil
}

// encode response from endpoint
func encodePublishToMqttResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}

	// fmt.Println(ctx)
	// fmt.Println(response)

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

package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	//Space
	GetAllSpacesEndpoint  endpoint.Endpoint
	GetRootSpacesEndpoint endpoint.Endpoint
	GetSpacesEndpoint     endpoint.Endpoint
	RemoveSpacesEndpoint  endpoint.Endpoint
	GetSpaceEndpoint      endpoint.Endpoint
	AddSpaceEndpoint      endpoint.Endpoint
	RemoveSpaceEndpoint   endpoint.Endpoint
	UpdateSpaceEndpoint   endpoint.Endpoint
	//Senzor
	GetAllSenzorsEndpoint endpoint.Endpoint
	GetSenzorsEndpoint    endpoint.Endpoint
	RemoveSenzorsEndpoint endpoint.Endpoint
	GetSenzorEndpoint     endpoint.Endpoint
	AddSenzorEndpoint     endpoint.Endpoint
	RemoveSenzorEndpoint  endpoint.Endpoint
	UpdateSenzorEndpoint  endpoint.Endpoint
	//Place
	GetAllPlacesEndpoint endpoint.Endpoint
	GetPlaceEndpoint     endpoint.Endpoint
	AddPlaceEndpoint     endpoint.Endpoint
	RemovePlaceEndpoint  endpoint.Endpoint
	UpdatePlaceEndpoint  endpoint.Endpoint
	//MQTT
	PublishToMqttEndpoint endpoint.Endpoint
}

//-----------------------------
// SPACE
//-----------------------------

type GetAllSpacesRequest struct{}

type GetAllSpacesResponse struct {
	Data []CellarSpace `json:"data"`
}

func MakeGetAllSpacesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//call service
		result, err := svc.GetAllSpaces()
		if err != nil {
			return nil, err
		}

		return GetAllSpacesResponse{Data: result}, nil
	}
}

type GetRootSpacesRequest struct{}

type GetRootSpacesResponse struct {
	Data []CellarSpace `json:"data"`
}

func MakeGetRootSpacesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//call service
		result, err := svc.GetRootSpaces()
		if err != nil {
			return nil, err
		}

		return GetRootSpacesResponse{Data: result}, nil
	}
}

type GetSpacesRequest struct {
	Path string `json:"path"`
}

type GetSpacesResponse struct {
	Data []CellarSpace `json:"data"`
}

func MakeGetSpacesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpacesRequest)

		//call service
		result, err := svc.GetSpaces(req.Path)
		if err != nil {
			return nil, err
		}

		return GetSpacesResponse{Data: result}, nil
	}
}

type RemoveSpacesRequest struct {
	Path string `json:"path"`
}

type RemoveSpacesResponse struct {
}

func MakeRemoveSpacesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveSpacesRequest)

		//call service
		err := svc.RemoveSpaces(req.Path)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type GetSpaceRequest struct {
	Id string `json:"id"`
}

type GetSpaceResponse struct {
	Data CellarSpace `json:"data"`
}

func MakeGetSpaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSpaceRequest)

		//call service
		result, err := svc.GetSpace(req.Id)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

type AddSpaceRequest struct {
	Item CellarSpace `json:"item"`
}

type AddSpaceResponse struct {
	Item CellarSpace `json:"item"`
}

func MakeAddSpaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddSpaceRequest)

		//call service
		item, err := svc.AddSpace(req.Item)
		if err != nil {
			return nil, err
		}

		return item, nil
	}
}

type RemoveSpaceRequest struct {
	Id string `json:"id"`
}

type RemoveSpaceResponse struct {
}

func MakeRemoveSpaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveSpaceRequest)

		//call service
		err := svc.RemoveSpace(req.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type UpdateSpaceRequest struct {
	Item CellarSpace `json:"item"`
}

type UpdateSpaceResponse struct {
	Item CellarSpace `json:"item"`
}

func MakeUpdateSpaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateSpaceRequest)

		//call service
		item, err := svc.UpdateSpace(req.Item)
		if err != nil {
			return nil, err
		}

		return item, nil
	}
}

//-----------------------------
// SENZOR
//-----------------------------

type GetAllSenzorsRequest struct{}

type GetAllSenzorsResponse struct {
	Data []CellarSenzor `json:"data"`
}

func MakeGetAllSenzorsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//call service
		result, err := svc.GetAllSenzors()
		if err != nil {
			return nil, err
		}

		//fmt.Println(result)

		return GetAllSenzorsResponse{Data: result}, nil
	}
}

type GetSenzorsRequest struct {
	Path string `json:"path"`
}

type GetSenzorsResponse struct {
	Data []CellarSenzor `json:"data"`
}

func MakeGetSenzorsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSenzorsRequest)

		//call service
		result, err := svc.GetSenzors(req.Path)
		if err != nil {
			return nil, err
		}

		return GetSenzorsResponse{Data: result}, nil
	}
}

type RemoveSenzorsRequest struct {
	Path string `json:"path"`
}

type RemoveSenzorsResponse struct {
}

func MakeRemoveSenzorsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveSenzorsRequest)

		//call service
		err := svc.RemoveSenzors(req.Path)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type GetSenzorRequest struct {
	Id string `json:"id"`
}

type GetSenzorResponse struct {
	Data CellarSenzor `json:"data"`
}

func MakeGetSenzorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetSenzorRequest)

		//call service
		result, err := svc.GetSenzor(req.Id)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

type AddSenzorRequest struct {
	Item CellarSenzor `json:"item"`
}

type AddSenzorResponse struct {
	Item CellarSenzor `json:"item"`
}

func MakeAddSenzorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddSenzorRequest)

		//call service
		item, err := svc.AddSenzor(req.Item)
		if err != nil {
			return nil, err
		}

		//fmt.Println(item)

		return item, nil
	}
}

type RemoveSenzorRequest struct {
	Id string `json:"id"`
}

type RemoveSenzorResponse struct {
}

func MakeRemoveSenzorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveSenzorRequest)

		//call service
		err := svc.RemoveSenzor(req.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type UpdateSenzorRequest struct {
	Item CellarSenzor `json:"item"`
}

type UpdateSenzorResponse struct {
	Item CellarSenzor `json:"item"`
}

func MakeUpdateSenzorEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateSenzorRequest)

		//call service
		item, err := svc.UpdateSenzor(req.Item)
		if err != nil {
			return nil, err
		}

		return item, nil
	}
}

//-----------------------------
// PLACE
//-----------------------------

type GetAllPlacesRequest struct{}

type GetAllPlacesResponse struct {
	Data []CellarPlace `json:"data"`
}

func MakeGetAllPlacesEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//call service
		result, err := svc.GetAllPlaces()
		if err != nil {
			return nil, err
		}

		return GetAllPlacesResponse{Data: result}, nil
	}
}

type GetPlaceRequest struct {
	Id string `json:"id"`
}

type GetPlaceResponse struct {
	Data CellarPlace `json:"data"`
}

func MakeGetPlaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPlaceRequest)

		//call service
		result, err := svc.GetPlace(req.Id)
		if err != nil {
			return nil, err
		}

		return result, nil
	}
}

type AddPlaceRequest struct {
	Item CellarPlace `json:"item"`
}

type AddPlaceResponse struct {
	Item CellarPlace `json:"item"`
}

func MakeAddPlaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddPlaceRequest)

		//call service
		item, err := svc.AddPlace(req.Item)
		if err != nil {
			return nil, err
		}

		return item, nil
	}
}

type RemovePlaceRequest struct {
	Id string `json:"id"`
}

type RemovePlaceResponse struct {
}

func MakeRemovePlaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemovePlaceRequest)

		//call service
		err := svc.RemovePlace(req.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

type UpdatePlaceRequest struct {
	Item CellarPlace `json:"item"`
}

type UpdatePlaceResponse struct {
	Item CellarPlace `json:"item"`
}

func MakeUpdatePlaceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePlaceRequest)

		//call service
		item, err := svc.UpdatePlace(req.Item)
		if err != nil {
			return nil, err
		}

		return item, nil
	}
}

//-----------------------------
// MQTT
//-----------------------------

type PublishToMqttRequest struct {
	Topic string `json:"topic"`
	Value string `json:"value"`
}

type PublishToMqttResponse struct {
	Result string `json:"result"`
}

func MakePublishToMqttEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PublishToMqttRequest)

		//call service
		result2 := svc.PublishToMqtt(req.Topic, req.Value)
		if result2 != nil {
			return PublishToMqttResponse{Result: "ERROR"}, result2
		}

		return PublishToMqttResponse{Result: "OK"}, nil
	}
}

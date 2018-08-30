package service

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	pb "github.com/lukaskellerstein/Microservices/11_go-kit/06_Tracing/02_TwoMicroservices/Api/pb/externall"
	"google.golang.org/grpc"
	mgo "gopkg.in/mgo.v2"
)

//-----------------------------
//-----------------------------
// SERVICE METHODS
//-----------------------------
//-----------------------------

type Service interface {
	//Space
	GetAllSpaces() ([]Space, error)
	GetSpaceInfo(id string) (SpaceInfo, error)
	GetSpaceTimeline(id string) ([]MeetingInfo, error)
	GetSpaceState(id string) (string, error)
	//Calendar
	GetSpaceCalendar(id string) ([]CalendarItem, error)
	GetDayInfo(spaceid string, year int, month int, day int) ([]MeetingInfo, error)
	GetMeetingInfo(meetingid string) (MeetingInfo, error)
	AddNewMeeting(item MeetingInfo) error
	UpdateMeeting(item MeetingInfo) (MeetingInfo, error)
	DeleteMeeting(id string) error
	//Reception
	CallForClean(spaceid string) error
	CallReception(spaceid string) error
	SomethingElse(spaceid string, text string) error
	GetSortiment(spaceid string) ([]CellarSortimentItem, error)
	PlaceOrder(spaceid string, item CellarOrder) error
	//User
	ValidatePin(pin string) (bool, error)
}

type OfficeApiService struct {
	MongoDbUrl         string
	Mqtturl            string
	IoTMicroserviceUrl string
}

var Database = "OfficeDatabase"
var session *mgo.Session

func NewService(mongodburl string, mqtturl string, iotmicroserviceurl string) *OfficeApiService {

	var err error
	if session, err = mgo.Dial(mongodburl); err != nil {
		log.Fatal(err)
	}

	return &OfficeApiService{
		MongoDbUrl:         mongodburl,
		Mqtturl:            mqtturl,
		IoTMicroserviceUrl: iotmicroserviceurl,
	}
}

//-----------------------------
// SPACE
//-----------------------------

func (s OfficeApiService) GetAllSpaces() (result []Space, err error) {
	// gRPC -----------------------------------
	conn, err := grpc.Dial(s.IoTMicroserviceUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewIoTServiceClient(conn)

	request := &pb.GetAllSpacesRequest{}

	allspacesResponse, err := client.GetAllSpaces(context.Background(), request)
	if err != nil {
		return nil, err
	}

	spaces := (*allspacesResponse).Data

	for _, item := range spaces {

		//fmt.Println(item)

		asdf := Space{
			ID:    bson.ObjectIdHex(item.Id),
			Name:  item.Name,
			Image: item.Image,
			Path:  item.Path,
		}

		result = append(result, asdf)
	}

	return result, nil
}

func (s OfficeApiService) GetSpaceInfo(id string) (SpaceInfo, error) {
	return SpaceInfo{}, nil
}

func (s OfficeApiService) GetSpaceTimeline(id string) ([]MeetingInfo, error) {
	return []MeetingInfo{}, nil
}

func (s OfficeApiService) GetSpaceState(id string) (string, error) {
	return "", nil
}

//-----------------------------
// CALENDAR
//-----------------------------

func (s OfficeApiService) GetSpaceCalendar(id string) ([]CalendarItem, error) {
	return []CalendarItem{}, nil
}

func (s OfficeApiService) GetDayInfo(spaceid string, year int, month int, day int) ([]MeetingInfo, error) {
	return []MeetingInfo{}, nil
}

func (s OfficeApiService) GetMeetingInfo(meetingid string) (MeetingInfo, error) {
	return MeetingInfo{}, nil
}

func (s OfficeApiService) AddNewMeeting(item MeetingInfo) error {
	return nil
}

func (s OfficeApiService) UpdateMeeting(item MeetingInfo) (MeetingInfo, error) {
	return MeetingInfo{}, nil
}

func (s OfficeApiService) DeleteMeeting(id string) error {
	return nil
}

//-----------------------------
// RECEPTION
//-----------------------------

func (s OfficeApiService) CallForClean(spaceid string) error {
	return nil
}
func (s OfficeApiService) CallReception(spaceid string) error {
	return nil
}
func (s OfficeApiService) SomethingElse(spaceid string, text string) error {
	return nil
}
func (s OfficeApiService) GetSortiment(spaceid string) ([]CellarSortimentItem, error) {
	return []CellarSortimentItem{}, nil
}
func (s OfficeApiService) PlaceOrder(spaceid string, item CellarOrder) error {
	return nil
}

//-----------------------------
// USER
//-----------------------------

func (s OfficeApiService) ValidatePin(pin string) (bool, error) {
	return true, nil
}

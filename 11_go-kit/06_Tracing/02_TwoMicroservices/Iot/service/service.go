package service

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//-----------------------------
//-----------------------------
// SERVICE METHODS
//-----------------------------
//-----------------------------

type Service interface {
	//Space
	GetAllSpaces() ([]CellarSpace, error)
	GetRootSpaces() ([]CellarSpace, error)
	GetSpaces(path string) ([]CellarSpace, error)
	RemoveSpaces(path string) error
	GetSpace(id string) (CellarSpace, error)
	AddSpace(item CellarSpace) (CellarSpace, error)
	RemoveSpace(id string) error
	UpdateSpace(item CellarSpace) (CellarSpace, error)
	//Senzor
	GetAllSenzors() ([]CellarSenzor, error)
	GetSenzors(path string) ([]CellarSenzor, error)
	RemoveSenzors(path string) error
	GetSenzor(id string) (CellarSenzor, error)
	AddSenzor(item CellarSenzor) (CellarSenzor, error)
	RemoveSenzor(id string) error
	UpdateSenzor(item CellarSenzor) (CellarSenzor, error)
	//Place
	GetAllPlaces() ([]CellarPlace, error)
	GetPlace(id string) (CellarPlace, error)
	AddPlace(item CellarPlace) (CellarPlace, error)
	RemovePlace(id string) error
	UpdatePlace(item CellarPlace) (CellarPlace, error)
	//MQTT
	PublishToMqtt(topic string, value string) error
}

type IotService struct {
	MongoDbUrl string
	Mqtturl    string
}

var Database = "HubDatabase"
var session *mgo.Session

func NewService(mongodburl string, mqtturl string) *IotService {

	var err error
	if session, err = mgo.Dial(mongodburl); err != nil {
		log.Fatal(err)
	}

	return &IotService{
		MongoDbUrl: mongodburl,
		Mqtturl:    mqtturl,
	}
}

//-----------------------------
// SPACE
//-----------------------------

func (s IotService) GetAllSpaces() ([]CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//SELECT
	var result []CellarSpace
	err := table.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) GetRootSpaces() ([]CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	// var bsonQuery = "{path: /^\/[A-Za-z0-9]*$/ }"

	//SELECT
	var result []CellarSpace
	err := table.Find(bson.M{"path": bson.M{"$regex": bson.RegEx{`^\/[A-Za-z0-9]*$`, ""}}}).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) GetSpaces(path string) ([]CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//SELECT
	var result []CellarSpace
	err := table.Find(bson.M{"path": bson.M{"$regex": bson.RegEx{`^` + path + `$`, ""}}}).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) RemoveSpaces(path string) error {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//SELECT
	_, err := table.RemoveAll(bson.M{"path": bson.M{"$regex": bson.RegEx{`^` + path, ""}}})
	if err != nil {
		return err
	}

	return nil
}

func (s IotService) GetSpace(id string) (CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//SELECT
	var result CellarSpace
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return CellarSpace{}, err
	}

	return result, nil
}

func (s IotService) AddSpace(item CellarSpace) (CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//New ID
	item.ID = bson.NewObjectId()

	//SELECT
	err := table.Insert(&item)
	if err != nil {
		return CellarSpace{}, err
	}

	return item, nil
}

func (s IotService) RemoveSpace(id string) error {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	//SELECT
	err := table.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}

	return nil
}

func (s IotService) UpdateSpace(item CellarSpace) (CellarSpace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Spaces")

	// Update
	colQuerier := bson.M{"_id": item.ID}
	err := table.Update(colQuerier, item)
	if err != nil {
		return CellarSpace{}, err
	}

	return item, nil
}

//-----------------------------
// SENZOR
//-----------------------------

func (s IotService) GetAllSenzors() ([]CellarSenzor, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//SELECT
	var result []CellarSenzor
	err := table.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) GetSenzors(path string) ([]CellarSenzor, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//SELECT
	var result []CellarSenzor
	err := table.Find(bson.M{"path": bson.M{"$regex": bson.RegEx{`^` + path + `$`, ""}}}).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) RemoveSenzors(path string) error {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//SELECT
	_, err := table.RemoveAll(bson.M{"path": bson.M{"$regex": bson.RegEx{`^` + path, ""}}})
	if err != nil {
		return err
	}

	return nil
}

func (s IotService) GetSenzor(id string) (CellarSenzor, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//SELECT
	var result CellarSenzor
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return CellarSenzor{}, err
	}

	return result, nil
}

func (s IotService) AddSenzor(item CellarSenzor) (CellarSenzor, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//New ID
	item.ID = bson.NewObjectId()

	//SELECT
	err := table.Insert(&item)
	if err != nil {
		return CellarSenzor{}, err
	}

	return item, nil
}

func (s IotService) RemoveSenzor(id string) error {
	sess := session.Clone()
	defer sess.Close()

	// fmt.Println("1")
	// fmt.Println(id)

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	//SELECT
	err := table.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// fmt.Println("2")

	return nil
}

func (s IotService) UpdateSenzor(item CellarSenzor) (CellarSenzor, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Senzors")

	// Update
	colQuerier := bson.M{"_id": item.ID}
	err := table.Update(colQuerier, item)
	if err != nil {
		return CellarSenzor{}, err
	}

	return item, nil
}

//-----------------------------
// PLACE
//-----------------------------

func (s IotService) GetAllPlaces() ([]CellarPlace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Places")

	//SELECT
	var result []CellarPlace
	err := table.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s IotService) GetPlace(id string) (CellarPlace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Places")

	//SELECT
	var result CellarPlace
	err := table.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return CellarPlace{}, err
	}

	return result, nil
}

func (s IotService) AddPlace(item CellarPlace) (CellarPlace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Places")

	//New ID
	item.ID = bson.NewObjectId()

	//SELECT
	err := table.Insert(&item)
	if err != nil {
		return CellarPlace{}, err
	}

	return item, nil
}

func (s IotService) RemovePlace(id string) error {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Places")

	//SELECT
	err := table.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return err
	}

	return nil
}

func (s IotService) UpdatePlace(item CellarPlace) (CellarPlace, error) {
	sess := session.Clone()
	defer sess.Close()

	//SELECT TABLE
	table := sess.DB(Database).C("Places")

	// Update
	colQuerier := bson.M{"_id": item.ID}
	err := table.Update(colQuerier, item)
	if err != nil {
		return CellarPlace{}, err
	}

	return item, nil
}

//--------------------------------------------------
//MQTT ---------------------------------------------
//--------------------------------------------------

func (s IotService) PublishToMqtt(topic string, value string) error {

	clientID := RandStringBytesMaskImprSrc(10)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			//low-level exception logging
			fmt.Println(err.Error())
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
	err2 := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  s.Mqtturl + ":1883",
		ClientID: []byte(clientID),
	})
	if err2 != nil {
		//low-level exception logging
		return err2
	}

	// Publish a message.
	err := cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS1,
		TopicName: []byte(topic),
		Message:   []byte(value),
	})
	if err != nil {
		return err
	}

	return nil
}

//---------------------------------------------------------
//HELPERS -------------------------------------------------
//---------------------------------------------------------

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

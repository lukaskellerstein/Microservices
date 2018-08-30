package service

import "gopkg.in/mgo.v2/bson"

//-----------------------------
//-----------------------------
// ENTITIES
//-----------------------------
//-----------------------------

type CellarSpace struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Image string        `json:"image" bson:"image"`
	Path  string        `json:"path" bson:"path"`
}

type CellarPlace struct {
	ID         bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	State      string        `json:"state" bson:"state"`
	Path       string        `json:"path" bson:"path"`
	Country    string        `json:"country" bson:"country"`
	City       string        `json:"city" bson:"city"`
	Street     string        `json:"street" bson:"street"`
	Zipcode    string        `json:"zipcode" bson:"zipcode"`
	Latitude   string        `json:"latitude" bson:"latitude"`
	Longtitude string        `json:"longtitude" bson:"longtitude"`
}

type CellarSenzor struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	State        string        `json:"state" bson:"state"`
	Path         string        `json:"path" bson:"path"`
	Type         string        `json:"type" bson:"type"`
	Firmware     string        `json:"firmware" bson:"firmware"`
	IpAddress    string        `json:"ipaddress" bson:"ipaddress"`
	WifiSSID     string        `json:"wifiSSID" bson:"wifiSSID"`
	WifiPassword string        `json:"wifiPassword" bson:"wifiPassword"`
	MQTTUrl      string        `json:"mqttUrl" bson:"mqttUrl"`
}

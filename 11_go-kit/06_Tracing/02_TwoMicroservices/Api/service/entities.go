package service

import "gopkg.in/mgo.v2/bson"
import "time"

//-----------------------------
//-----------------------------
// ENTITIES
//-----------------------------
//-----------------------------

type CellarMeeting struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	State  string        `json:"state" bson:"state"`
	Start  time.Time     `json:"start" bson:"start"`
	End    time.Time     `json:"end" bson:"end"`
	Author string        `json:"author" bson:"author"`
	Path   string        `json:"path" bson:"path"`
}

type CellarOrder struct {
	ID       bson.ObjectId               `json:"id" bson:"_id,omitempty"`
	State    string                      `json:"state" bson:"state"`
	Author   string                      `json:"author" bson:"author"`
	SumPrice string                      `json:"sumprice" bson:"sumprice"`
	Items    map[CellarSortimentItem]int `json:"items" bson:"items	"`
	Path     string                      `json:"path" bson:"path"`
}

type CellarSortimentItem struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	State string        `json:"state" bson:"state"`
	Path  string        `json:"path" bson:"path"`
	Price int           `json:"price" bson:"price"`
	Unit  string        `json:"unit" bson:"unit"`
}

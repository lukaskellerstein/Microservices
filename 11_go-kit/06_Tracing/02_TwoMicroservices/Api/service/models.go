package service

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//-----------------------------
//-----------------------------
// MODELS
//-----------------------------
//-----------------------------

type Space struct {
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name  string        `json:"name" bson:"name"`
	Image string        `json:"image" bson:"image"`
	Path  string        `json:"path" bson:"path"`
}

type SpaceInfo struct {
	Name          string      `json:"name" bson:"name"`
	State         string      `json:"state" bson:"state"`
	ActualMeeting MeetingInfo `json:"meetinginfo" bson:"meetinginfo"`
}

type MeetingInfo struct {
	Start  time.Time `json:"start" bson:"start"`
	End    time.Time `json:"end" bson:"end"`
	Author string    `json:"author" bson:"author"`
}

type CalendarItem struct {
	Date         time.Time `json:"date" bson:"date"`
	MeetingCount int       `json:"meetingcount" bson:"meetingcount"`
}

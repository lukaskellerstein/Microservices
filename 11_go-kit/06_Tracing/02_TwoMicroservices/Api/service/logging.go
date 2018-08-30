package service

import (
	"time"

	"github.com/go-kit/kit/log"
)

func NewLoggingMiddleware(logger log.Logger, s Service) Service {
	return &LoggingMiddleware{
		Service: s,
		logger:  logger}
}

type LoggingMiddleware struct {
	Service
	logger log.Logger
}

//-----------------------------
// SPACE
//-----------------------------

func (mw LoggingMiddleware) GetAllSpaces() (result []Space, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetAllSpaces",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetAllSpaces()
	return result, err
}

func (mw LoggingMiddleware) GetSpaceInfo(id string) (result SpaceInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpaceInfo",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpaceInfo(id)
	return result, err
}

func (mw LoggingMiddleware) GetSpaceTimeline(id string) (result []MeetingInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpaceTimeline",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpaceTimeline(id)
	return result, err
}

func (mw LoggingMiddleware) GetSpaceState(id string) (result string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpaceState",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpaceState(id)
	return result, err
}

//-----------------------------
// CALENDAR
//-----------------------------

func (mw LoggingMiddleware) GetSpaceCalendar(id string) (result []CalendarItem, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpaceCalendar",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpaceCalendar(id)
	return result, err
}

func (mw LoggingMiddleware) GetDayInfo(spaceid string, year int, month int, day int) (result []MeetingInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetDayInfo",
			"spaceid", spaceid,
			"year", year,
			"month", month,
			"day", day,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetDayInfo(spaceid, year, month, day)
	return result, err
}

func (mw LoggingMiddleware) GetMeetingInfo(meetingid string) (result MeetingInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetMeetingInfo",
			"meetingid", meetingid,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetMeetingInfo(meetingid)
	return result, err
}

func (mw LoggingMiddleware) AddNewMeeting(item MeetingInfo) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddNewMeeting",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.AddNewMeeting(item)
	return err
}

func (mw LoggingMiddleware) UpdateMeeting(item MeetingInfo) (result MeetingInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateMeeting",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateMeeting(item)
	return result, err
}

func (mw LoggingMiddleware) DeleteMeeting(id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "DeleteMeeting",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.DeleteMeeting(id)
	return err
}

//-----------------------------
// RECEPTION
//-----------------------------

func (mw LoggingMiddleware) CallForClean(spaceid string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "CallForClean",
			"spaceid", spaceid,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.CallForClean(spaceid)
	return err
}

func (mw LoggingMiddleware) CallReception(spaceid string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "CallReception",
			"spaceid", spaceid,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.CallReception(spaceid)
	return err
}

func (mw LoggingMiddleware) SomethingElse(spaceid string, text string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "SomethingElse",
			"spaceid", spaceid,
			"text", text,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.SomethingElse(spaceid, text)
	return err
}

func (mw LoggingMiddleware) GetSortiment(spaceid string) (result []CellarSortimentItem, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSortiment",
			"spaceid", spaceid,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSortiment(spaceid)
	return result, err
}

func (mw LoggingMiddleware) PlaceOrder(spaceid string, item CellarOrder) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "PlaceOrder",
			"spaceid", spaceid,
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.PlaceOrder(spaceid, item)
	return err
}

//-----------------------------
// USER
//-----------------------------

func (mw LoggingMiddleware) ValidatePin(pin string) (result bool, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "ValidatePin",
			"pin", pin,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.ValidatePin(pin)
	return result, err
}

package service

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

func NewMetricsMiddleware(requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	s Service) Service {
	return &MetricsMiddleware{
		Service:        s,
		requestCount:   requestCount,
		requestLatency: requestLatency,
	}
}

type MetricsMiddleware struct {
	Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

//-----------------------------
// SPACE
//-----------------------------

func (mw MetricsMiddleware) GetAllSpaces() (result []Space, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllSpaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetAllSpaces()
	return result, err
}

func (mw MetricsMiddleware) GetSpaceInfo(id string) (result SpaceInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpaceInfo"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpaceInfo(id)
	return result, err
}

func (mw MetricsMiddleware) GetSpaceTimeline(id string) (result []MeetingInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpaceTimeline"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpaceTimeline(id)
	return result, err
}

func (mw MetricsMiddleware) GetSpaceState(id string) (result string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpaceState"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpaceState(id)
	return result, err
}

//-----------------------------
// CALENDAR
//-----------------------------

func (mw MetricsMiddleware) GetSpaceCalendar(id string) (result []CalendarItem, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpaceCalendar"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpaceCalendar(id)
	return result, err
}

func (mw MetricsMiddleware) GetDayInfo(spaceid string, year int, month int, day int) (result []MeetingInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetDayInfo"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetDayInfo(spaceid, year, month, day)
	return result, err
}

func (mw MetricsMiddleware) GetMeetingInfo(meetingid string) (result MeetingInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetMeetingInfo"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetMeetingInfo(meetingid)
	return result, err
}

func (mw MetricsMiddleware) AddNewMeeting(item MeetingInfo) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddNewMeeting"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.AddNewMeeting(item)
	return err
}

func (mw MetricsMiddleware) UpdateMeeting(item MeetingInfo) (result MeetingInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateMeeting"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.UpdateMeeting(item)
	return result, err
}

func (mw MetricsMiddleware) DeleteMeeting(id string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteMeeting"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.DeleteMeeting(id)
	return err
}

//-----------------------------
// RECEPTION
//-----------------------------

func (mw MetricsMiddleware) CallForClean(spaceid string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CallForClean"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.CallForClean(spaceid)
	return err
}

func (mw MetricsMiddleware) CallReception(spaceid string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CallReception"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.CallReception(spaceid)
	return err
}

func (mw MetricsMiddleware) SomethingElse(spaceid string, text string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SomethingElse"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.SomethingElse(spaceid, text)
	return err
}

func (mw MetricsMiddleware) GetSortiment(spaceid string) (result []CellarSortimentItem, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSortiment"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSortiment(spaceid)
	return result, err
}

func (mw MetricsMiddleware) PlaceOrder(spaceid string, item CellarOrder) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "PlaceOrder"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.PlaceOrder(spaceid, item)
	return err
}

//-----------------------------
// USER
//-----------------------------

func (mw MetricsMiddleware) ValidatePin(pin string) (result bool, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "ValidatePin"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.ValidatePin(pin)
	return result, err
}

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

func (mw MetricsMiddleware) GetAllSpaces() (result []CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllSpaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetAllSpaces()
	return result, err
}

func (mw MetricsMiddleware) GetRootSpaces() (result []CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetRootSpaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetRootSpaces()
	return result, err
}

func (mw MetricsMiddleware) GetSpaces(path string) (result []CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpaces(path)
	return result, err
}
func (mw MetricsMiddleware) RemoveSpaces(path string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RemoveSpaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.RemoveSpaces(path)
	return err
}

func (mw MetricsMiddleware) GetSpace(id string) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSpace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSpace(id)
	return result, err
}
func (mw MetricsMiddleware) AddSpace(item CellarSpace) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddSpace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.AddSpace(item)
	return result, err
}
func (mw MetricsMiddleware) RemoveSpace(id string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RemoveSpace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.RemoveSpace(id)
	return err
}

func (mw MetricsMiddleware) UpdateSpace(item CellarSpace) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateSpace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.UpdateSpace(item)
	return result, err
}

//-----------------------------
// SENZOR
//-----------------------------

func (mw MetricsMiddleware) GetAllSenzors() (result []CellarSenzor, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllSenzors"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetAllSenzors()
	return result, err
}
func (mw MetricsMiddleware) GetSenzors(path string) (result []CellarSenzor, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSenzors"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSenzors(path)
	return result, err
}
func (mw MetricsMiddleware) RemoveSenzors(path string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RemoveSenzors"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.RemoveSenzors(path)
	return err
}

func (mw MetricsMiddleware) GetSenzor(id string) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetSenzor"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetSenzor(id)
	return result, err
}
func (mw MetricsMiddleware) AddSenzor(item CellarSenzor) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddSenzor"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.AddSenzor(item)
	return result, err
}
func (mw MetricsMiddleware) RemoveSenzor(id string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RemoveSenzor"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.RemoveSenzor(id)
	return err
}

func (mw MetricsMiddleware) UpdateSenzor(item CellarSenzor) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateSenzor"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.UpdateSenzor(item)
	return result, err
}

//-----------------------------
// PLACE
//-----------------------------

func (mw MetricsMiddleware) GetAllPlaces() (result []CellarPlace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllPlaces"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetAllPlaces()
	return result, err
}
func (mw MetricsMiddleware) GetPlace(id string) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetPlace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.GetPlace(id)
	return result, err
}
func (mw MetricsMiddleware) AddPlace(item CellarPlace) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddPlace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.AddPlace(item)
	return result, err
}
func (mw MetricsMiddleware) RemovePlace(id string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "RemovePlace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.RemovePlace(id)
	return err
}

func (mw MetricsMiddleware) UpdatePlace(item CellarPlace) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdatePlace"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result, err = mw.Service.UpdatePlace(item)
	return result, err
}

//-----------------------------
// MQTT
//-----------------------------

// Implement service functions and add label method for our metrics
func (mw MetricsMiddleware) PublishToMqtt(topic, value string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "PublishToMqtt"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	err = mw.Service.PublishToMqtt(topic, value)
	return err
}

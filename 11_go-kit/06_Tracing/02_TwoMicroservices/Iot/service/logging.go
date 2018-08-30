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

func (mw LoggingMiddleware) GetAllSpaces() (result []CellarSpace, err error) {
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

func (mw LoggingMiddleware) GetRootSpaces() (result []CellarSpace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetRootSpaces",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetRootSpaces()
	return result, err
}

func (mw LoggingMiddleware) GetSpaces(path string) (result []CellarSpace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpaces",
			"path", path,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpaces(path)
	return result, err
}

func (mw LoggingMiddleware) RemoveSpaces(path string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "RemoveSpaces",
			"path", path,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.RemoveSpaces(path)
	return err
}

func (mw LoggingMiddleware) GetSpace(id string) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSpace",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSpace(id)
	return result, err
}
func (mw LoggingMiddleware) AddSpace(item CellarSpace) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddSpace",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddSpace(item)
	return result, err
}
func (mw LoggingMiddleware) RemoveSpace(id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "RemoveSpace",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.RemoveSpace(id)
	return err
}

func (mw LoggingMiddleware) UpdateSpace(item CellarSpace) (result CellarSpace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateSpace",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateSpace(item)
	return result, err
}

//-----------------------------
// SENZOR
//-----------------------------

func (mw LoggingMiddleware) GetAllSenzors() (result []CellarSenzor, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetAllSenzors",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetAllSenzors()
	return result, err
}
func (mw LoggingMiddleware) GetSenzors(path string) (result []CellarSenzor, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSenzors",
			"path", path,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSenzors(path)
	return result, err
}
func (mw LoggingMiddleware) RemoveSenzors(path string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "RemoveSenzors",
			"path", path,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.RemoveSenzors(path)
	return err
}
func (mw LoggingMiddleware) GetSenzor(id string) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetSenzor",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetSenzor(id)
	return result, err
}
func (mw LoggingMiddleware) AddSenzor(item CellarSenzor) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddSenzor",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddSenzor(item)
	return result, err
}
func (mw LoggingMiddleware) RemoveSenzor(id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "RemoveSenzor",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.RemoveSenzor(id)
	return err
}
func (mw LoggingMiddleware) UpdateSenzor(item CellarSenzor) (result CellarSenzor, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdateSenzor",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdateSenzor(item)
	return result, err
}

//-----------------------------
// PLACE
//-----------------------------

func (mw LoggingMiddleware) GetAllPlaces() (result []CellarPlace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetAllPlaces",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetAllPlaces()
	return result, err
}

func (mw LoggingMiddleware) GetPlace(id string) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "GetPlace",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.GetPlace(id)
	return result, err
}

func (mw LoggingMiddleware) AddPlace(item CellarPlace) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "AddPlace",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.AddPlace(item)
	return result, err
}

func (mw LoggingMiddleware) RemovePlace(id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "RemovePlace",
			"id", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.RemovePlace(id)
	return err
}

func (mw LoggingMiddleware) UpdatePlace(item CellarPlace) (result CellarPlace, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "UpdatePlace",
			"item", item,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	result, err = mw.Service.UpdatePlace(item)
	return result, err
}

//-----------------------------
// MQTT
//-----------------------------
func (mw LoggingMiddleware) PublishToMqtt(topic, value string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "PublishToMqtt",
			"topic", topic,
			"value", value,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	err = mw.Service.PublishToMqtt(topic, value)
	return
}

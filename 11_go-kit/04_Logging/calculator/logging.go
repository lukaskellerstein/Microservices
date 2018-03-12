package calculator

import (
	"time"

	"github.com/go-kit/kit/log"
)

// implement function to return ServiceMiddleware
func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}

// Make a new type and wrap into Service interface
// Add logger property to this type
type loggingMiddleware struct {
	Service
	logger log.Logger
}

// Implement Service Interface for LoggingMiddleware
func (mw loggingMiddleware) Plus(a, b int) (result int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Plus",
			"A", a,
			"B", b,
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.Plus(a, b)
	return
}

func (mw loggingMiddleware) Minus(a, b int) (result int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Minus",
			"A", a,
			"B", b,
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.Minus(a, b)
	return
}

func (mw loggingMiddleware) Multiply(a, b int) (result int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Multiply",
			"A", a,
			"B", b,
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.Multiply(a, b)
	return
}

func (mw loggingMiddleware) Divide(a, b int) (result int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"function", "Divide",
			"A", a,
			"B", b,
			"result", result,
			"took", time.Since(begin),
		)
	}(time.Now())
	result = mw.Service.Divide(a, b)
	return
}

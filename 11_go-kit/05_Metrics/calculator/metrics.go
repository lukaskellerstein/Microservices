package calculator

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

func Metrics(requestCount metrics.Counter,
	requestLatency metrics.Histogram) ServiceMiddleware {
	return func(next Service) Service {
		return metricsMiddleware{
			next,
			requestCount,
			requestLatency,
		}
	}
}

// Make a new type and wrap into Service interface
// Add expected metrics property to this type
type metricsMiddleware struct {
	Service
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

// Implement service functions and add label method for our metrics
func (mw metricsMiddleware) Plus(a, b int) (result int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Plus"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result = mw.Service.Plus(a, b)
	fmt.Println("r:", result)
	return
}

func (mw metricsMiddleware) Minus(a, b int) (result int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Minus"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result = mw.Service.Minus(a, b)
	fmt.Println("r:", result)
	return
}

func (mw metricsMiddleware) Multiply(a, b int) (result int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Multiply"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result = mw.Service.Multiply(a, b)
	fmt.Println("r:", result)
	return
}

func (mw metricsMiddleware) Divide(a, b int) (result int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Divide"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	result = mw.Service.Divide(a, b)
	fmt.Println("r:", result)
	return
}

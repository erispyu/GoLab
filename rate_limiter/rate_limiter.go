package rate_limiter

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

var limiterMap sync.Map

func getLimiter(apiName, token string, limit int, timeRange time.Duration) *rate.Limiter {
	limiterKey := apiName + "_" + token + "_" + timeRange.String()
	limiter, ok := limiterMap.Load(limiterKey)
	if !ok {
		rateLimit := rate.Every(timeRange)
		limiter = rate.NewLimiter(rateLimit, limit)
		limiterMap.Store(limiterKey, limiter)
	}
	return limiter.(*rate.Limiter)
}

func AllowRequest(apiName, token string, limit int, timeRange time.Duration) bool {
	limiter := getLimiter(apiName, token, limit, timeRange)
	return limiter.Allow()
}

package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"time"
)

const QPSLimit = 500

func RateLimitMiddleware() gin.HandlerFunc {
	rateLimiter := ratelimit.New(QPSLimit)
	return func(c *gin.Context) {
		now := time.Now()
		time.Sleep(rateLimiter.Take().Sub(now))
		c.Next()
	}
}

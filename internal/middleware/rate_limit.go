package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var requests = make(map[string]int)
var mu sync.Mutex

func RateLimit() gin.HandlerFunc {
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			requests = make(map[string]int)
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		requests[ip]++
		if requests[ip] > 100 {
			mu.Unlock()
			c.AbortWithStatusJSON(429, gin.H{"error": "too many requests"})
			return
		}
		mu.Unlock()
		c.Next()
	}
}

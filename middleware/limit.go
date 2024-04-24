package middleware

import (
	"demo1/utils"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type Limit struct {
	//tokens per second
	rate              int64
	capacity          int64
	tokens            int64
	lastSuccessReqSec int64

	lock sync.Mutex
}

func New(rate, capacity int64) *Limit {
	return &Limit{
		rate:              rate,
		capacity:          capacity,
		tokens:            0,
		lastSuccessReqSec: time.Now().Unix(),
	}
}
func (l *Limit) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	now := time.Now().Unix()
	span := now - l.lastSuccessReqSec
	l.tokens = l.tokens + span*l.rate
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	if l.tokens > 0 {
		l.tokens--
		l.lastSuccessReqSec = now
		return true
	} else {
		return false
	}
}

var limiter *Limit

func init() {
	limiter = New(1, 3)
}

func LimiterMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		} else {
			utils.FailWithMsg(c, "rate limit")
			c.Abort()
			return
		}
	}
}

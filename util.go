package pika

import (
	"time"
)

func nowTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
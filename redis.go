package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)


func RedisClient() *redis.Pool{
	pool = &redis.Pool{
		MaxIdle: 10,
		IdleTimeout: 240*time.Second,
		Dial:func() (redis.Conn, error){
			return redis.Dial("tcp","localhost:6379")
		},
	}
	return pool
}

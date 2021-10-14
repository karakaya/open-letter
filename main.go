package main

import (
	"database/sql"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	dbErr error
	pool *redis.Pool
)

func main() {
	r := router()
	db = connect()

	RedisClient()

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Printf("http serve err %v \n", err)
	}

}

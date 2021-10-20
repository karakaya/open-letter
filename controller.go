package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
)

func index(w http.ResponseWriter, r *http.Request) {
	conn := pool.Get()
	defer conn.Close()
	rand, err := redis.String(conn.Do("RANDOMKEY"))
	if err != nil {
		panic(err)
	}

	letter, err := redis.StringMap(conn.Do("HGETALL", rand))
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(letter)
	if err != nil {
		panic(err)
	}

	w.Write(res)

}

func writeLetterView(w http.ResponseWriter, r *http.Request) {

}
func saveLetter(w http.ResponseWriter, r *http.Request) {

	var letter Letter
	err := json.NewDecoder(r.Body).Decode(&letter)
	if err != nil {
		panic(err)
	}
	letter.Time = time.Now()

	go func() {
		stmt, err := db.Prepare("INSERT INTO letters(title,body) values($1,$2)")
		if err != nil {
			panic(err)
		}

		_, err = stmt.Exec(letter.Title, letter.Body)
		if err != nil {
			panic(err)
		}

	}()

	conn := pool.Get()
	defer conn.Close()

	_, err = conn.Do("HSET", "letter:"+letter.Title, "title", "redis-title", "body", "redis-body")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)

}

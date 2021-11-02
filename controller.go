package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
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

func viewLetter(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	conn := pool.Get()

	defer conn.Close()
	var theLetter Letter
	letter, err := redis.StringMap(conn.Do("HGETALL", "letter:"+title))
	if len(letter) > 0 {
		w.WriteHeader(http.StatusFound)
		response, err := json.Marshal(letter)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(response)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if len(letter) == 0 {
		log.Printf("the letter (%s) not found in cache, looking from the database", title)
		stmt, err := db.Prepare("select * from letters where title = $1")
		if err != nil {
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Println(err)
			return
		}
		err = stmt.QueryRow(title).Scan(&theLetter.ID, &theLetter.Title, &theLetter.Body)
		if err != nil {
			log.Println(err)
			return
		}
		response, err := json.Marshal(theLetter)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusFound)
		w.Write(response)

	}

}
func saveLetter(w http.ResponseWriter, r *http.Request) {

	var letter Letter
	err := json.NewDecoder(r.Body).Decode(&letter)
	if err != nil {
		panic(err)
	}

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

	_, err = conn.Do("HSET", "letter:"+letter.Title, "title", letter.Title, "body", letter.Body)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)

}

package main

import "github.com/gorilla/mux"

func router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/letter/{title}", viewLetter).Methods("GET")
	r.HandleFunc("/write-letter", writeLetterView).Methods("GET")
	r.HandleFunc("/write-letter", saveLetter).Methods("POST")
	return r
}

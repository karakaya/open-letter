package main

import "time"

type Letter struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Time time.Time `json:"time,omitempty"`
}

package main

type Letter struct {
	ID    uint   `json:"id" redis:"id"`
	Title string `json:"title" redis:"title"`
	Body  string `json:"body" redis:"body"`
	//Time time.Time `json:"time,omitempty" redis:"time"`
}

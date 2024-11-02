package main

import (
	"fmt"
	"log"
	"net/http"
)

// var next_id int  shared data
type next_chan chan int

func (ch next_chan) handler(w http.ResponseWriter, r *http.Request) {


	//fmt.Fprintf(w, "<h1>You got %d<h1>", next_id)
	//next_id++    race condition
	fmt.Fprintf(w, "<h1>You got %d<h1>", <-ch)
}

func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	var n_ch next_chan = make(chan int)

	go counter(n_ch)

	http.HandleFunc("/", n_ch.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

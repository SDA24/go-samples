package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {

	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url: url, err: err, latency: 0}
	} else {
		t := time.Since(start).Round(time.Microsecond)
		ch <- result{url: url, err: nil, latency: t}
		resp.Body.Close()
	}
}

func main() {
	result := make(chan result)
	list := []string{
		"https://google.com",
		"https://github.com",
		"https://docker.com",
		"https://go.dev",
	}

	for _, url := range list {
		go get(url, result)
	}
	for range list {
		r := <-result
		if r.err != nil {
			fmt.Printf("%-20s |  %s\n", r.url, r.err)
		} else {
			fmt.Printf("%-20s |  %s\n", r.url, r.latency)
		}
	}
}

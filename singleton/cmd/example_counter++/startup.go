package main

import "sync"

var (
	counter = 0
	wg      = sync.WaitGroup{}
)

func increment_counter() {
	counter++
	wg.Done()
}

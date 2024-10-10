package main

import "sync"

var (
	counter = 0
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment_counter() {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

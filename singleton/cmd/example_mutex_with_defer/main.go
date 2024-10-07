package main

import (
	"fmt"
	"sync"
)

var (
	wg      sync.WaitGroup
	mutex   sync.Mutex
	counter = 0
	n       = 2
)

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		mutex.Lock()
		counter++
		defer mutex.Unlock()
		if n%2 == 0 {
			fmt.Println("goroutine 1: n%2 = 0")
			return
		}
	}()
	go func() {
		defer wg.Done()
		mutex.Lock()
		counter++
		defer mutex.Unlock()
		if n%2 == 0 {
			fmt.Println("goroutine 2: n%2 = 0")
			return
		}
	}()
	wg.Wait()
	fmt.Println("done, counter is: ", counter)
}

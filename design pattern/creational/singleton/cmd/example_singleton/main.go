package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("calling new_class: ")
			new_class()
		}()
	}
	wg.Wait()
}

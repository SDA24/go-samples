package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment_counter()
	}
	wg.Wait()
	fmt.Println("counter is: ", counter)
}

package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment_counter()
	}
	wg.Wait()
	fmt.Println("counter is: ", counter)
}

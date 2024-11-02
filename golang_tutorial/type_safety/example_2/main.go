package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var count int

	for {
		var value float64

		_, err := fmt.Fscanln(os.Stdin, &value)
		if err != nil {
			break
		}

		sum += value
		count++
	}

	if count == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is: ", sum/float64(count))
}

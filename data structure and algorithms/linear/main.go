package main

import "fmt"

func linear(data_list []int, key int) bool {
	for i := 0; i < len(data_list); i++ {
		if data_list[i] == key {
			return true
		}
	}
	return false
}

func main() {
	var key = 24
	var items = []int{23, 6, 76, 3, 678, 9, 0, 24, 67, 30}
	fmt.Println(linear(items, key))
}

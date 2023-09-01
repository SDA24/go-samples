package main

import (
	call_external_api "LSP/service"
	"fmt"
	"log"
)

func main() {
	factory := call_external_api.ApiFactory{}

	example_api := factory.CreateRepository("https://api.example.com/data")
	example_data, err := example_api.GetData()
	if err != nil {
		log.Fatalf("Error calling Example API: %v", err)
	}
	fmt.Printf("Example API data: %s\n", example_data)

	another_example := factory.CreateRepository("https://api.anotherexample.com/data")
	another_example_data, err := another_example.GetData()
	if err != nil {
		log.Fatalf("Error calling Another Example API: %v", err)
	}
	fmt.Printf("Another Example API data: %s\n", another_example_data)
}

package main

import "fmt"

// Global type
type singleton map[string]string

var (
	//Global variable
	instance singleton
)

func new_class() singleton {
	if instance == nil {
		fmt.Println("instance is nil")
		instance = make(singleton) // <== not thread safe
	}
	return instance
}

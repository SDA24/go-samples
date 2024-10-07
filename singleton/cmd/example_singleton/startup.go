package main

import (
	"fmt"
	"sync"
)

// global type
type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func new_class() singleton {
	once.Do(func() { //<== atomic, does not allowing repeat
		instance = make(singleton) //<== thread safe
		instance["name"] = "Saeed"
		fmt.Printf("instance is: %+v\n", instance)
	})
	return instance
}

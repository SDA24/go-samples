package main

import (
	"fmt"
	hello "hello_world/service"
	"os"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}

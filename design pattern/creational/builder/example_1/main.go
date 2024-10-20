package main

import (
	element_service "builder/service"
	"builder/service_abstraction"
	"fmt"
)

func main() {
	adapter := element_service.NewHtmlBuilder("ul")
	adapter.AddChildFluent("li", "hello").AddChild("li", "world")
	serve := service_abstraction.New_element(adapter)
	fmt.Println(serve.String())
}

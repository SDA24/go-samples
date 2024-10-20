package main

import (
	person_service "builder/service"
	person_service_abstraction "builder/service_abstraction"
	"fmt"
)

func main() {
	adapter := person_service.NewPersonBuilder()
	adapter.
		Lives().
		At("123 IRAN Road").
		In("IRAN").
		WithPostcode("84256").
		Works().
		At("SDA24").
		AsA("Programmer").
		Earning(123000)
	person := adapter.Build()
	serve := person_service_abstraction.New_person(person)
	fmt.Println(string(serve.Person_info_json()))
}

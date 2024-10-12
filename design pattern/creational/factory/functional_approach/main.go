package main

import "fmt"

type Employee struct {
	Name, Position string
	Annual_income  int
}

func New_employee_factory(position string, annual int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{Name: name, Position: position, Annual_income: annual}
	}
}

func main() {
	developer_factory := New_employee_factory("golang developer", 30000)
	develeper_porson := developer_factory("Saeed")
	fmt.Println(develeper_porson)
}

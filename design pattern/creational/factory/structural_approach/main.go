package main

import "fmt"

type Employee struct {
	Name, Position string
	Annual_income  int
}

type Employee_factory struct {
	Position      string
	Annual_income int
}

func New_employee_factory(position string, annual int) *Employee_factory {
	return &Employee_factory{Position: position, Annual_income: annual}
}

func (e *Employee_factory) Creat(name string) *Employee {
	return &Employee{
		Name:          name,
		Position:      e.Position,
		Annual_income: e.Annual_income,
	}
}

func main() {
	employee := New_employee_factory("Developer", 40).Creat("Saeed")
	fmt.Println(*employee)
}

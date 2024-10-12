package main

import "fmt"

type Employee struct {
	Position, Name string
	Annual_income  int
}

const (
	Developer = iota
	Manager
)

func New_employee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{Position: "Developer", Name: "", Annual_income: 30000}
	case Manager:
		return &Employee{Position: "Manager", Name: "", Annual_income: 35000}
	default:
		panic("unsupported role")
	}
}

func main() {
	person := New_employee(Developer)
	person.Name = "Saeed"
	fmt.Println(*person)
}

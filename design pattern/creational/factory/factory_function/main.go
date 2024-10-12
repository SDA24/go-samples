package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func New_person(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	// initialized directly
	inited_person := Person{"Saeed", 40}
	fmt.Println(inited_person)

	// use a constructor
	constructed_person := New_person("Delaram", 3)
	fmt.Println(*constructed_person)
}

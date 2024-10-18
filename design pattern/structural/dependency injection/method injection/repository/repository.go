package repository

import "fmt"

type Irepository interface {
	Save(data string) error
}

type database_repository struct{}

func New_database_repository() database_repository {
	return database_repository{}
}

func (d database_repository) Save(data string) error {
	fmt.Println("saving data: ", data)
	return nil
}

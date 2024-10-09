package main

import "gorm.io/gorm"

var (
	db gorm.DB
)

func init_db() {
	db = gorm.DB{}
}

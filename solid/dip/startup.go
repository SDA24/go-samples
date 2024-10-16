package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func getDB() *gorm.DB {
	once.Do(func() {
		db = func() *gorm.DB {
			dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("Failed to connect to database: " + err.Error())
			}
			fmt.Println("Connection to database!")
			return db
		}()
	})
	return db
}

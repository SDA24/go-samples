package main

import (
	"constructor_injection/logger"
	"constructor_injection/repository"
	"constructor_injection/service"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	logger := logger.New_console_logger()
	repository := repository.New_database_repository()
	my_service := service.New_my_service(logger, repository)
	err := my_service.Do_work(ctx, "sample data")
	if err != nil {
		fmt.Println("an error occurred: ", err)
	}
}

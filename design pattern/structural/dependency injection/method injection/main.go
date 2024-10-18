package main

import (
	"bufio"
	"fmt"
	"method_injection/logger"
	"method_injection/repository"
	"method_injection/service"
	"os"
)

func main() {
	fmt.Println("input text l/n:")
	reader := bufio.NewReader(os.Stdin)
	rune, size, err := reader.ReadRune()
	if err != nil {
		fmt.Println("an error occurred: ", err)
	}
	fmt.Printf("input character is %c , size is %d\n", rune, size)

	repository := repository.New_database_repository()
	service := service.New_my_service(repository)
	switch rune {
	case 'l':
		service.Set_logger(logger.New_logger())
		fmt.Println("logger is activated")
	case 'n':
		service.Set_notifier(logger.New_notifier())
		fmt.Println("notifier is activated")
	}
	fmt.Println("work completed")
}

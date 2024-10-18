package logger

import "fmt"

type Ilogger interface {
	Log(message string)
}

type logger struct{}

func New_logger() logger {
	return logger{}
}

func (l logger) Log(message string) {
	fmt.Println("log: ", message)
}

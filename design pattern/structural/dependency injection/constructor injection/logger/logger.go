package logger

import "fmt"

type Ilogger interface {
	Log(message string)
}

type console_logger struct{}

func New_console_logger() console_logger {
	return console_logger{}
}

func (c console_logger) Log(message string) {
	fmt.Println("log: ", message)
}

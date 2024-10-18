package logger

import (
	"fmt"
)

type Inotifier interface {
	Send(message string) error
}

type notifier struct{}

func New_notifier() notifier {
	return notifier{}
}

func (n notifier) Send(message string) error {
	fmt.Println("send notiffication")
	return nil
}

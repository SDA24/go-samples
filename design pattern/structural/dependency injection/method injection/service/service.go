package service

import (
	"context"
	"fmt"
	"method_injection/logger"
	"method_injection/repository"
)

type Iservice interface {
	Do_work(ctx context.Context, data string) error
	Set_logger(logger.Ilogger) Iservice
	Set_notifier(logger.Inotifier) Iservice
}

type my_service struct {
	logger     logger.Ilogger
	notifier   logger.Inotifier
	repository repository.Irepository
}

func New_my_service(repository repository.Irepository) *my_service {
	return &my_service{
		repository: repository,
	}
}

func (m *my_service) Set_logger(logger logger.Ilogger) Iservice {
	m.logger = logger
	return m
}
func (m *my_service) Set_notifier(notifier logger.Inotifier) Iservice {
	m.notifier = notifier
	return m
}
func (m *my_service) Do_work(ctx context.Context, data string) error {
	if m.logger != nil {
		m.logger.Log("starting work")
	} else if m.notifier != nil {
		err := m.notifier.Send("starting work")
		if err != nil {
			return err
		}
	}
	m.repository.Save(data)
	fmt.Println("sample data is: ", data)
	return nil
}

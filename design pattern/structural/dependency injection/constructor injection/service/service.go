package service

import (
	"constructor_injection/logger"
	"constructor_injection/repository"
	"context"
)

type Iservice interface {
	Do_work(ctx context.Context, data string) error
}

type my_service struct {
	logger     logger.Ilogger
	repository repository.Irepository
}

func New_my_service(logger logger.Ilogger, repository repository.Irepository) my_service {
	return my_service{
		logger:     logger,
		repository: repository,
	}
}

func (m my_service) Do_work(ctx context.Context, data string) error {
	m.logger.Log("starting work")
	if err := m.repository.Save(data); err != nil {
		m.logger.Log("an error occurred")
		return err
	}
	m.logger.Log("work completed")
	return nil
}

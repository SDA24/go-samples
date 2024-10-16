package service

import (
	repo_abstraction "dip/repository_abstraction"
	"fmt"
)

type EmailService struct {
	repository repo_abstraction.UserRepository
	// some email sender
}

func NewEmailService(repository repo_abstraction.UserRepository) *EmailService {
	return &EmailService{
		repository: repository,
	}
}
func (s *EmailService) SendRegistrationEmail(userID uint) error {
	user, err := s.repository.GetByID(userID)
	if err != nil {
		return err
	}
	// send email for specific user
	fmt.Println("we will send email for this user: ",user)
	return nil
}

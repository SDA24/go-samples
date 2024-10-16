package service

import (
	"fmt"

	"adapter/repository"
	repo_abstraction "adapter/repository_abstraction"

	"gorm.io/gorm"
)

type EmailService struct {
	DB *gorm.DB
	// some email sender
}

func NewEmailService(db *gorm.DB) *EmailService {
	return &EmailService{
		DB: db,
	}
}
func (s *EmailService) SendRegistrationEmail(userID uint) error {
	adapter := repository.NewUserDatabaseRepository(s.DB)
	user, err := repo_abstraction.New_user_repository(adapter).GetByID(userID)
	if err != nil {
		return err
	}
	// send email for specific user
	fmt.Println("we will send email for this user: ", user)
	return nil
}

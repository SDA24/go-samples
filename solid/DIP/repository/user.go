package repository

import (
	repo_abstraction "dip/repository_abstraction"

	"gorm.io/gorm"
)

type UserDatabaseRepository struct {
	db *gorm.DB
}

var _ repo_abstraction.UserRepository = &UserDatabaseRepository{}

func NewUserDatabaseRepository(db *gorm.DB) repo_abstraction.UserRepository {
	return &UserDatabaseRepository{
		db: db,
	}
}

func (r *UserDatabaseRepository) GetByID(id uint) (*repo_abstraction.User, error) {
	user := UserGorm{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user.ToUser(), nil
}

type UserGorm struct {
	// some fields
}

func (g UserGorm) ToUser() *repo_abstraction.User {
	return &repo_abstraction.User{
		// some fields
	}
}

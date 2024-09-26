package repository

import (
	"github.com/matthewyuh246/blogbackend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(user *models.User, email string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"errors"
	"time"

	"github.com/iqbal13/finalprojecthactiv8/helpers"
	"github.com/iqbal13/finalprojecthactiv8/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Register(context.Context, *models.User) error
	Login(context.Context, *models.User) error
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Register(ctx context.Context, user *models.User) (err error) {

	if err = userRepository.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}

	return
}

func (userRepository *userRepository) Login(ctx context.Context, user *models.User) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	password := user.Password

	if err = userRepository.db.WithContext(ctx).Where("email = ? OR username = ?", user.Email, user.Username).Take(&user).Error; err != nil {
		return errors.New("the email you entered are not found")
	}

	if isValid := helpers.Compare([]byte(user.Password), []byte(password)); !isValid {
		return errors.New("the credential you entered are wrong")
	}

	return
}

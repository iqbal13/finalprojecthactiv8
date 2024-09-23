package usecase

import (
	"context"

	"github.com/iqbal13/finalprojecthactiv8/models"
	"github.com/iqbal13/finalprojecthactiv8/repository"
)

type UserUseCase interface {
	Register(context.Context, *models.User) error
	Login(context.Context, *models.User) error
}
type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *userUseCase {
	return &userUseCase{userRepository}
}

func (userUseCase *userUseCase) Register(ctx context.Context, user *models.User) (err error) {
	if err = userUseCase.userRepository.Register(ctx, user); err != nil {
		return err
	}

	return
}

func (userUseCase *userUseCase) Login(ctx context.Context, user *models.User) (err error) {
	if err = userUseCase.userRepository.Login(ctx, user); err != nil {
		return err
	}

	return
}

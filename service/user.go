package service

import (
	"context"
	"gin_example_with_generic/generic"
	"gin_example_with_generic/store/model"
	"gin_example_with_generic/store/repository"
	"gin_example_with_generic/types/request"
	"gorm.io/gorm"
)

type UserService struct {
	*generic.Service[request.User, model.User]
	userRepository repository.UserInterface
}

func NewUserService(db *gorm.DB) UserInterface {
	userRepository := repository.NewUserRepository(db)
	return &UserService{
		generic.NewService[request.User, model.User](userRepository),
		userRepository,
	}
}

func (u *UserService) ListByName(ctx context.Context, name string) ([]*model.User, error) {
	return u.userRepository.ListByName(ctx, name)
}

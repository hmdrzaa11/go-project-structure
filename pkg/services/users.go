package services

import "github.com/hmdrzaa11/example-go-api/pkg/domains"

type UserService interface {
	GetUserById(id int) (*domains.User, error)
	GetUserByEmail(email string) (*domains.User, error)
}

type DefaultUsersService struct {
	repo domains.UserRepository
}

func NewDefaultUserService(repo domains.UserRepository) *DefaultUsersService {
	return &DefaultUsersService{repo: repo}
}

func (us *DefaultUsersService) GetUserById(id int) (*domains.User, error) {
	return nil, nil
}

func (us *DefaultUsersService) GetUserByEmail(email string) (*domains.User, error) {
	return nil, nil
}

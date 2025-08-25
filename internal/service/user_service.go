package service

import (
	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
)

type UserService interface {
	List() ([]domain.User, error)
	Get(id string) (*domain.User, error)
}

type userService struct {
	client ports.UserClient
}

func NewUserService(c ports.UserClient) UserService {
	return &userService{client: c}
}

func (s *userService) List() ([]domain.User, error) {
	return s.client.ListUsers()
}

func (s *userService) Get(id string) (*domain.User, error) {
	return s.client.GetUser(id)
}

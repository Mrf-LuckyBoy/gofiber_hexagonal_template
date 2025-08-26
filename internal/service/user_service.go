package service

import (
	"fmt"

	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/cache"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
)

type UserService interface {
	List() ([]domain.User, error)
	Get(id string) (*domain.User, error)
}

type userService struct {
	client ports.UserClient
	cache  *cache.RistrettoCache
}

func NewUserService(c ports.UserClient, cache *cache.RistrettoCache) UserService {
	return &userService{client: c, cache: cache}
}

func (s *userService) List() ([]domain.User, error) {
	if data, found := s.cache.Get("users:all"); found {
		return data.([]domain.User), nil
	}
	users, err := s.client.ListUsers()
	if err != nil {
		return nil, err
	}

	s.cache.Set("users:all", users, int64(len(users)))
	return users, nil
}

func (s *userService) Get(id string) (*domain.User, error) {
	key := fmt.Sprintf("user:%s", id)
	if data, found := s.cache.Get(key); found {
		return data.(*domain.User), nil
	}

	user, err := s.client.GetUser(id)
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, user, 1)
	return user, nil
}

package service_test

import (
	"testing"

	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserClient struct {
	mock.Mock
}

func (m *mockUserClient) ListUsers() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *mockUserClient) GetUser(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

// // --- Mock repository (DB) ---
// type mockUserRepository struct {
// 	mock.Mock
// }

// func (m *mockUserRepository) Save(u *domain.User) error {
// 	args := m.Called(u)
// 	return args.Error(0)
// }

func TestUserService_List(t *testing.T) {
	mockClient := new(mockUserClient)
	mockClient.On("ListUsers").Return([]domain.User{
		{Name: "Emard Inc", Email: "Haleigh.Ruecker@gmail.com", PhoneNumber: "359.999.3884 x89984", ID: "1"},
	}, nil)

	service := service.NewUserService(mockClient, nil)

	users, err := service.List()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "Emard Inc", users[0].Name)

	mockClient.AssertExpectations(t)
}

// internal/adapters/thirdparty/user_api.go
package thirdparty

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
)

type UserAPIClient struct {
	baseURL string
}

func NewUserAPIClient(baseURL string) ports.UserClient {
	return &UserAPIClient{baseURL: baseURL}
}

func (u *UserAPIClient) ListUsers() ([]domain.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users", u.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []domain.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserAPIClient) GetUser(id string) (*domain.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%s", u.baseURL, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user domain.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

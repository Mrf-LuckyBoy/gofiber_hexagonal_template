// internal/core/ports/user_client.go
package ports

import "github.com/Mrf-LuckyBoy/test-go/internal/core/domain"

type UserClient interface {
	ListUsers() ([]domain.User, error)
	GetUser(id string) (*domain.User, error)
}

package ports

import "github.com/Mrf-LuckyBoy/test-go/internal/core/domain"

type BookService interface {
	Create(title, author string) (*domain.Book, error)
	GetByID(id string) (*domain.Book, error)
	List() ([]domain.Book, error)
	Update(id, title, author string) (*domain.Book, error)
	Delete(id string) error
}

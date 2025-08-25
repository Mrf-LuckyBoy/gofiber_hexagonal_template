package ports

import "github.com/Mrf-LuckyBoy/test-go/internal/core/domain"

type BookRepository interface {
	Create(book *domain.Book) error
	GetByID(id string) (*domain.Book, error)
	List() ([]domain.Book, error)
	Update(book *domain.Book) error
	Delete(id string) error
}

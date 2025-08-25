package mariadb

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
)

var _ ports.BookRepository = (*BookRepositoryMariaDB)(nil)

type BookRepositoryMariaDB struct {
	db *gorm.DB
}

func NewBookRepositoryMariaDB(db *gorm.DB) *BookRepositoryMariaDB {
	return &BookRepositoryMariaDB{db: db}
}

// AutoMigrate ensures schema is ready
func (r *BookRepositoryMariaDB) AutoMigrate() error {
	return r.db.AutoMigrate(&domain.Book{})
}

func (r *BookRepositoryMariaDB) Create(book *domain.Book) error {
	return r.db.Create(book).Error
}

func (r *BookRepositoryMariaDB) GetByID(id string) (*domain.Book, error) {
	var b domain.Book
	if err := r.db.First(&b, "id = ?", id).Error; err != nil {
		return nil, errors.New("book not found")
	}
	return &b, nil
}

func (r *BookRepositoryMariaDB) List() ([]domain.Book, error) {
	var books []domain.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepositoryMariaDB) Update(book *domain.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepositoryMariaDB) Delete(id string) error {
	return r.db.Delete(&domain.Book{}, "id = ?", id).Error
}

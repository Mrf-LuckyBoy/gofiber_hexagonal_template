package service

import (
	"errors"
	"time"

	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
	"github.com/google/uuid"
)

type bookService struct {
	repo ports.BookRepository
}

func NewBookService(repo ports.BookRepository) ports.BookService {
	return &bookService{repo: repo}
}

func (s *bookService) Create(title, author string) (*domain.Book, error) {
	if title == "" || author == "" {
		return nil, errors.New("title and author are required")
	}
	now := time.Now().UTC()
	b := &domain.Book{
		ID:        uuid.NewString(),
		Title:     title,
		Author:    author,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.Create(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *bookService) GetByID(id string) (*domain.Book, error) {
	if id == "" {
		return nil, errors.New("id required")
	}
	return s.repo.GetByID(id)
}

func (s *bookService) List() ([]domain.Book, error) {
	return s.repo.List()
}

func (s *bookService) Update(id, title, author string) (*domain.Book, error) {
	if id == "" {
		return nil, errors.New("id required")
	}
	b, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if title != "" {
		b.Title = title
	}
	if author != "" {
		b.Author = author
	}
	b.UpdatedAt = time.Now().UTC()
	if err := s.repo.Update(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *bookService) Delete(id string) error {
	if id == "" {
		return errors.New("id required")
	}
	return s.repo.Delete(id)
}

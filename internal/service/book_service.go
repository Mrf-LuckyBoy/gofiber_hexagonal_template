package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/cache"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
	"github.com/google/uuid"
)

type bookService struct {
	repo  ports.BookRepository
	cache *cache.RistrettoCache
}

func NewBookService(repo ports.BookRepository, cache *cache.RistrettoCache) ports.BookService {
	return &bookService{repo: repo, cache: cache}
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
	key := fmt.Sprintf("book:%s", id)
	if data, found := s.cache.Get(key); found {
		return data.(*domain.Book), nil
	}

	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	s.cache.Set(key, book, 1)
	return book, nil
}

func (s *bookService) List() ([]domain.Book, error) {
	if data, found := s.cache.Get("books:all"); found {
		return data.([]domain.Book), nil
	}
	books, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	s.cache.Set("books:all", books, int64(len(books)))
	return books, nil
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

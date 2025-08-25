package memory

import (
	"errors"
	"sync"

	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/ports"
)

var _ ports.BookRepository = (*BookRepositoryMemory)(nil)

type BookRepositoryMemory struct {
	mu    sync.RWMutex
	items map[string]domain.Book
}

func NewBookRepositoryMemory() *BookRepositoryMemory {
	return &BookRepositoryMemory{
		items: make(map[string]domain.Book),
	}
}

func (r *BookRepositoryMemory) Create(book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[book.ID]; ok {
		return errors.New("book already exists")
	}
	r.items[book.ID] = *book
	return nil
}

func (r *BookRepositoryMemory) GetByID(id string) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	b, ok := r.items[id]
	if !ok {
		return nil, errors.New("book not found")
	}
	// return copy
	bb := b
	return &bb, nil
}

func (r *BookRepositoryMemory) List() ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]domain.Book, 0, len(r.items))
	for _, v := range r.items {
		out = append(out, v)
	}
	return out, nil
}

func (r *BookRepositoryMemory) Update(book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[book.ID]; !ok {
		return errors.New("book not found")
	}
	r.items[book.ID] = *book
	return nil
}

func (r *BookRepositoryMemory) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.items[id]; !ok {
		return errors.New("book not found")
	}
	delete(r.items, id)
	return nil
}

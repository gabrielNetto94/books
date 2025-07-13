package bookrepository

import (
	"books/internal/core/domain"
	"books/internal/core/repositories/cache"
	"books/internal/core/repositories/db"
	"context"
	"encoding/json"
	"fmt"
)

type UserRepositoryImpl struct {
	db    db.DatabaseRepositoryInterface
	cache cache.CacheRepositoryInterface
}

func NewBookRepository(db db.DatabaseRepositoryInterface, cache cache.CacheRepositoryInterface) *UserRepositoryImpl {
	return &UserRepositoryImpl{db, cache}
}

type BookRepository interface {
	Save(ctx context.Context, book domain.Book) error
	FindById(ctx context.Context, id string) (domain.Book, error)
	ListAll(ctx context.Context) ([]domain.Book, error)
	Update(ctx context.Context, book domain.Book) error
}

func (s *UserRepositoryImpl) Save(ctx context.Context, book domain.Book) error {

	if err := s.db.Create(ctx, book); err != nil {
		return err
	}
	_ = s.cache.Set(ctx, book.Id, book)
	return nil
}

func (s *UserRepositoryImpl) Update(ctx context.Context, book domain.Book) error {

	// err := s.db.Save(book).Error
	// if err == nil {
	// 	_ = s.cache.Del(ctx, book.Id)
	// 	return nil
	// }

	// return err
	return fmt.Errorf("not implemented")
}

func (s *UserRepositoryImpl) FindById(ctx context.Context, id string) (domain.Book, error) {
	var book = domain.Book{Id: id}
	if err := s.cache.Get(ctx, id, &book); err == nil {
		return book, nil
	}

	if err := s.db.Find(ctx, &book); err != nil {
		return book, err
	}

	bookBytes, err := json.Marshal(book)
	if err == nil {
		_ = s.cache.Set(ctx, book.Id, bookBytes)
	}

	return book, nil
}

func (s *UserRepositoryImpl) ListAll(ctx context.Context) ([]domain.Book, error) {

	// var books []domain.Book

	// return books, s.db.WithContext(ctx).Find(&books).Error
	return nil, fmt.Errorf("not implemented")
}

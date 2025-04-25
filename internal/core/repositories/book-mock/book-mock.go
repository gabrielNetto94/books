package bookmock

import (
	"books/internal/core/domain"
	datafake "books/pkg/data-fake"
	"math/rand"

	"github.com/google/uuid"
)

type UserRepositoryImplMock struct {
	fake datafake.DataFake
}

func NewBookRepositoryMock(dataFake datafake.DataFake) *UserRepositoryImplMock {
	return &UserRepositoryImplMock{dataFake}
}

func (s *UserRepositoryImplMock) Save(book domain.Book) error {
	return nil
}

func (s *UserRepositoryImplMock) FindById(id string) (domain.Book, error) {
	return domain.Book{Id: uuid.NewString(), Title: s.fake.FirstName(), Author: s.fake.FullName()}, nil
}

func (s *UserRepositoryImplMock) ListAll() ([]domain.Book, error) {
	count := rand.Intn(10) + 1
	books := make([]domain.Book, count)
	for i := range count {
		books[i] = domain.Book{Id: uuid.NewString(), Title: s.fake.FirstName(), Author: s.fake.FullName()}
	}
	return books, nil
}
func (s *UserRepositoryImplMock) Update(book domain.Book) error {
	return nil
}

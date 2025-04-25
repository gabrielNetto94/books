package bookmock

import (
	"books/internal/core/domain"
)

type UserRepositoryImplMock struct {
}

func NewBookRepositoryMock() *UserRepositoryImplMock {
	return &UserRepositoryImplMock{}
}

func (s *UserRepositoryImplMock) Save(book domain.Book) error {

	return nil
}
func (s *UserRepositoryImplMock) FindById(id string) (domain.Book, error) {
	return domain.Book{Id: "asda", Title: "book"}, nil
}
func (s *UserRepositoryImplMock) ListAll() ([]domain.Book, error) {

	var asd = []domain.Book{
		{Id: "asda", Title: "asdas"},
		{Id: "asdasd", Title: ""},
	}
	return asd, nil
}
func (s *UserRepositoryImplMock) Update(book domain.Book) error {
	return nil
}

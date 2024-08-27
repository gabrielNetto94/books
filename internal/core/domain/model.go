package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Book struct {
	Id     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

// @todo verificar se mantém model separada para cada domínio
type DomainError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Error   error  `json:"error,omitempty"`
}

func (b Book) Validate() error {
	var errs []string

	if _, err := uuid.Parse(b.Id); err != nil {
		errs = append(errs, "Invalid book ID")
	}
	if b.Author == "" {
		errs = append(errs, "Invalid author")
	}
	if b.Title == "" {
		errs = append(errs, "Invalid title")
	}
	if b.Desc == "" {
		errs = append(errs, "Invalid description")
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.New(strings.Join(errs, ", "))
}

package domain

import (
	"errors"
	"strings"
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

func (b Book) Validate() (err error) {

	var errStrings []string
	if b.Author == "" {
		errStrings = append(errStrings, "Invalid author")
	}
	if b.Title == "" {
		errStrings = append(errStrings, "Invalid tittle")
	}
	if b.Desc == "" {
		errStrings = append(errStrings, "Invalid desc")
	}

	if len(errStrings) == 0 {
		return nil
	}

	return errors.New(strings.Join(errStrings, ","))
}

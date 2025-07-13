package services

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"context"

	userrepository "books/internal/core/repositories/user"
	"books/internal/infra/log"
	"books/internal/ports"

	"github.com/google/uuid"
)

type UserService struct {
	repo userrepository.UserRepository
	log  log.Logger
}

func NewUserService(repo userrepository.UserRepository, log log.Logger) ports.UserServiceInterface {
	return &UserService{repo, log}
}

func (s *UserService) CreateUser(ctx context.Context, user domain.User) (string, *domain.DomainError) {

	user.Id = uuid.New().String()
	err := s.repo.Save(ctx, user)
	if err != nil {
		s.log.Error("Failed to save user: ", err)
		return "", &domain.DomainError{
			Message: "error on create user",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return user.Id, nil
}

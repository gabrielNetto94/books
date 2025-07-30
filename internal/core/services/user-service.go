package services

import (
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"context"

	userrepository "books/internal/core/repositories/user"
	"books/internal/infra/log"
	"books/internal/ports"
	"books/pkg/observability"

	"github.com/google/uuid"
)

type UserService struct {
	repo   userrepository.UserRepository
	log    log.Logger
	tracer observability.Tracer
}

func NewUserService(repo userrepository.UserRepository, log log.Logger, tracer observability.Tracer) ports.UserServiceInterface {
	return &UserService{repo, log, tracer}
}

func (s *UserService) CreateUser(ctx context.Context, user domain.User) (string, *domain.DomainError) {

	ctx, span := s.tracer.Span(ctx, "UserService.CreateUser")
	defer span.End()

	user.Id = uuid.New().String()

	err := s.repo.Save(ctx, user)
	if err != nil {
		span.RecordError(err)
		s.log.Error("Failed to save user: ", err)
		return "", &domain.DomainError{
			Message: "error on create user",
			Code:    errorscode.ErrInternalError,
			Error:   err,
		}
	}
	return user.Id, nil
}

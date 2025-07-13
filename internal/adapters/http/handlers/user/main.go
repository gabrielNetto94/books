package userhandler

import (
	httputils "books/internal/adapters/http/http-utils"
	errorscode "books/internal/consts/errors-code"
	"books/internal/core/domain"
	"books/internal/infra/log"
	"books/internal/ports"
	"encoding/json"
	"net/http"
)

type UserHTTPHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserHandlers struct {
	service ports.UserServiceInterface
	log     log.Logger
}

// NewUserHandlers cria uma nova instância de UserHandlers.
func NewUserHandlers(service ports.UserServiceInterface, log log.Logger) UserHTTPHandler {
	return &UserHandlers{service, log}
}

func (u UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	u.log.Info("CreateUser called")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		httputils.HandleError(w, domain.DomainError{
			Message: "Invalid request",
			Error:   err,
			Code:    errorscode.ErrInvalidInput,
		})
		return
	}

	userId, err := u.service.CreateUser(r.Context(), user)
	if err != nil {
		u.log.Error("Error creating user: ", err.Error.Error())
		httputils.HandleError(w, *err)
		return
	}

	httputils.JsonResponse(w, http.StatusCreated, map[string]string{"userId": userId})
}

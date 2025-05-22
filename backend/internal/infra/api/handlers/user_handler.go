package handlers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type UserHandler struct {
	userGetAllUseCase usecases.IUserGetAllUseCase
	userCreateUseCase usecases.IUserCreateUseCase
}

func NewUserHandler(userGetAllUseCase usecases.IUserGetAllUseCase, userCreateUseCase usecases.IUserCreateUseCase) *UserHandler {
	return &UserHandler{
		userGetAllUseCase: userGetAllUseCase,
		userCreateUseCase: userCreateUseCase,
	}
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userGetAllUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	}

	response := mappers.UserGetListOutputToResponse(users)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreateInput dto.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userCreateInput); err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	user, err := h.userCreateUseCase.Execute(r.Context(), mappers.UserCreateInputToRequest(userCreateInput))
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	response := mappers.UserGetOutputToResponse(user)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

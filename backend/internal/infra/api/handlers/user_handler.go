package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type UserHandler struct {
	userGetAllUseCase     usecases.IUserGetAllUseCase
	userCreateUseCase     usecases.IUserCreateUseCase
	userGetByEmailUseCase usecases.IUserGetByEmailUseCase
}

func NewUserHandler(userGetAllUseCase usecases.IUserGetAllUseCase, userCreateUseCase usecases.IUserCreateUseCase, userGetByEmailUseCase usecases.IUserGetByEmailUseCase) *UserHandler {
	return &UserHandler{
		userGetAllUseCase:     userGetAllUseCase,
		userCreateUseCase:     userCreateUseCase,
		userGetByEmailUseCase: userGetByEmailUseCase,
	}
}

func (h *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userGetAllUseCase.Execute(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to get users", err)
		return
	}

	response := mappers.UserGetListOutputToResponse(users)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userCreateInput dto.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&userCreateInput); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to decode request body", err)
		return
	}

	user, err := h.userCreateUseCase.Execute(r.Context(), mappers.UserCreateInputToRequest(userCreateInput))
	if err != nil {
		if strings.Contains(err.Error(), "email already in use") {
			http.Error(w, `{"message": "Email đã được sử dụng."}`, http.StatusBadRequest)
			return
		}
		http.Error(w, `{"message": "Tạo người dùng không thành công: `+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := mappers.UserGetOutputToResponse(user)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		utils.WriteError(w, http.StatusBadRequest, "email is required", nil)
		return
	}

	user, err := h.userGetByEmailUseCase.Execute(r.Context(), email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to get user by email", err)
		return
	}

	response := mappers.UserGetOutputToResponse(*user)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}

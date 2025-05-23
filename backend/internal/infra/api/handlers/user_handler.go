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
	userUpdateUseCase     usecases.IUserUpdateUseCase
	userGetUseCase        usecases.IUserGetUseCase
}

func NewUserHandler(userGetAllUseCase usecases.IUserGetAllUseCase, userCreateUseCase usecases.IUserCreateUseCase, userGetByEmailUseCase usecases.IUserGetByEmailUseCase, userUpdateUseCase usecases.IUserUpdateUseCase, userUserGetUseCase usecases.IUserGetUseCase) *UserHandler {
	return &UserHandler{
		userGetAllUseCase:     userGetAllUseCase,
		userCreateUseCase:     userCreateUseCase,
		userGetByEmailUseCase: userGetByEmailUseCase,
		userUpdateUseCase:     userUpdateUseCase,
		userGetUseCase:        userUserGetUseCase,
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

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := utils.UserIdFromContext(r.Context())
	if id == 0 {
		http.Error(w, `{"message": "Extract userID from token failed"}`, http.StatusBadRequest)
		return
	}
	// Parse request body
	var userUpdateInput dto.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&userUpdateInput); err != nil {
		http.Error(w, `{"message": "Invalid customer data. Please check the input fields."}`, http.StatusBadRequest)
		return
	}

	// Gọi usecase để xử lý logic
	updatedUser, err := h.userUpdateUseCase.Execute(r.Context(), id, mappers.UserUpdateInputToRequest(userUpdateInput))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	// Trả về response thành công
	response := mappers.UserUpdateOutputToResponse(updatedUser)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to encode response", err)
		return
	}
}

func (h *UserHandler) GetUserByToken(w http.ResponseWriter, r *http.Request) {
	// Lấy `id` từ query string
	tokenUserID := utils.UserIdFromContext(r.Context())
	if tokenUserID == 0 {
		http.Error(w, `{"message": "Extract userID from token failed"}`, http.StatusBadRequest)
		return
	}

	user, err := h.userGetUseCase.Execute(r.Context(), tokenUserID)
	if err != nil {
		http.Error(w, `{"message": "Failed to get user by token"}`, http.StatusInternalServerError)
		return
	}
	response := mappers.UserGetOutputToResponseByToken(user)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "failed to encode response", err)
		return
	}
}

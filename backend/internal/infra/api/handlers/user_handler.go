package handlers

import (
	"encoding/json"
	"net/http"

	usecases "github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type UserHandler struct {
	userGetByEmailUseCase usecases.IUserGetByEmailUseCase
}

func NewUserHandler(userGetByEmailUseCase usecases.IUserGetByEmailUseCase) *UserHandler {
	return &UserHandler{
		userGetByEmailUseCase: userGetByEmailUseCase,
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

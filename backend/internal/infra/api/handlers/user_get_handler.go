package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases"
)

type UserGetHandler struct {
	userGetHandler usecases.IUserGetAllUseCase
}

func NewUserGetHandler(userGetHandler usecases.IUserGetAllUseCase) *UserGetHandler {
	return &UserGetHandler{
		userGetHandler: userGetHandler,
	}
}

func (h *UserGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.userGetHandler.Execute(r.Context())
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		return
	}

	response := []usecases.UserGetOutput{}
	for _, user := range users {
		response = append(response, usecases.UserGetOutput{
			UserID:   user.UserID,
			Name:     user.Username,
			Password: user.Password,
			Role:     user.Role,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

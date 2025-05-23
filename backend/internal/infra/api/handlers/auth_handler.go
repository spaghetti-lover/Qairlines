package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type AuthHandler struct {
	loginUseCase auth.ILoginUseCase
}

func NewAuthHandler(loginUseCase auth.ILoginUseCase) *AuthHandler {
	return &AuthHandler{loginUseCase: loginUseCase}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input auth.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	output, err := h.loginUseCase.Execute(r.Context(), input)
	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			http.Error(w, fmt.Sprintf(`{"message": "%s"}`, appErr.Message), http.StatusUnauthorized)
			return
		}
		http.Error(w, `{"message": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	response := mappers.LoginOutputToResponse(*output)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

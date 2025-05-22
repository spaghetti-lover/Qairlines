package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
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
		if errors.Is(err, errors.New("ERR_INVALID_CREDENTIALS")) {
			message := utils.GetErrorMessage("ERR_INVALID_CREDENTIALS", "vi")
			utils.WriteError(w, http.StatusUnauthorized, `{"message": "`+message+`"}`, nil)
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "", err)
		return
	}

	response := mappers.LoginOutputToResponse(*output)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

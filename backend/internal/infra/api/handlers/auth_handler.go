package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"
	"github.com/spaghetti-lover/qairlines/pkg/utils"
)

type PasswordChangeRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type AuthHandler struct {
	loginUseCase          auth.ILoginUseCase
	changePasswordUseCase auth.IChangePasswordUseCase
}

func NewAuthHandler(loginUseCase auth.ILoginUseCase, changePasswordUseCase auth.IChangePasswordUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase:          loginUseCase,
		changePasswordUseCase: changePasswordUseCase,
	}
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

func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Lấy userID từ URL
	vars := mux.Vars(r)
	userID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	var req PasswordChangeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Gọi usecase để xử lý logic
	err = h.changePasswordUseCase.Execute(r.Context(), auth.ChangePasswordInput{
		UserID:      userID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	// Trả về response thành công
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully."})
}

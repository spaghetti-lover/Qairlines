package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"
	"github.com/spaghetti-lover/qairlines/pkg/token"
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

	if input.Email == "" || input.Password == "" {
		http.Error(w, `{"message": "Email and password are required."}`, http.StatusBadRequest)
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

// ChangePassword handles password change requests
func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {

	// Lấy token payload từ context
	authPayload, ok := r.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		http.Error(w, `{"message": "Authentication failed. Invalid token1."}`, http.StatusUnauthorized)
		return
	}

	// Parse request
	var request dto.ChangePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"message": "Invalid request format."}`, http.StatusBadRequest)
		return
	}

	// Validate request
	if request.Email == "" || request.OldPassword == "" || request.NewPassword == "" {
		http.Error(w, `{"message": "Email, old password, and new password are required."}`, http.StatusBadRequest)
		return
	}

	// Convert request to use case input
	input := mappers.ChangePasswordRequestToInput(request)

	// Call use case
	err := h.changePasswordUseCase.Execute(r.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrOldPasswordIncorrect):
			http.Error(w, `{"message": "Old password is incorrect."}`, http.StatusBadRequest)
		case errors.Is(err, auth.ErrPasswordValidationFailed):
			http.Error(w, `{"message": "New password does not meet the required criteria."}`, http.StatusUnprocessableEntity)
		case errors.Is(err, auth.ErrUserNotFound):
			http.Error(w, `{"message": "User not found with the provided email."}`, http.StatusNotFound)
		default:
			log.Printf("Error type: %T, Error value: %v", err, err)
			http.Error(w, `{"message": "An unexpected error occurred. Please try again later."}`, http.StatusInternalServerError)
		}
		return
	}

	// Return success response
	response := dto.ChangePasswordResponse{
		Message: "Password changed successfully.",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

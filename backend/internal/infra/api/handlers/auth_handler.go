package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghetti-lover/qairlines/internal/domain/usecases/auth"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/dto"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/middleware"
	appErrors "github.com/spaghetti-lover/qairlines/pkg/errors"
	"github.com/spaghetti-lover/qairlines/pkg/token"
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

func (h *AuthHandler) Login(c *gin.Context) {
	var input auth.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}

	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email and password are required."})
		return
	}

	output, err := h.loginUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		if appErr, ok := err.(*appErrors.AppError); ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": appErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	response := mappers.LoginOutputToResponse(*output)
	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	// Lấy token payload từ context
	authPayload, ok := c.Request.Context().Value(middleware.AuthorizationPayloadKey).(*token.Payload)
	if !ok || authPayload == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed. Invalid token."})
		return
	}

	// Parse request
	var request dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}

	// Validate request
	if request.Email == "" || request.OldPassword == "" || request.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email, old password, and new password are required."})
		return
	}

	// Convert request to use case input
	input := mappers.ChangePasswordRequestToInput(request)

	// Call use case
	err := h.changePasswordUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrOldPasswordIncorrect):
			c.JSON(http.StatusBadRequest, gin.H{"message": "Old password is incorrect."})
		case errors.Is(err, auth.ErrPasswordValidationFailed):
			c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "New password does not meet the required criteria."})
		case errors.Is(err, auth.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found with the provided email."})
		default:
			log.Printf("Error type: %T, Error value: %v", err, err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred. Please try again later."})
		}
		return
	}

	// Return success response
	response := dto.ChangePasswordResponse{
		Message: "Password changed successfully.",
	}
	c.JSON(http.StatusOK, response)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/spaghetti-lover/qairlines/internal/domain/usecases/user"
	"github.com/spaghetti-lover/qairlines/internal/infra/api/mappers"
)

type UserHandler struct {
	userGetByEmailUseCase usecases.IUserGetByEmailUseCase
}

func NewUserHandler(userGetByEmailUseCase usecases.IUserGetByEmailUseCase) *UserHandler {
	return &UserHandler{
		userGetByEmailUseCase: userGetByEmailUseCase,
	}
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is required"})
		return
	}

	user, err := h.userGetByEmailUseCase.Execute(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user by email", "error": err.Error()})
		return
	}

	response := mappers.UserGetOutputToResponse(*user)
	c.JSON(http.StatusOK, response)
}

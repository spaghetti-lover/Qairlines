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

func (h *UserHandler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email is required"})
		return
	}

	user, err := h.userGetByEmailUseCase.Execute(ctx.Request.Context(), email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get user by email", "error": err.Error()})
		return
	}

	response := mappers.UserGetOutputToResponse(*user)
	ctx.JSON(http.StatusOK, response)
}

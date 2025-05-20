package dto

import "github.com/spaghetti-lover/qairlines/internal/domain/entities"

type UserGetResponse struct {
	UserID   int64             `json:"user_id"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Role     entities.UserRole `json:"role"`
}

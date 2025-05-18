package dto

type UserCreateRequest struct {
	Name  string `query:"name"`
	Email string `query:"email"`
}

type UserCreateResponse struct {
	UserID string `json:"user_id"`
}

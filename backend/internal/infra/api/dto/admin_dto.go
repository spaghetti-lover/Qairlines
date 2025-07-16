package dto

type CreateAdminRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateAdminResponse struct {
	Message string `json:"message"`
	Admin   struct {
		ID        int64  `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		CreatedAt struct {
			Seconds int64 `json:"seconds"`
		} `json:"createdAt"`
	} `json:"admin"`
}

type AdminResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	CreatedAt struct {
		Seconds int64 `json:"seconds"`
	} `json:"createdAt"`
}

type ListAdminsResponse struct {
	Message string          `json:"message"`
	Data    []AdminResponse `json:"data"`
}

type GetCurrentAdminResponse struct {
	Message string `json:"message"`
	Data    struct {
		UID       string `json:"uid"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	} `json:"data"`
}

type AdminUpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type AdminUpdateResponse struct {
	Message string `json:"message"`
	Data    struct {
		UID       string `json:"uid"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	} `json:"data"`
}

type ChangePasswordRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

type ListAdminsParams struct {
	Limit int `json:"limit" binding:"required,min=1,max=100" default:"10"`
	Page  int `json:"page" binding:"required,min=1" default:"1"`
}

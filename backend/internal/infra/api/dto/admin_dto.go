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
		ID        string `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		CreatedAt struct {
			Seconds int64 `json:"seconds"`
		} `json:"createdAt"`
	} `json:"admin"`
}

type AdminResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	CreatedAt struct {
		Seconds int64 `json:"seconds"`
	} `json:"createdAt"`
}

type GetAllAdminsResponse struct {
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

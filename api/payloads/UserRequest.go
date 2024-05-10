package payloads

type CreateRequest struct {
	UserName     string `json:"username" validate:"required,max=30,min=5" `
	Email        string `json:"email" validate:"required,email"`
	ActiveStatus bool   `json:"active_status" validate:"required"`
	UserPassword string `json:"user_password" validate:"required,max=100,min=5"`
}

type UserDetails struct {
	Role      int    `json:"role"`
	CompanyID string `json:"company_id"`
}

type LoginRequest struct {
	Username     string `json:"username" validate:"required"`
	UserPassword string `json:"user_password" validate:"required"`
	Client       string `json:"client" validate:"required"`
}

type LoginCredential struct {
	Client    string `json:"client"`
	IpAddress string `json:"ip_address"`
	Session   string `json:"session"`
}

type UserDetail struct {
	UserID    int    `json:"user_id"`
	Username  string `json:"username"`
	Authorize string `json:"authorized"`
	CompanyID string `json:"company_id"`
	Role      int    `json:"role"`
	IpAddress string `json:"ip_address"`
	Client    string `json:"client"`
}

type CurrentUserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

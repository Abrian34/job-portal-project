package payloads

type CreateRequest struct {
	UserName     string `json:"username" validate:"required"`
	UserPassword string `json:"user_password" validate:"required"`
	RoleId       int    `json:"role_id"`
}

type UserDetails struct {
	Role      int    `json:"role"`
	CompanyID string `json:"company_id"`
}

type LoginRequest struct {
	Username     string `json:"username" validate:"required"`
	UserPassword string `json:"user_password" validate:"required"`
}

type LoginCredential struct {
	Client    string `json:"client"`
	IpAddress string `json:"ip_address"`
	Session   string `json:"session"`
}

type UserDetail struct {
	UserId          int    `json:"user_id"`
	UserCode        string `json:"user_code"`
	UserDisplayName string `json:"user_display_name"`
	RoleId          int    `json:"role_id"`
	RoleName        string `json:"role_name"`
	UserName        string `json:"username"`
	UserPassword    string `json:"user_password" `
	ActiveStatus    bool   `json:"active_status"`
}

type CurrentUserResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type RoleResponse struct {
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

type LoginResponse struct {
	User  UserDetail `json:"user"`
	Token string     `json:"token"`
}

type RegisterResponse struct {
	UserID int `json:"user_id"`
}

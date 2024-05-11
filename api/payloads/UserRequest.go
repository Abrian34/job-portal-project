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
	RoleId      int                `json:"role_id"`
	RoleName    string             `json:"role_name"`
	Permissions []PermissionDetail `json:"permissions"`
}

type PermissionDetail struct {
	PermissionId   int    `json:"permission_id"`
	PermissionName string `json:"permission_name"`
}

type LoginResponse struct {
	User        UserDetail         `json:"user"`
	Permissions []PermissionDetail `json:"permissions"`
	Token       string             `json:"token"`
}

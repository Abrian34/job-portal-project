package entitypayloads

type RolePayload struct {
	RoleId          int    `json:"role_id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	ActiveStatus    bool   `json:"active_status"`
}

type RoleRequest struct {
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	ActiveStatus    bool   `json:"active_status"`
}

type RoleUpdate struct {
	RoleId          int    `json:"role_id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
}

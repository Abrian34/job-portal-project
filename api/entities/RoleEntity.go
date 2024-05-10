package entities

var CreateRoleTable = "role_tbl"

type Role struct {
	RoleId       int    `gorm:"column:role_id;size:30;not null;primaryKey" json:"role_id"`
	PermissionId int    `gorm:"column:permission_id;size:30;not null;primaryKey" json:"permission_id"`
	RoleCode     string `gorm:"column:role_code;size:50;not null" json:"role_code"`
	RoleName     string `gorm:"column:role_name;size:50;not null" json:"role_name"`
	ActiveStatus bool   `gorm:"column:active_status;size:50;not null" json:"active_status"`
}

func (*Role) TableName() string {
	return CreateRoleTable
}

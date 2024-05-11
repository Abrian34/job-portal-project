package entities

var CreateRolePermissionTable = "role_permission_table"

type RolePermission struct {
	RoleId       int    `gorm:"column:role_id;size:30;not null;" json:"role_id"`
	PermissionId string `gorm:"column:permission_id;size:50;not null" json:"permission_id"`
}

func (*RolePermission) TableName() string {
	return CreateRolePermissionTable
}

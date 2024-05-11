package entities

var CreatePermissionTable = "permission_tbl"

type Permission struct {
	PermissionId   int    `gorm:"column:permission_id;size:30;not null;primaryKey" json:"permission_id"`
	PermissionName string `gorm:"column:permission_name;size:50;not null" json:"permission_name"`
	Roles          []Role `gorm:"many2many:role_permission_table" json:"roles"`
	// RolePermission RolePermission
}

func (*Permission) TableName() string {
	return CreatePermissionTable
}

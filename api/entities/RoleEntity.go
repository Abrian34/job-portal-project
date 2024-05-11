package entities

var CreateRoleTable = "role_tbl"

type Role struct {
	RoleId      int          `gorm:"column:role_id;size:30;not null;primaryKey" json:"role_id"`
	RoleName    string       `gorm:"column:role_name;size:50;not null" json:"role_name"`
	Permissions []Permission `gorm:"many2many:role_permission_table" json:"permissions"`
	// RolePermission RolePermission
}

func (*Role) TableName() string {
	return CreateRoleTable
}

package entities

var CreateRoleTable = "role_tbl"

type Role struct {
	RoleId          int    `gorm:"column:role_id;size:30;not null;primaryKey" json:"role_id"`
	RoleName        string `gorm:"column:role_name;size:50;not null" json:"role_name"`
	RoleDescription string `gorm:"column:role_description;size:50;not null" json:"role_description"`
	ActiveStatus    bool   `gorm:"column:active_status;not null;default:true" json:"active_status"`
	User            User   `gorm:"foreignKey:RoleId" json:"user"`
}

func (*Role) TableName() string {
	return CreateRoleTable
}

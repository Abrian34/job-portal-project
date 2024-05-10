package entities

var CreateUserTable = "user_tbl"

type User struct {
	UserId          int    `gorm:"column:user_id;size:30;not null;primaryKey" json:"user_id"`
	UserCode        string `gorm:"column:user_code;size:50;not null" json:"user_code"`
	RoleId          int    `gorm:"column:role_id;size:30;not null" json:"role_id"`
	UserDisplayName string `gorm:"column:user_display_name;size:50;not null" json:"user_display_name"`
	UserName        string `gorm:"column:username;size:50;not null" json:"username"`
	UserPassword    string `gorm:"column:user_password;size:50;not null" json:"user_password"`
	ActiveStatus    bool   `gorm:"column:active_status;size:50;not null" json:"active_status"`
}

func (*User) TableName() string {
	return CreateUserTable
}

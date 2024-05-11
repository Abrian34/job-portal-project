package userrepo

import (
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"

	"gorm.io/gorm"
)

type UserRepository interface {
	CheckUserExists(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	FindUser(*gorm.DB, string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	ViewUser(*gorm.DB) ([]entities.User, *exceptions.BaseErrorResponse)
	GetCurrentUser(*gorm.DB, int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse)
	GetByID(*gorm.DB, int) (entities.User, *exceptions.BaseErrorResponse)
	GetByUsername(*gorm.DB, string) (payloads.UserDetail, *exceptions.BaseErrorResponse)
	// GetEmails(*gorm.DB, []int) ([]string, *exceptions.BaseErrorResponse)
	GetUsernameByUserID(*gorm.DB, int) (string, *exceptions.BaseErrorResponse)
	GetUserIDByUsername(*gorm.DB, string) (int, *exceptions.BaseErrorResponse)
	// GetByEmail(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	GetRoleByCompanyAndUserID(*gorm.DB, int, int) (int, *exceptions.BaseErrorResponse)
	GetUserDetailByUsername(*gorm.DB, string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	Create(*gorm.DB, payloads.CreateRequest, int) (int, *exceptions.BaseErrorResponse)
	Update(*gorm.DB, payloads.CreateRequest, int) (bool, *exceptions.BaseErrorResponse)
	Delete(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	GetAllRole(*gorm.DB) ([]payloads.RoleResponse, *exceptions.BaseErrorResponse)
	GetPermissionsByRoleID(*gorm.DB, int) []payloads.PermissionDetail
	GetRoleById(*gorm.DB, int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
	GetRoleWithPermissions(*gorm.DB, int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
}

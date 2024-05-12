package userservices

import (
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
)

type UserService interface {
	GetCurrentUser(int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse)
	FindUser(string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	// CheckUserExists(string) (bool, *exceptions.BaseErrorResponse)
	ViewUser() ([]entities.User, *exceptions.BaseErrorResponse)
	// GetEmails([]int) ([]string, *exceptions.BaseErrorResponse)
	GetByID(int) (entities.User, *exceptions.BaseErrorResponse)
	GetUsernameByUserID(int) (string, *exceptions.BaseErrorResponse)
	GetUserIDByUsername(string) (int, *exceptions.BaseErrorResponse)
	GetUser(username string) (payloads.UserDetail, *exceptions.BaseErrorResponse)
	GetUserDetailByUsername(string) (payloads.UserDetails, *exceptions.BaseErrorResponse)
	// UpdateUser(payloads.CreateRequest, int) (bool, *exceptions.BaseErrorResponse)
	DeleteUser(int) (bool, *exceptions.BaseErrorResponse)
	GetAllRole() ([]payloads.RoleResponse, *exceptions.BaseErrorResponse)
	// GetPermissionsByRoleID(int) []payloads.PermissionDetail
	GetRoleById(int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
	// GetRoleWithPermissions(int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
}

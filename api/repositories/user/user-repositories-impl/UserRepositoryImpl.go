package userrepoimpl

import (
	"errors"
	"fmt"
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
	userrepo "job-portal-project/api/repositories/user"
	"job-portal-project/api/utils/constant"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

// GetCurrentUser implements userrepo.UserRepository.
func (*UserRepositoryImpl) GetCurrentUser(tx *gorm.DB, userID int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse) {
	var user entities.User
	var userResponse payloads.CurrentUserResponse
	err := tx.Model(user).
		Select(
			"user_id",
			"username",
		).
		Where(entities.User{
			UserId: userID,
		}).
		Scan(&userResponse).
		Error

	if err != nil {
		return userResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return userResponse, nil
}

// GetRoleByCompanyAndUserID implements userrepo.UserRepository.
func (*UserRepositoryImpl) GetRoleByCompanyAndUserID(tx *gorm.DB, companyID int, userID int) (int, *exceptions.BaseErrorResponse) {
	var user entities.User
	var role int
	err := tx.Model(user).
		Select(
			"MenuUserAccess__MenuAccess.role_id",
		).
		InnerJoins("MenuUserAccess",
			tx.Select("1"),
		).
		InnerJoins("MenuUserAccess.MenuAccess",
			tx.Select("1").
				Where("MenuUserAccess__MenuAccess.company_id = ?", companyID),
		).
		Where(entities.User{
			UserId: userID,
		}).
		Scan(&role).
		Error

	if err != nil {
		return role, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if role == 0 {
		return role, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusForbidden,
			Message:    fmt.Sprintf("%s %s", constant.PermissionError, " of any menus"),
			Err:        errors.New(constant.PermissionError),
		}
	}

	return role, nil
}

func NewUserRepository() userrepo.UserRepository {
	return &UserRepositoryImpl{}
}

// func (*UserRepositoryImpl) CheckUserExists(tx *gorm.DB, username string) (bool, *exceptions.BaseErrorResponse) {
// 	var user entities.User
// 	var exists bool
// 	err := tx.Model(user).
// 		Select(
// 			"count(id)",
// 		).
// 		Where(entities.User{UserName: username}).
// 		Find(&exists).
// 		Error
// 	if exists {
// 		return exists, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusConflict,
// 			Err:        err,
// 		}
// 	}
// 	if err != nil {
// 		return exists, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	return exists, nil
// }

func (*UserRepositoryImpl) FindUser(tx *gorm.DB, username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	var user entities.User
	var userDetail payloads.UserDetails
	err := tx.Model(user).
		Select(
			"role",
			"company_id",
		).
		Where(entities.User{UserName: username}).
		Scan(&userDetail).
		Error

	if err != nil {

		return userDetail, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return userDetail, nil
}

func (*UserRepositoryImpl) ViewUser(tx *gorm.DB) ([]entities.User, *exceptions.BaseErrorResponse) {
	var user []entities.User
	row, err := tx.Model(user).Scan(&user).Rows()

	if err != nil {

		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	var users []entities.User
	for row.Next() {
		var user entities.User

		err := row.Scan(
			&user.UserId,
			&user.UserName,
			&user.UserPassword)

		if err != nil {

			return users, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
		users = append(users, user)
	}

	return users, nil
}

func (*UserRepositoryImpl) GetByUsername(tx *gorm.DB, username string) (payloads.UserDetail, *exceptions.BaseErrorResponse) {
	var user entities.User
	response := payloads.UserDetail{}

	row, err := tx.
		Model(user).
		Where(entities.User{
			UserName: username,
		}).
		Scan(&response).
		Rows()

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.UserName == "" {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	defer row.Close()

	return response, nil
}

// func (*UserRepositoryImpl) GetByEmail(tx *gorm.DB, email string) (bool, *exceptions.BaseErrorResponse) {
// 	var user entities.User

// 	row, err := tx.
// 		Model(user).
// 		Where(entities.User{
// 			Email: email,
// 		}).
// 		Scan(&user).
// 		Rows()

// 	if err != nil {

// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}

// 	if user.Email == "" {

// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// // GetEmails implements repositories.UserRepository.
// func (*UserRepositoryImpl) GetEmails(tx *gorm.DB, userIDs []int) ([]string, *exceptions.BaseErrorResponse) {
// 	var email []string

// 	row, err := tx.
// 		Model(entities.User{}).
// 		Select("email").
// 		Where("id in (?)", userIDs).
// 		Scan(email).
// 		Rows()

// 	if err != nil {

// 		return email, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

//		return email, nil
//	}
func (*UserRepositoryImpl) GetUserDetailByUsername(tx *gorm.DB, username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	var user payloads.UserDetails

	row, err := tx.
		Model(entities.User{}).
		Where(entities.User{UserName: username}).
		Scan(&user).
		Rows()

	if err != nil {

		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return user, nil
}

func (*UserRepositoryImpl) GetByID(tx *gorm.DB, userID int) (entities.User, *exceptions.BaseErrorResponse) {
	var user entities.User

	rows, err := tx.Model(user).
		Where(entities.User{
			UserId: userID,
		}).Scan(&user).
		Rows()

	if err != nil {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if user.UserName == "" {
		return user, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}
	defer rows.Close()

	return user, nil
}

// GetUsernameByUserID implements repositories.UserRepository.
func (*UserRepositoryImpl) GetUserIDByUsername(tx *gorm.DB, username string) (int, *exceptions.BaseErrorResponse) {
	var user entities.User

	err := tx.
		Select("id").
		Where(entities.User{
			UserName: username,
		}).
		Find(&user).
		Error

	if err != nil {
		logrus.Info(err)

		return 0, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.UserId == 0 {
		return 0, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return user.UserId, nil
}

// GetUsernameByUserID implements repositories.UserRepository.
func (*UserRepositoryImpl) GetUsernameByUserID(tx *gorm.DB, userID int) (string, *exceptions.BaseErrorResponse) {
	var user entities.User

	err := tx.
		Select("username").
		Where(entities.User{
			UserId: userID,
		}).
		Find(&user).
		Error

	if err != nil {

		return user.UserName, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if user.UserName == "" {
		return user.UserName, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return user.UserName, nil
}

func (*UserRepositoryImpl) Create(tx *gorm.DB, userReq payloads.CreateRequest, roleID int) (int, *exceptions.BaseErrorResponse) {
	user := entities.User{
		UserName:     userReq.UserName,
		UserPassword: userReq.UserPassword,
		ActiveStatus: true,
		RoleId:       roleID,
	}
	err := tx.
		Create(&user).
		Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return 0, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusConflict,
				Err:        err,
			}
		} else {
			return 0, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Err:        err,
			}
		}
	}

	return user.UserId, nil
}

// func (*UserRepositoryImpl) Update(tx *gorm.DB, userReq payloads.CreateRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		UserName:     userReq.UserName,
// 		UserPassword: userReq.UserPassword,
// 		ActiveStatus: userReq.ActiveStatus,
// 		// OtpEnabled:   true,
// 	}

// 	err := tx.
// 		Where(userID).
// 		Updates(&user).
// 		Error
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}

//		return true, nil
//	}
func (*UserRepositoryImpl) Delete(tx *gorm.DB, userID int) (bool, *exceptions.BaseErrorResponse) {

	err := tx.
		Where(userID).
		Delete(userID).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}

func (r *UserRepositoryImpl) GetAllRole(tx *gorm.DB) ([]payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	var roles []entities.Role
	err := tx.Model(&entities.Role{}).Find(&roles).Error
	if err != nil {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	var roleResponses []payloads.RoleResponse
	for _, role := range roles {
		roleResponse := payloads.RoleResponse{
			RoleId:   role.RoleId,
			RoleName: role.RoleName,
		}
		roleResponses = append(roleResponses, roleResponse)
	}

	return roleResponses, nil
}

// func (r *UserRepositoryImpl) GetPermissionsByRoleID(tx *gorm.DB, roleID int) []payloads.PermissionDetail {
// 	var permissions []entities.Permission
// 	err := tx.Model(&entities.Permission{}).Where("role_id = ?", roleID).Find(&permissions).Error
// 	if err != nil {
// 		return []payloads.PermissionDetail{}
// 	}

// 	var permissionResponses []payloads.PermissionDetail
// 	for _, permission := range permissions {
// 		permissionResponse := payloads.PermissionDetail{
// 			PermissionId:   permission.PermissionId,
// 			PermissionName: permission.PermissionName,
// 		}
// 		permissionResponses = append(permissionResponses, permissionResponse)
// 	}

// 	return permissionResponses
// }

func (r *UserRepositoryImpl) GetRoleById(tx *gorm.DB, RoleId int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	entity := entities.Role{}
	response := payloads.RoleResponse{}

	rows, err := tx.Model(&entity).
		Where(entities.Role{
			RoleId: RoleId,
		}).
		First(&response).
		Rows()

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	defer rows.Close()

	return response, nil
}

// func (repo *UserRepositoryImpl) GetRoleWithPermissions(tx *gorm.DB, roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
// 	var role entities.Role

// 	// Retrieve role by ID
// 	if err := tx.Preload("Permissions").First(&role, roleID).Error; err != nil {
// 		return payloads.RoleResponse{}, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}

// 	// Convert the entity to a payload struct
// 	response := payloads.RoleResponse{
// 		RoleId:   role.RoleId,
// 		RoleName: role.RoleName,
// 	}

// 	// Convert each permission to PermissionDetail
// 	for _, permission := range role.Permissions {
// 		response.Permissions = append(response.Permissions, payloads.PermissionDetail{
// 			PermissionId:   permission.PermissionId,
// 			PermissionName: permission.PermissionName,
// 		})
// 	}

// 	return response, nil
// }

package userserviceimpl

import (
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/helper"
	"job-portal-project/api/payloads"
	userrepo "job-portal-project/api/repositories/user"
	userservices "job-portal-project/api/services/user"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository userrepo.UserRepository
	DB             *gorm.DB
}

// GetCurrentUser implements userservices.UserService.
func (service *UserServiceImpl) GetCurrentUser(userID int) (payloads.CurrentUserResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.GetCurrentUser(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func NewUserService(userRepository userrepo.UserRepository, db *gorm.DB) userservices.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

// func (service *UserServiceImpl) CheckUserExists(username string) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)

// 	get, err := service.UserRepository.CheckUserExists(tx, username)

// 	if err != nil {
// 		return false, err
// 	}

// 	return get, nil
// }

func (service *UserServiceImpl) FindUser(username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.FindUser(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetEmails implements services.UserService.
// func (service *UserServiceImpl) GetEmails(users []int) ([]string, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	get, err := service.UserRepository.GetEmails(tx, users)

// 	if err != nil {
// 		return get, err
// 	}

// 	return get, nil
// }

// // GetByEmail implements services.UserService.
// func (service *UserServiceImpl) GetByEmail(email string) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	get, err := service.UserRepository.GetByEmail(tx, email)

// 	if err != nil {
// 		return get, err
// 	}

// 	return get, nil
// }

// GetUserIDByUsername implements services.UserService.
func (service *UserServiceImpl) GetUserIDByUsername(username string) (int, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUserIDByUsername(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUsernameByUserID implements services.UserService.
func (service *UserServiceImpl) GetUsernameByUserID(userID int) (string, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUsernameByUserID(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetByID implements services.UserService.
func (service *UserServiceImpl) GetByID(userID int) (entities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.UserRepository.GetByID(tx, userID)

	if err != nil {
		return get, err
	}

	return get, nil
}

// GetUser implements services.UserService.
func (service *UserServiceImpl) GetUser(username string) (payloads.UserDetail, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetByUsername(tx, username)
	if err != nil {
		return get, err
	}
	return get, nil
}

// GetUserDetailByUsername implements services.UserService.
func (service *UserServiceImpl) GetUserDetailByUsername(username string) (payloads.UserDetails, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.GetUserDetailByUsername(tx, username)

	if err != nil {
		return get, err
	}

	return get, nil
}

// ViewUser implements services.UserService.
func (service *UserServiceImpl) ViewUser() ([]entities.User, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := service.UserRepository.ViewUser(tx)

	if err != nil {
		return get, err
	}

	return get, nil
}

// UpdateUser implements services.UserService.
// func (service *UserServiceImpl) UpdateUser(userReq payloads.CreateRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	update, err := service.UserRepository.Update(tx, userReq, userID)

// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

// DeleteUser implements services.UserService.
func (service *UserServiceImpl) DeleteUser(userID int) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	deleteUser, err := service.UserRepository.Delete(tx, userID)

	if err != nil {
		return deleteUser, err
	}

	return deleteUser, nil
}

func (service *UserServiceImpl) GetAllRole() ([]payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := service.UserRepository.GetAllRole(tx)

	if err != nil {
		return result, err
	}

	return result, nil
}

// func (service *UserServiceImpl) GetPermissionsByRoleID(roleID int) []payloads.PermissionDetail {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	result := service.UserRepository.GetPermissionsByRoleID(tx, roleID)

// 	return result
// }

func (service *UserServiceImpl) GetRoleById(roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := service.UserRepository.GetRoleById(tx, roleID)
	if err != nil {
		return result, err
	}

	return result, nil
}

// func (service *UserServiceImpl) GetRoleWithPermissions(roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	result, err := service.UserRepository.GetRoleWithPermissions(tx, roleID)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

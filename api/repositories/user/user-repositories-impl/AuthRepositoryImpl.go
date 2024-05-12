package userrepoimpl

import (
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
	userrepo "job-portal-project/api/repositories/user"
	"net/http"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() userrepo.AuthRepository {
	return &AuthRepositoryImpl{}
}

// CheckPasswordResetTime implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) CheckPasswordResetTime(tx *gorm.DB, tokenReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	var exists bool
// 	err := tx.Model(entities.User{}).
// 		Select("count(company_id)").
// 		Where(
// 			"password_reset_token = ? AND password_reset_at > ?",
// 			tokenReq.PasswordResetToken, tokenReq.PasswordResetAt,
// 		).
// 		Find(&exists).
// 		Error

// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	if !exists {
// 		return exists, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Err:        errors.New("Invalid Token"),
// 		}
// 	}
// 	return exists, nil
// }

// UpdateCredential implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) UpdateCredential(tx *gorm.DB, loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		IpAddress: loginReq.IpAddress,
// 		LastLogin: time.Now(),
// 		ID:        userID,
// 	}

// 	err := tx.
// 		Where(entities.User{ID: userID}).
// 		Updates(&user).Error

// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	return true, nil
// }

// UpdatePassword implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) UpdatePassword(tx *gorm.DB, password string, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		Password: password,
// 	}
// 	row, err := tx.
// 		Where(entities.User{ID: userID}).
// 		Updates(&user).
// 		Rows()
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// UpdatePasswordByToken implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) UpdatePasswordByToken(tx *gorm.DB, passReq payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		Password: passReq.Password,
// 	}
// 	row, err := tx.
// 		Where(entities.User{PasswordResetToken: passReq.PasswordResetToken}).
// 		Updates(&user).
// 		Rows()
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// ResetPassword implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) ResetPassword(tx *gorm.DB, updateReq payloads.ResetPasswordRequest) (bool, *exceptions.BaseErrorResponse) {
// 	var user entities.User
// 	var nullString *string
// 	var nullTime *time.Time
// 	row, err := tx.
// 		Model(&user).
// 		Where(entities.User{PasswordResetToken: updateReq.PasswordResetToken}).
// 		Updates(map[string]interface{}{
// 			"password_reset_token": nullString,
// 			"password_reset_at":    nullTime,
// 		}).
// 		Rows()
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// UpdatePasswordTokenByEmail implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) UpdatePasswordTokenByEmail(tx *gorm.DB, emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		PasswordResetToken: emailReq.PasswordResetToken,
// 		PasswordResetAt:    emailReq.PasswordResetAt,
// 	}

// 	row, err := tx.
// 		Where(entities.User{Email: emailReq.Email}).
// 		Updates(&user).
// 		Rows()

// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// UpdateUserOTP implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) UpdateUserOTP(tx *gorm.DB, otpReq payloads.OTPRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		OtpVerified: otpReq.OtpVerified,
// 		OtpEnabled:  otpReq.OtpEnabled,
// 	}

// 	row, err := tx.
// 		Where(entities.User{ID: userID}).
// 		Updates(&user).
// 		Rows()
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// GenerateOTP implements repositories.AuthRepository.
// func (*AuthRepositoryImpl) GenerateOTP(tx *gorm.DB, userReq payloads.SecretUrlRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	user := entities.User{
// 		OtpSecret:  userReq.Secret,
// 		OtpAuthUrl: userReq.Url,
// 	}

// 	row, err := tx.
// 		Where(entities.User{ID: userID}).
// 		Updates(&user).
// 		Rows()
// 	if err != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusInternalServerError,
// 			Err:        err,
// 		}
// 	}
// 	defer row.Close()

// 	return true, nil
// }

// func (repo *AuthRepositoryImpl) GetRoleWithPermissions(tx *gorm.DB, roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
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

func (repo *AuthRepositoryImpl) GetRoleByUserID(tx *gorm.DB, userID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	var user entities.User
	if err := tx.First(&user, userID).Error; err != nil {
		return payloads.RoleResponse{}, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	var role entities.Role
	if err := tx.First(&role, user.RoleId).Error; err != nil {
		return payloads.RoleResponse{}, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	// Convert the entities.Role to payloads.RoleResponse
	roleResponse := payloads.RoleResponse{
		RoleId:   role.RoleId,
		RoleName: role.RoleName,
	}

	return roleResponse, nil
}

func (*AuthRepositoryImpl) CheckUserExists(tx *gorm.DB, username string) (bool, *exceptions.BaseErrorResponse) {
	var user entities.User
	var exists bool
	err := tx.Model(user).
		Select(
			"count(id)",
		).
		Where(entities.User{UserName: username}).
		Find(&exists).
		Error
	if exists {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusConflict,
			Err:        err,
		}
	}
	if err != nil {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return exists, nil
}

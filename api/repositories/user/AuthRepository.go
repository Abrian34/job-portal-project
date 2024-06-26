package userrepo

import (
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"

	"gorm.io/gorm"
)

type AuthRepository interface {
	// CheckPasswordResetTime(*gorm.DB, payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	// GenerateOTP(*gorm.DB, payloads.SecretUrlRequest, int) (bool, *exceptions.BaseErrorResponse)
	// UpdateUserOTP(*gorm.DB, payloads.OTPRequest, int) (bool, *exceptions.BaseErrorResponse)
	// UpdatePassword(*gorm.DB, string, int) (bool, *exceptions.BaseErrorResponse)
	// UpdatePasswordByToken(*gorm.DB, payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse)
	// UpdatePasswordTokenByEmail(*gorm.DB, payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	// ResetPassword(*gorm.DB, payloads.ResetPasswordRequest) (bool, *exceptions.BaseErrorResponse)
	// UpdateCredential(*gorm.DB, payloads.LoginCredential, int) (bool, *exceptions.BaseErrorResponse)

	// GetRoleWithPermissions(tx *gorm.DB, roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
	CheckUserExists(*gorm.DB, string) (bool, *exceptions.BaseErrorResponse)
	GetRoleByUserID(*gorm.DB, int) (payloads.RoleResponse, *exceptions.BaseErrorResponse)
}

package userrepoimpl

import (
	"errors"
	entities "job-portal-project/api/entities"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
	userrepo "job-portal-project/api/repositories/user"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() userrepo.AuthRepository {
	return &AuthRepositoryImpl{}
}

// CheckPasswordResetTime implements repositories.AuthRepository.
func (*AuthRepositoryImpl) CheckPasswordResetTime(tx *gorm.DB, tokenReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	var exists bool
	err := tx.Model(entities.User{}).
		Select("count(company_id)").
		Where(
			"password_reset_token = ? AND password_reset_at > ?",
			tokenReq.PasswordResetToken, tokenReq.PasswordResetAt,
		).
		Find(&exists).
		Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	if !exists {
		return exists, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("Invalid Token"),
		}
	}
	return exists, nil
}

// UpdateCredential implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdateCredential(tx *gorm.DB, loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		IpAddress: loginReq.IpAddress,
		LastLogin: time.Now(),
		ID:        userID,
	}

	err := tx.
		Where(entities.User{ID: userID}).
		Updates(&user).Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	return true, nil
}

// UpdatePassword implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePassword(tx *gorm.DB, password string, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		Password: password,
	}
	row, err := tx.
		Where(entities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdatePasswordByToken implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePasswordByToken(tx *gorm.DB, passReq payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		Password: passReq.Password,
	}
	row, err := tx.
		Where(entities.User{PasswordResetToken: passReq.PasswordResetToken}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// ResetPassword implements repositories.AuthRepository.
func (*AuthRepositoryImpl) ResetPassword(tx *gorm.DB, updateReq payloads.ResetPasswordRequest) (bool, *exceptions.BaseErrorResponse) {
	var user entities.User
	var nullString *string
	var nullTime *time.Time
	row, err := tx.
		Model(&user).
		Where(entities.User{PasswordResetToken: updateReq.PasswordResetToken}).
		Updates(map[string]interface{}{
			"password_reset_token": nullString,
			"password_reset_at":    nullTime,
		}).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdatePasswordTokenByEmail implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdatePasswordTokenByEmail(tx *gorm.DB, emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		PasswordResetToken: emailReq.PasswordResetToken,
		PasswordResetAt:    emailReq.PasswordResetAt,
	}

	row, err := tx.
		Where(entities.User{Email: emailReq.Email}).
		Updates(&user).
		Rows()

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// UpdateUserOTP implements repositories.AuthRepository.
func (*AuthRepositoryImpl) UpdateUserOTP(tx *gorm.DB, otpReq payloads.OTPRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		OtpVerified: otpReq.OtpVerified,
		OtpEnabled:  otpReq.OtpEnabled,
	}

	row, err := tx.
		Where(entities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

// GenerateOTP implements repositories.AuthRepository.
func (*AuthRepositoryImpl) GenerateOTP(tx *gorm.DB, userReq payloads.SecretUrlRequest, userID int) (bool, *exceptions.BaseErrorResponse) {
	user := entities.User{
		OtpSecret:  userReq.Secret,
		OtpAuthUrl: userReq.Url,
	}

	row, err := tx.
		Where(entities.User{ID: userID}).
		Updates(&user).
		Rows()
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer row.Close()

	return true, nil
}

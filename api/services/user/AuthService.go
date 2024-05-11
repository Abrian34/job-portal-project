package userservices

import (
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
)

type AuthService interface {
	Login(loginReq payloads.LoginRequest) (payloads.LoginResponse, *exceptions.BaseErrorResponse)
	// CheckPasswordResetTime(payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	Register(payloads.CreateRequest, int) (int, *exceptions.BaseErrorResponse)
	// GenerateOTP(int) (string, *exceptions.BaseErrorResponse)
	// UpdateUserOTP(entities.OTPInput, string) (*payloads.ResponseAuth, *exceptions.BaseErrorResponse)
	// UpdateCredential(payloads.LoginCredential, int) (bool, *exceptions.BaseErrorResponse)
	// UpdatePassword(*payloads.UserDetail, payloads.ChangePasswordInput) (bool, *exceptions.BaseErrorResponse)
	// UpdatePasswordTokenByEmail(payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	// UpdatePasswordByToken(payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse)
	// ResetPassword(string, payloads.ResetPasswordInput) (bool, *exceptions.BaseErrorResponse)
}

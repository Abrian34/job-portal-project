package userservices

import (
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/payloads"
)

type AuthService interface {
	Login(payloads.LoginRequest) (masterentities.User, *exceptions.BaseErrorResponse)
	CheckPasswordResetTime(payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	Register(payloads.CreateRequest, int) (int, *exceptions.BaseErrorResponse)
	GenerateOTP(int) (string, *exceptions.BaseErrorResponse)
	UpdateUserOTP(masterentities.OTPInput, string) (*payloads.ResponseAuth, *exceptions.BaseErrorResponse)
	UpdateCredential(payloads.LoginCredential, int) (bool, *exceptions.BaseErrorResponse)
	UpdatePassword(*payloads.UserDetail, payloads.ChangePasswordInput) (bool, *exceptions.BaseErrorResponse)
	UpdatePasswordTokenByEmail(payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse)
	UpdatePasswordByToken(payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse)
	ResetPassword(string, payloads.ResetPasswordInput) (bool, *exceptions.BaseErrorResponse)
}

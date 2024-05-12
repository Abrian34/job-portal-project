package userserviceimpl

import (
	"job-portal-project/api/exceptions"
	"job-portal-project/api/helper"
	"job-portal-project/api/payloads"
	userrepo "job-portal-project/api/repositories/user"
	"job-portal-project/api/securities"
	userservices "job-portal-project/api/services/user"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB             *gorm.DB
	AuthRepository userrepo.AuthRepository
	UserRepository userrepo.UserRepository
	Validate       *validator.Validate
}

func NewAuthService(
	db *gorm.DB,
	authRepository userrepo.AuthRepository,
	userRepository userrepo.UserRepository,
	validate *validator.Validate,
) userservices.AuthService {
	return &AuthServiceImpl{
		DB:             db,
		AuthRepository: authRepository,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements services.AuthService.
func (service *AuthServiceImpl) Login(loginReq payloads.LoginRequest) (payloads.LoginResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.GetByUsername(tx, loginReq.Username)
	if err != nil {
		return payloads.LoginResponse{}, err
	}

	pass := securities.VerifyPassword(user.UserPassword, loginReq.UserPassword)
	if pass != nil {
		return payloads.LoginResponse{}, err
	}

	role, err := service.AuthRepository.GetRoleWithPermissions(tx, user.RoleId)
	if err != nil {
		return payloads.LoginResponse{}, err
	}

	// token, err := securities.GenerateToken(user.UserName, user.UserId, role.RoleName, loginReq.Client)
	token, err := securities.GenerateToken(user)
	if err != nil {
		return payloads.LoginResponse{}, err
	}

	userDetails := payloads.UserDetail{
		UserId:          user.UserId,
		UserCode:        user.UserCode,
		UserDisplayName: user.UserDisplayName,
		RoleId:          role.RoleId,
		RoleName:        role.RoleName,
		UserName:        user.UserName,
		UserPassword:    user.UserPassword,
		ActiveStatus:    user.ActiveStatus,
	}

	// Construct login response with user details and role
	loginResponse := payloads.LoginResponse{
		User:        userDetails,
		Permissions: role.Permissions,
		Token:       token,
	}

	return loginResponse, nil
}

// Register implements services.AuthService.
func (service *AuthServiceImpl) Register(userReq payloads.CreateRequest, roleID int) (int, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	_, err := service.AuthRepository.CheckUserExists(tx, userReq.UserName)
	if err != nil {
		return 0, err
	}
	hash, err := securities.HashPassword(userReq.UserPassword)
	if err != nil {
		return 0, err
	}
	userReq.UserName = strings.ToLower(strings.ReplaceAll(userReq.UserName, " ", ""))
	userReq.UserPassword = strings.ReplaceAll(userReq.UserPassword, " ", "")
	userReq.UserPassword = hash

	get, err := service.UserRepository.Create(tx, userReq, roleID)

	if err != nil {
		return get, err
	}

	return get, nil
}

func (service *AuthServiceImpl) GetRoleWithPermissions(roleID int) (payloads.RoleResponse, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := service.AuthRepository.GetRoleWithPermissions(tx, roleID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (service *AuthServiceImpl) CheckUserExists(username string) (bool, *exceptions.BaseErrorResponse) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	get, err := service.AuthRepository.CheckUserExists(tx, username)

	if err != nil {
		return false, err
	}

	return get, nil
}

// CheckPasswordResetTime implements services.AuthService.
// func (service *AuthServiceImpl) CheckPasswordResetTime(emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	get, err := service.AuthRepository.CheckPasswordResetTime(tx, emailReq)

// 	if err != nil {
// 		return get, err
// 	}
// 	return get, nil
// }

// UpdatePassword implements services.AuthService.
// func (service *AuthServiceImpl) UpdatePassword(claims *payloads.UserDetail, changePasswordRequest payloads.ChangePasswordInput) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	getUser, err := service.UserRepository.GetByID(
// 		tx,
// 		claims.UserID,
// 	)
// 	if err != nil {
// 		return false, err
// 	}

// 	hashPwd := getUser.Password
// 	pwd := changePasswordRequest.OldPassword

// 	_, errors := securities.VerifyPassword(hashPwd, pwd)
// 	if errors != nil {
// 		return false,
// 			&exceptions.BaseErrorResponse{
// 				StatusCode: http.StatusBadRequest,
// 				Err:        errors,
// 			}
// 	}

// 	pass, err := securities.HashPassword(changePasswordRequest.NewPassword)
// 	if err != nil {
// 		return false, err
// 	}

// 	update, err := service.AuthRepository.UpdatePassword(tx, pass, claims.UserID)

// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

// UpdatePasswordByToken implements services.AuthService.
// func (service *AuthServiceImpl) UpdatePasswordByToken(passReq payloads.UpdatePasswordByTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	update, err := service.AuthRepository.UpdatePasswordByToken(tx, passReq)

// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

// ResetPassword implements services.AuthService.
// func (service *AuthServiceImpl) ResetPassword(resetToken string, passReq payloads.ResetPasswordInput) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)

// 	_, errors := securities.VerifyPassword(passReq.Password, passReq.PasswordConfirm)

// 	if errors != nil {
// 		return false, &exceptions.BaseErrorResponse{
// 			Err: errors,
// 		}
// 	}
// 	hashedPassword, err := securities.HashPassword(passReq.Password)

// 	if err != nil {
// 		return false, err
// 	}
// 	passwordResetToken := utils.Encode(resetToken)

// 	_, err = service.AuthRepository.CheckPasswordResetTime(
// 		tx,
// 		payloads.UpdateEmailTokenRequest{
// 			PasswordResetToken: utils.StringPtr(passwordResetToken),
// 			PasswordResetAt:    utils.TimePtr(time.Now()),
// 		})
// 	if err != nil {
// 		return false, err
// 	}

// 	_, err = service.AuthRepository.UpdatePasswordByToken(
// 		tx,
// 		payloads.UpdatePasswordByTokenRequest{
// 			Password:           hashedPassword,
// 			PasswordResetToken: utils.StringPtr(passwordResetToken),
// 		})

// 	if err != nil {
// 		return false, err
// 	}
// 	update, err := service.AuthRepository.ResetPassword(tx, payloads.ResetPasswordRequest{
// 		PasswordResetToken: utils.StringPtr(passwordResetToken),
// 	})

// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

// UpdatePasswordTokenByEmail implements services.AuthService.
// func (service *AuthServiceImpl) UpdatePasswordTokenByEmail(emailReq payloads.UpdateEmailTokenRequest) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	update, err := service.AuthRepository.UpdatePasswordTokenByEmail(tx, emailReq)

// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

// UpdateUserOTP implements services.AuthService.
// func (service *AuthServiceImpl) UpdateUserOTP(otpReq entities.OTPInput, remoteAddr string) (*payloads.ResponseAuth, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	txRedis := service.DBRedis
// 	defer helper.CommitOrRollback(tx)

// 	var loginCredential payloads.LoginCredential

// 	loginCredential.IpAddress = remoteAddr
// 	loginCredential.Client = otpReq.Client
// 	getUser, err := service.UserRepository.GetByID(
// 		tx,
// 		otpReq.UserID,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	token, err := securities.GenerateToken(getUser.Username, getUser.ID, getUser.RoleID, getUser.CompanyID, loginCredential.Client)

// 	if err != nil {
// 		return nil, err
// 	}

// 	loginCredential.Session = token
// 	_, err = service.AuthRepository.UpdateCredential(
// 		tx,
// 		loginCredential,
// 		int(getUser.ID),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = service.RedisRepository.UpdateCredential(txRedis, loginCredential, int(getUser.ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	valid := totp.Validate(otpReq.Token, getUser.OtpSecret)
// 	if !valid {
// 		return nil, &exceptions.BaseErrorResponse{
// 			StatusCode: http.StatusBadRequest,
// 			Err:        errors.New("token not valid"),
// 		}
// 	}

// 	updateOTP := payloads.OTPRequest{
// 		OtpVerified: true,
// 		OtpEnabled:  true,
// 	}
// 	response := payloads.ResponseAuth{
// 		Token:   token,
// 		Role:    getUser.RoleID,
// 		Company: getUser.CompanyID,
// 		UserID:  getUser.ID,
// 	}
// 	_, err = service.AuthRepository.UpdateUserOTP(tx, updateOTP, getUser.ID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

// GenerateOTP implements services.AuthService.
// func (service *AuthServiceImpl) GenerateOTP(userID int) (string, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)

// 	key, err := totp.Generate(totp.GenerateOpts{
// 		Issuer:      config.EnvConfigs.Issuer,
// 		AccountName: config.EnvConfigs.AccountName,
// 		SecretSize:  15,
// 	})

// 	if err != nil {
// 		return "", &exceptions.BaseErrorResponse{
// 			Err: err,
// 		}
// 	}

// 	updateSecretUrl := payloads.SecretUrlRequest{
// 		Secret: key.Secret(),
// 		Url:    key.URL(),
// 	}

// 	_, errors := service.AuthRepository.GenerateOTP(tx, updateSecretUrl, userID)

// 	if errors != nil {
// 		return "", errors
// 	}

// 	otpResponse := payloads.SecretUrlResponse{
// 		Secret: key.Secret(),
// 		Url:    key.URL(),
// 		UserID: userID,
// 	}

// 	img, _ := qrcode.NewQRCodeWriter().Encode(otpResponse.Url, gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
// 	fileName := fmt.Sprintf("%v-*.png", otpResponse.UserID)

// 	file, err := os.CreateTemp("", fileName)
// 	if err != nil {
// 		return "", &exceptions.BaseErrorResponse{
// 			Err: err,
// 		}
// 	}

// 	defer os.Remove(file.Name())

// 	_ = png.Encode(file, img)

// 	imgFile, err := os.Open(file.Name())
// 	if err != nil {
// 		return "", &exceptions.BaseErrorResponse{
// 			Err: err,
// 		}
// 	}
// 	defer imgFile.Close()

// 	return file.Name(), nil
// }

// UpdateCredential implements services.AuthService.
// func (service *AuthServiceImpl) UpdateCredential(loginReq payloads.LoginCredential, userID int) (bool, *exceptions.BaseErrorResponse) {
// 	tx := service.DB.Begin()
// 	defer helper.CommitOrRollback(tx)
// 	update, err := service.AuthRepository.UpdateCredential(tx, loginReq, userID)
// 	if err != nil {
// 		return update, err
// 	}

// 	return update, nil
// }

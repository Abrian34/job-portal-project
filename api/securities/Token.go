package securities

import (
	"job-portal-project/api/config"
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const AuthTokenValidTime = time.Hour * 3

func GenerateToken(userDetail payloads.UserDetail) (string, *exceptions.BaseErrorResponse) {
	secret := config.EnvConfigs.JWTKey
	claims := jwt.MapClaims{}

	claims["user_id"] = userDetail.UserId
	claims["username"] = userDetail.UserName
	claims["user_code"] = userDetail.UserCode
	claims["user_display_name"] = userDetail.UserDisplayName
	claims["role_id"] = userDetail.RoleId
	claims["role_name"] = userDetail.RoleName
	claims["active_status"] = userDetail.ActiveStatus

	expirationTime := time.Now().Add(AuthTokenValidTime)
	claims["exp"] = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return tokenString, nil
}

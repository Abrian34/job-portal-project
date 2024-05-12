package securities

import (
	"errors"
	"fmt"
	"job-portal-project/api/config"
	"job-portal-project/api/utils/constant"

	// redisservices "job-portal-project/api/services/redis"

	"strings"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func GetAuthentication(request *http.Request) error {
	token, err := VerifyToken(request)
	if err != nil {
		return errors.New(constant.SessionError)
	}

	if !token.Valid {
		return errors.New(constant.SessionError)
	}

	return nil
}

func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(request)
	if tokenString == "" {
		return nil, errors.New("session invalid, please re-login")
	}
	secretKey := config.EnvConfigs.JWTKey
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errors.New("session invalid, please re-login")
	}

	return token, nil
}

func ExtractToken(request *http.Request) string {
	// Get the query string parameters.
	keys := request.URL.Query()
	token := keys.Get("token")

	if token != "" {
		return token
	}

	// Get the Authorization header.
	authHeader := request.Header.Get("Authorization")

	// If the Authorization header is not empty, split it into two parts.
	if authHeader != "" {
		bearerToken := strings.Split(authHeader, " ")

		// If the Authorization header is in the format "Bearer token", return the token.
		if len(bearerToken) == 2 {
			return bearerToken[1]
		}
	}

	// If no token is found, return an empty string.
	return ""
}

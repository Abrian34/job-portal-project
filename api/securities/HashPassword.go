package securities

import (
	"job-portal-project/api/exceptions"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, *exceptions.BaseErrorResponse) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	hash := string(hashPassword)
	return hash, nil
}

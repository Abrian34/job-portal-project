package securities

import (
	"job-portal-project/api/exceptions"
	"job-portal-project/api/payloads"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractAuthToken(request *http.Request) (*payloads.UserDetail, *exceptions.BaseErrorResponse) {
	token, err := VerifyToken(request)
	if err != nil {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	userID := int(claims["user_id"].(float64))
	username := claims["username"].(string)
	userCode := claims["user_code"].(string)
	userDisplayName := claims["user_display_name"].(string)
	roleId := int(claims["role_id"].(float64))
	roleName := claims["role_name"].(string)
	activeStatus := claims["active_status"].(bool)

	userDetail := payloads.UserDetail{
		UserId:          userID,
		UserName:        username,
		UserCode:        userCode,
		UserDisplayName: userDisplayName,
		RoleId:          roleId,
		RoleName:        roleName,
		ActiveStatus:    activeStatus,
	}

	return &userDetail, nil
}

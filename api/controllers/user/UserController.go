package usercontrollers

import (
	"job-portal-project/api/exceptions"
	"job-portal-project/api/helper"
	jsonchecker "job-portal-project/api/helper/json/json-checker"
	"job-portal-project/api/payloads"
	"job-portal-project/api/securities"
	userservices "job-portal-project/api/services/user"
	"job-portal-project/api/utils/constant"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController interface {
	GetCurrentUser(writer http.ResponseWriter, request *http.Request)
	GetUser(writer http.ResponseWriter, request *http.Request)
	GetUserIDByUsername(writer http.ResponseWriter, request *http.Request)
	GetUsernameByUserID(writer http.ResponseWriter, request *http.Request)
	FindUser(writer http.ResponseWriter, request *http.Request)
	Register(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
	GetRoleById(writer http.ResponseWriter, request *http.Request)
}

type UserControllerImpl struct {
	UserService userservices.UserService
	AuthService userservices.AuthService
}

func (controller *UserControllerImpl) GetCurrentUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.GetCurrentUser(claims.UserId)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, constant.GetDataSuccess, http.StatusOK)
}

func NewUserController(userService userservices.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	var registrationRequest payloads.CreateRequest
	err := jsonchecker.ReadFromRequestBody(request, &registrationRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}

	roleID, _ := strconv.Atoi(chi.URLParam(request, "role_id"))

	userID, err := controller.AuthService.Register(registrationRequest, roleID)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	response := payloads.RegisterResponse{
		UserID: userID,
	}
	payloads.NewHandleSuccess(writer, response, "User registered successfully", http.StatusCreated)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	var loginRequest payloads.LoginRequest
	err := jsonchecker.ReadFromRequestBody(request, &loginRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}

	loginResponse, err := controller.AuthService.Login(loginRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, loginResponse, "User Login successfully", http.StatusCreated)
}

func (controller *UserControllerImpl) GetUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.GetUser(claims.UserName)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, constant.GetDataSuccess, http.StatusOK)
}

func (controller *UserControllerImpl) GetRoleById(writer http.ResponseWriter, request *http.Request) {
	RoleId, _ := strconv.Atoi(chi.URLParam(request, "role_id"))
	RoleResponse, errors := controller.UserService.GetRoleById(RoleId)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.NewHandleSuccess(writer, RoleResponse, constant.GetDataSuccess, http.StatusOK)
}

func (controller *UserControllerImpl) GetUsernameByUserID(writer http.ResponseWriter, request *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(request, "user_id"))
	if err != nil {
		exceptions.NewAppException(writer, request, &exceptions.BaseErrorResponse{
			Err: err,
		})
		return
	}
	userResponse, errors := controller.UserService.GetUsernameByUserID(userID)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.HandleSuccess(writer, userResponse, constant.GetDataSuccess, http.StatusOK)
}

func (controller *UserControllerImpl) GetUserIDByUsername(writer http.ResponseWriter, request *http.Request) {
	username := chi.URLParam(request, "username")
	userResponse, err := controller.UserService.GetUserIDByUsername(username)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, constant.GetDataSuccess, http.StatusOK)
}

func (controller *UserControllerImpl) FindUser(writer http.ResponseWriter, request *http.Request) {
	claims, _ := securities.ExtractAuthToken(request)

	userResponse, err := controller.UserService.FindUser(claims.UserName)

	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.HandleSuccess(writer, userResponse, constant.GetDataSuccess, http.StatusOK)
}

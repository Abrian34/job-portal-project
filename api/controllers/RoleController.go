package controller

import (
	exceptions "job-portal-project/api/exceptions"
	"job-portal-project/api/utils/constant"

	helper "job-portal-project/api/helper"
	jsonchecker "job-portal-project/api/helper/json/json-checker"
	"job-portal-project/api/payloads"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	service "job-portal-project/api/services"
	"job-portal-project/api/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController interface {
	GetRoleList(writer http.ResponseWriter, request *http.Request)
	GetRoleById(writer http.ResponseWriter, request *http.Request)
	SaveRole(writer http.ResponseWriter, request *http.Request)
	ChangeStatusRole(writer http.ResponseWriter, request *http.Request)
}

type RoleControllerImpl struct {
	RoleService service.RoleService
}

func NewRoleController(RoleService service.RoleService) RoleController {
	return &RoleControllerImpl{
		RoleService: RoleService,
	}
}

func (r *RoleControllerImpl) GetRoleList(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	queryParams := map[string]string{
		"role_name":        queryValues.Get("role_name"),
		"role_description": queryValues.Get("role_description"),
	}
	pagination := pagination.Pagination{
		Limit:  utils.NewGetQueryInt(queryValues, "limit"),
		Page:   utils.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}

	filterCondition := utils.BuildFilterCondition(queryParams)
	result, err := r.RoleService.GetRoleList(filterCondition, pagination)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccessPagination(writer, result.Rows, "Get Data Successfully!", 200, result.Limit, result.Page, result.TotalRows, result.TotalPages)
}

func (r *RoleControllerImpl) GetRoleById(writer http.ResponseWriter, request *http.Request) {
	RoleId, _ := strconv.Atoi(chi.URLParam(request, "role_id"))
	RoleResponse, errors := r.RoleService.GetRoleById(RoleId)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.NewHandleSuccess(writer, RoleResponse, constant.GetDataSuccess, http.StatusOK)
}

func (r *RoleControllerImpl) SaveRole(writer http.ResponseWriter, request *http.Request) {
	RoleRequest := entitypayloads.RoleRequest{}
	var message string

	err := jsonchecker.ReadFromRequestBody(request, &RoleRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	create, err := r.RoleService.SaveRole(RoleRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, message, http.StatusCreated)
}

func (r *RoleControllerImpl) ChangeStatusRole(writer http.ResponseWriter, request *http.Request) {

	RoleId, _ := strconv.Atoi(chi.URLParam(request, "role_id"))

	response, err := r.RoleService.ChangeStatusRole(int(RoleId))
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, response, "Update Data Successfully!", http.StatusOK)
}

func (r *RoleControllerImpl) UpdateRole(writer http.ResponseWriter, request *http.Request) {
	RoleRequest := entitypayloads.RoleUpdate{}
	var message string

	err := jsonchecker.ReadFromRequestBody(request, &RoleRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	create, err := r.RoleService.UpdateRole(RoleRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, message, http.StatusCreated)
}

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

type JobApplicationController interface {
	GetJobApplicationListByJobId(writer http.ResponseWriter, request *http.Request)
	GetJobApplicationById(writer http.ResponseWriter, request *http.Request)
	SaveJobApplication(writer http.ResponseWriter, request *http.Request)
	ChangeStatusJobApplication(writer http.ResponseWriter, request *http.Request)
	UpdateJobApplication(writer http.ResponseWriter, request *http.Request)
}

type JobApplicationControllerImpl struct {
	JobApplicationService service.JobApplicationService
}

func NewJobApplicationController(JobApplicationService service.JobApplicationService) JobApplicationController {
	return &JobApplicationControllerImpl{
		JobApplicationService: JobApplicationService,
	}
}

func (r *JobApplicationControllerImpl) GetJobApplicationListByJobId(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()
	jobId, _ := strconv.Atoi(chi.URLParam(request, "job_id"))
	queryParams := map[string]string{
		"job_id":             queryValues.Get("job_id"),
		"application_status": queryValues.Get("application_status"),
	}
	pagination := pagination.Pagination{
		Limit:  utils.NewGetQueryInt(queryValues, "limit"),
		Page:   utils.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}

	filterCondition := utils.BuildFilterCondition(queryParams)
	result, err := r.JobApplicationService.GetJobApplicationListByJobId(filterCondition, jobId, pagination)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccessPagination(writer, result.Rows, "Get Data Successfully!", 200, result.Limit, result.Page, result.TotalRows, result.TotalPages)
}

func (r *JobApplicationControllerImpl) GetJobApplicationById(writer http.ResponseWriter, request *http.Request) {
	JobApplicationId, _ := strconv.Atoi(chi.URLParam(request, "job_application_id"))
	JobApplicationResponse, errors := r.JobApplicationService.GetJobApplicationById(JobApplicationId)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.NewHandleSuccess(writer, JobApplicationResponse, constant.GetDataSuccess, http.StatusOK)
}

func (r *JobApplicationControllerImpl) SaveJobApplication(writer http.ResponseWriter, request *http.Request) {
	JobApplicationRequest := entitypayloads.JobApplicationRequest{}
	var message string

	err := jsonchecker.ReadFromRequestBody(request, &JobApplicationRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	create, err := r.JobApplicationService.SaveJobApplication(JobApplicationRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, message, http.StatusCreated)
}

func (r *JobApplicationControllerImpl) ChangeStatusJobApplication(writer http.ResponseWriter, request *http.Request) {

	JobApplicationId, _ := strconv.Atoi(chi.URLParam(request, "job_application_id"))

	response, err := r.JobApplicationService.ChangeStatusJobApplication(int(JobApplicationId))
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, response, "Update Data Successfully!", http.StatusOK)
}

func (r *JobApplicationControllerImpl) UpdateJobApplication(writer http.ResponseWriter, request *http.Request) {
	JobApplicationRequest := entitypayloads.JobApplicationUpdate{}
	var message string

	err := jsonchecker.ReadFromRequestBody(request, &JobApplicationRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	create, err := r.JobApplicationService.UpdateJobApplication(JobApplicationRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, create, message, http.StatusCreated)
}

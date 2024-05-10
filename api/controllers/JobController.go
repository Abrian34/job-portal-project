package controller

import (
	exceptions "job-portal-project/api/exceptions"
	"job-portal-project/api/utils/constant"

	// "job-portal-project/api/helper"
	helper "job-portal-project/api/helper"
	jsonchecker "job-portal-project/api/helper/json/json-checker"
	"job-portal-project/api/payloads"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	service "job-portal-project/api/services"
	"job-portal-project/api/utils"
	"job-portal-project/api/validation"
	"net/http"
	"strconv"

	// "job-portal-project/api/utils/validation"

	"github.com/go-chi/chi/v5"
	// "github.com/julienschmidt/httprouter"
)

type JobController interface {
	GetJobList(writer http.ResponseWriter, request *http.Request)
	GetJobById(writer http.ResponseWriter, request *http.Request)
	SaveJob(writer http.ResponseWriter, request *http.Request)
	ChangeStatusJob(writer http.ResponseWriter, request *http.Request)
}

type JobControllerImpl struct {
	JobService service.JobService
}

func NewJobController(JobService service.JobService) JobController {
	return &JobControllerImpl{
		JobService: JobService,
	}
}

func (r *JobControllerImpl) GetJobList(writer http.ResponseWriter, request *http.Request) {
	queryValues := request.URL.Query()

	queryParams := map[string]string{
		"job_code":  queryValues.Get("job_code"),
		"job_title": queryValues.Get("job_title"),
		"job_level": queryValues.Get("job_level"),
	}
	pagination := pagination.Pagination{
		Limit:  utils.NewGetQueryInt(queryValues, "limit"),
		Page:   utils.NewGetQueryInt(queryValues, "page"),
		SortOf: queryValues.Get("sort_of"),
		SortBy: queryValues.Get("sort_by"),
	}

	filterCondition := utils.BuildFilterCondition(queryParams)
	result, err := r.JobService.GetJobList(filterCondition, pagination)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccessPagination(writer, result.Rows, "Get Data Successfully!", 200, result.Limit, result.Page, result.TotalRows, result.TotalPages)
}

func (r *JobControllerImpl) GetJobById(writer http.ResponseWriter, request *http.Request) {
	JobId, _ := strconv.Atoi(chi.URLParam(request, "job_id"))
	JobResponse, errors := r.JobService.GetJobById(JobId)

	if errors != nil {
		helper.ReturnError(writer, request, errors)
		return
	}
	payloads.NewHandleSuccess(writer, JobResponse, constant.GetDataSuccess, http.StatusOK)
}

func (r *JobControllerImpl) SaveJob(writer http.ResponseWriter, request *http.Request) {
	JobRequest := entitypayloads.JobPayload{}
	var message string

	err := jsonchecker.ReadFromRequestBody(request, &JobRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}
	err = validation.ValidationForm(writer, request, JobRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	create, err := r.JobService.SaveJob(JobRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	if JobRequest.JobId == 0 {
		message = "Create Data Successfully!"
	} else {
		message = "Update Data Successfully!"
	}

	payloads.NewHandleSuccess(writer, create, message, http.StatusCreated)
}

func (r *JobControllerImpl) ChangeStatusJob(writer http.ResponseWriter, request *http.Request) {

	JobId, _ := strconv.Atoi(chi.URLParam(request, "job_id"))

	response, err := r.JobService.ChangeStatusJob(int(JobId))
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}

	payloads.NewHandleSuccess(writer, response, "Update Data Successfully!", http.StatusOK)
}

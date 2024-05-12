package service

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"
)

type JobApplicationService interface {
	GetJobApplicationListByJobId(filterCondition []utils.FilterCondition, jobId int, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	GetJobApplicationById(id int) (entitypayloads.JobApplicationPayload, *exceptions.BaseErrorResponse)
	SaveJobApplication(req entitypayloads.JobApplicationRequest) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusJobApplication(id int) (bool, *exceptions.BaseErrorResponse)
	UpdateJobApplication(req entitypayloads.JobApplicationUpdate) (bool, *exceptions.BaseErrorResponse)
}

package service

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"
)

type JobService interface {
	GetJobList([]utils.FilterCondition, pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	GetJobById(int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse)
	SaveJob(entitypayloads.JobPayload) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusJob(int) (bool, *exceptions.BaseErrorResponse)
}

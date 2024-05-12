package repository

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"

	"gorm.io/gorm"
)

type JobApplicationRepository interface {
	GetJobApplicationById(tx *gorm.DB, Id int) (entitypayloads.JobApplicationPayload, *exceptions.BaseErrorResponse)
	SaveJobApplication(tx *gorm.DB, req entitypayloads.JobApplicationRequest) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusJobApplication(tx *gorm.DB, Id int) (bool, *exceptions.BaseErrorResponse)
	GetJobApplicationListByJobId(tx *gorm.DB, jobId int, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	UpdateJobApplication(tx *gorm.DB, req entitypayloads.JobApplicationUpdate) (bool, *exceptions.BaseErrorResponse)
}

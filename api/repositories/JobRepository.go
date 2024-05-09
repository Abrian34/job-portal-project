package repository

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"

	"gorm.io/gorm"
)

type JobRepository interface {
	GetJobById(tx *gorm.DB, Id int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse)
	SaveJob(tx *gorm.DB, req entitypayloads.JobPayload) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusJob(tx *gorm.DB, Id int) (bool, *exceptions.BaseErrorResponse)
	GetJobList(tx *gorm.DB, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
}

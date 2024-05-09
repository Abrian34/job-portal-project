package repository

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"

	"gorm.io/gorm"
)

type JobRepository interface {
	GetJobById(*gorm.DB, int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse)
	SaveJob(*gorm.DB, entitypayloads.JobPayload) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusJob(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	GetJobList(*gorm.DB, []utils.FilterCondition, pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
}

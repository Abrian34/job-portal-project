package repository

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetRoleById(*gorm.DB, int) (entitypayloads.RolePayload, *exceptions.BaseErrorResponse)
	SaveRole(*gorm.DB, entitypayloads.RoleRequest) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusRole(*gorm.DB, int) (bool, *exceptions.BaseErrorResponse)
	GetRoleList(*gorm.DB, []utils.FilterCondition, pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	UpdateRole(*gorm.DB, entitypayloads.RoleUpdate) (bool, *exceptions.BaseErrorResponse)
}

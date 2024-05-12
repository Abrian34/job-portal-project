package service

import (
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	"job-portal-project/api/utils"
)

type RoleService interface {
	GetRoleList([]utils.FilterCondition, pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse)
	GetRoleById(int) (entitypayloads.RolePayload, *exceptions.BaseErrorResponse)
	SaveRole(entitypayloads.RoleRequest) (bool, *exceptions.BaseErrorResponse)
	ChangeStatusRole(int) (bool, *exceptions.BaseErrorResponse)
	UpdateRole(entitypayloads.RoleUpdate) (bool, *exceptions.BaseErrorResponse)
}

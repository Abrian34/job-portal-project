package repository

import (
	masteroperationentitypayloads "job-portal-project/api/entitypayloads/master/operation"
	"job-portal-project/api/entitypayloads/pagination"
	exceptionsss_test "job-portal-project/api/exceptions"
	"job-portal-project/api/utils"

	"gorm.io/gorm"
)

type OperationKeyRepository interface {
	GetOperationKeyById(*gorm.DB, int) (masteroperationentitypayloads.OperationkeyListResponse, *exceptionsss_test.BaseErrorResponse)
	GetOperationKeyName(*gorm.DB, masteroperationentitypayloads.OperationKeyRequest) (masteroperationentitypayloads.OperationKeyNameResponse, *exceptionsss_test.BaseErrorResponse)
	SaveOperationKey(*gorm.DB, masteroperationentitypayloads.OperationKeyResponse) (bool, *exceptionsss_test.BaseErrorResponse)
	GetAllOperationKeyList(*gorm.DB, []utils.FilterCondition, pagination.Pagination) (pagination.Pagination, *exceptionsss_test.BaseErrorResponse)
	ChangeStatusOperationKey(*gorm.DB, int) (bool, *exceptionsss_test.BaseErrorResponse)
}

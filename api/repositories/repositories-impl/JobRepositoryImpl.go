package masterrepositoryimpl

import (
	"after-sales/api/config"
	masterentities "after-sales/api/entities/master"
	exceptionsss_test "after-sales/api/expectionsss"
	masterpayloads "after-sales/api/payloads/master"
	"after-sales/api/payloads/pagination"
	masterrepository "after-sales/api/repositories/master"
	"after-sales/api/utils"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type JobRepositoryImpl struct {
}

func StartJobRepositoryImpl() masterrepository.JobRepository {
	return &JobRepositoryImpl{}
}

func (r *JobRepositoryImpl) GetJobById(tx *gorm.DB, JobId int) (masterpayloads.JobPayload, *exceptionsss_test.BaseErrorResponse) {
	entities := masterentities.Job{}
	response := masterpayloads.JobPayload{}

	err := tx.Model(&entities).
		Where(masterentities.Job{
			JobId: JobId,
		}).
		First(&entities).
		Error

	if err != nil {
		return response, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	// Copying values from entities to response
	response.JobId = entities.JobId
	response.JobCode = entities.JobCode
	response.IsActive = entities.IsActive
	response.BrandId = entities.BrandId
	response.CustomerId = entities.CustomerId
	response.ProfitCenterId = entities.ProfitCenterId
	response.JobDateFrom = entities.JobDateFrom
	response.JobDateTo = entities.JobDateTo
	response.DealerId = entities.DealerId
	response.TopId = entities.TopId
	response.JobRemark = entities.JobRemark

	return response, nil
}

func (r *JobRepositoryImpl) SaveJob(tx *gorm.DB, req masterpayloads.JobResponse) (bool, *exceptionsss_test.BaseErrorResponse) {
	entities := masterentities.Job{
		JobCode:     req.JobCode,
		BrandId:           req.BrandId,
		DealerId:          req.DealerId,
		TopId:             req.TopId,
		JobDateFrom: req.JobDateFrom,
		JobDateTo:   req.JobDateTo,
		JobRemark:   req.JobRemark,
		ProfitCenterId:    req.ProfitCenterId,
		IsActive:          req.IsActive,
		JobId:       req.JobId,
		CustomerId:        req.CustomerId,
	}

	err := tx.Save(&entities).Error

	if err != nil {
		return false, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

func (r *JobRepositoryImpl) ChangeStatusJob(tx *gorm.DB, Id int) (bool, *exceptionsss_test.BaseErrorResponse) {
	var entities masterentities.Job

	result := tx.Model(&entities).
		Where("Job_id = ?", Id).
		First(&entities)

	if result.Error != nil {
		return false, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        result.Error,
		}
	}

	if entities.IsActive {
		entities.IsActive = false
	} else {
		entities.IsActive = true
	}

	result = tx.Save(&entities)

	if result.Error != nil {
		return false, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        result.Error,
		}
	}

	return true, nil
}

func (r *JobRepositoryImpl) GetJobList(tx *gorm.DB, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptionsss_test.BaseErrorResponse) {
	entities := masterentities.Job{}
	var responses []masteroperationpayloads.JobPayload

	// define table struct
	tableStruct := masteroperationpayloads.JobPayload{}

	//join table
	joinTable := utils.CreateJoinSelectStatement(tx, tableStruct)

	//apply filter
	whereQuery := utils.ApplyFilter(joinTable, filterCondition)
	//apply pagination and execute
	rows, err := joinTable.Scopes(pagination.Paginate(&entities, &pages, whereQuery)).Scan(&responses).Rows()

	if err != nil {
		return pages, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if len(responses) == 0 {
		return pages, &exceptionsss_test.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	defer rows.Close()

	pages.Rows = responses

	return pages, nil
}
package masterrepositoryimpl

import (
	entities "job-portal-project/api/entities/job"
	entitypayloads "job-portal-project/api/entitypayloads"
	exceptions "job-portal-project/api/exceptions"

	// "job-portal-project/api/entitypayloads/pagination"
	// masterrepository "job-portal-project/api/repositories"
	"job-portal-project/api/utils"
	"net/http"

	"gorm.io/gorm"
)

type JobRepositoryImpl struct {
}

func StartJobRepositoryImpl() masterrepository.JobRepository {
	return &JobRepositoryImpl{}
}

func (r *JobRepositoryImpl) GetJobById(tx *gorm.DB, JobId int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse) {
	entities := entities.Job{}
	response := entitypayloads.JobPayload{}

	err := tx.Model(&entities).
		Where(entities.Job{
			JobId: JobId,
		}).
		First(&entities).
		Error

	if err != nil {
		return response, &exceptions.BaseErrorResponse{
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

func (r *JobRepositoryImpl) SaveJob(tx *gorm.DB, req entitypayloads.JobResponse) (bool, *exceptions.BaseErrorResponse) {
	entities := entities.Job{
		JobCode:        req.JobCode,
		BrandId:        req.BrandId,
		DealerId:       req.DealerId,
		TopId:          req.TopId,
		JobDateFrom:    req.JobDateFrom,
		JobDateTo:      req.JobDateTo,
		JobRemark:      req.JobRemark,
		ProfitCenterId: req.ProfitCenterId,
		IsActive:       req.IsActive,
		JobId:          req.JobId,
		CustomerId:     req.CustomerId,
	}

	err := tx.Save(&entities).Error

	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

func (r *JobRepositoryImpl) ChangeStatusJob(tx *gorm.DB, Id int) (bool, *exceptions.BaseErrorResponse) {
	var entities entities.Job

	result := tx.Model(&entities).
		Where("Job_id = ?", Id).
		First(&entities)

	if result.Error != nil {
		return false, &exceptions.BaseErrorResponse{
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
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        result.Error,
		}
	}

	return true, nil
}

func (r *JobRepositoryImpl) GetJobList(tx *gorm.DB, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	entities := entities.Job{}
	var responses []entitypayloads.JobPayload

	// define table struct
	tableStruct := entitypayloads.JobPayload{}

	//join table
	joinTable := utils.CreateJoinSelectStatement(tx, tableStruct)

	//apply filter
	whereQuery := utils.ApplyFilter(joinTable, filterCondition)
	//apply pagination and execute
	rows, err := joinTable.Scopes(pagination.Paginate(&entities, &pages, whereQuery)).Scan(&responses).Rows()

	if err != nil {
		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	if len(responses) == 0 {
		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	defer rows.Close()

	pages.Rows = responses

	return pages, nil
}

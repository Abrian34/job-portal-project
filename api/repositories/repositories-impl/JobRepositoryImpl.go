package repositoryimpl

import (
	entities "job-portal-project/api/entities"
	exceptions "job-portal-project/api/exceptions"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"

	// "job-portal-project/api/entitypayloads/pagination"
	repository "job-portal-project/api/repositories"
	"job-portal-project/api/utils"
	"net/http"

	"gorm.io/gorm"
)

type JobRepositoryImpl struct {
}

func StartJobRepositoryImpl() repository.JobRepository {
	return &JobRepositoryImpl{}
}

func (*JobRepositoryImpl) GetJobById(tx *gorm.DB, Id int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse) {
	var JobMapping entities.Job
	var JobResponse entitypayloads.JobPayload

	rows, err := tx.
		Model(&JobMapping).
		Where(entities.Job{JobId: Id}).
		First(&JobResponse).
		Rows()

	if err != nil {

		return JobResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer rows.Close()

	return JobResponse, nil
}

func (r *JobRepositoryImpl) SaveJob(tx *gorm.DB, req entitypayloads.JobPayload) (bool, *exceptions.BaseErrorResponse) {
	entities := entities.Job{
		JobCode:        req.JobCode,
		JobId:          req.JobId,
		EmployerId:     req.EmployerId,
		JobPostDate:    req.JobPostDate,
		CompanyId:      req.CompanyId,
		JobTitle:       req.JobTitle,
		JobDescription: req.JobDescription,
		JobLevel:       req.JobLevel,
		JobVacancy:     req.JobVacancy,
		ActiveStatus:   req.ActiveStatus,
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

	if entities.ActiveStatus {
		entities.ActiveStatus = false
	} else {
		entities.ActiveStatus = true
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
	JobMapping := []entities.Job{}
	JobResponse := []entitypayloads.JobPayload{}
	query := tx.
		Model(entities.Job{}).
		Scan(&JobResponse)

	ApplyFilter := utils.ApplyFilter(query, filterCondition)

	err := ApplyFilter.
		Scopes(pagination.Paginate(&JobMapping, &pages, ApplyFilter)).
		// Order("").
		Scan(&JobResponse).
		Error

	if len(JobResponse) == 0 {
		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	if err != nil {

		return pages, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	// defer row.Close()
	pages.Rows = JobResponse

	return pages, nil
}

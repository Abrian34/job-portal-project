package repositoryimpl

import (
	"errors"
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

type JobApplicationRepositoryImpl struct {
}

func StartJobApplicationRepositoryImpl() repository.JobApplicationRepository {
	return &JobApplicationRepositoryImpl{}
}

func (*JobApplicationRepositoryImpl) GetJobApplicationById(tx *gorm.DB, Id int) (entitypayloads.JobApplicationPayload, *exceptions.BaseErrorResponse) {
	var JobApplicationMapping entities.JobApplication
	var JobApplicationResponse entitypayloads.JobApplicationPayload

	rows, err := tx.
		Model(&JobApplicationMapping).
		Where(entities.JobApplication{JobApplicationId: Id}).
		First(&JobApplicationResponse).
		Rows()

	if err != nil {

		return JobApplicationResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer rows.Close()

	return JobApplicationResponse, nil
}

func (r *JobApplicationRepositoryImpl) SaveJobApplication(tx *gorm.DB, req entitypayloads.JobApplicationRequest) (bool, *exceptions.BaseErrorResponse) {
	entities := entities.JobApplication{
		JobId:             req.JobId,
		UserId:            req.UserId,
		CoverLetter:       req.CoverLetter,
		ApplicationStatus: "Applying",
		ActiveStatus:      true,
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

func (r *JobApplicationRepositoryImpl) ChangeStatusJobApplication(tx *gorm.DB, Id int) (bool, *exceptions.BaseErrorResponse) {
	var entities entities.JobApplication

	result := tx.Model(&entities).
		Where("JobApplication_id = ?", Id).
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

func (r *JobApplicationRepositoryImpl) GetJobApplicationListByJobId(tx *gorm.DB, jobId int, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	JobApplicationMapping := []entities.JobApplication{}
	JobApplicationResponse := []entitypayloads.JobApplicationPayload{}
	query := tx.
		Model(entities.JobApplication{}).
		Where(entities.JobApplication{JobId: jobId}).
		Scan(&JobApplicationResponse)

	ApplyFilter := utils.ApplyFilter(query, filterCondition)

	err := ApplyFilter.
		Scopes(pagination.Paginate(&JobApplicationMapping, &pages, ApplyFilter)).
		Scan(&JobApplicationResponse).
		Error

	if len(JobApplicationResponse) == 0 {
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
	pages.Rows = JobApplicationResponse

	return pages, nil
}

func (r *JobApplicationRepositoryImpl) UpdateJobApplication(tx *gorm.DB, req entitypayloads.JobApplicationUpdate) (bool, *exceptions.BaseErrorResponse) {
	var model entities.JobApplication
	if err := tx.Where("JobApplication_id = ?", req.JobApplicationId).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusNotFound,
				Err:        errors.New("JobApplication not found"),
			}
		}
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	model.ApplicationStatus = req.ApplicationStatus

	if err := tx.Save(&model).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

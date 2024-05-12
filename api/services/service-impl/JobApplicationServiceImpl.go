package serviceimpl

import (
	exceptions "job-portal-project/api/exceptions"
	"job-portal-project/api/helper"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	repository "job-portal-project/api/repositories"
	service "job-portal-project/api/services"
	"job-portal-project/api/utils"

	// "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type JobApplicationServiceImpl struct {
	JobApplicationRepository repository.JobApplicationRepository
	DB                       *gorm.DB
	// RedisClient   *redis.Client // Redis client
}

func StartJobApplicationService(JobApplicationRepository repository.JobApplicationRepository, db *gorm.DB) service.JobApplicationService {
	return &JobApplicationServiceImpl{
		JobApplicationRepository: JobApplicationRepository,
		DB:                       db,
	}
}

func (s *JobApplicationServiceImpl) GetJobApplicationListByJobId(filterCondition []utils.FilterCondition, jobId int, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := s.JobApplicationRepository.GetJobApplicationListByJobId(tx, jobId, filterCondition, pages)

	if err != nil {
		return get, err
	}

	return get, nil
}
func (s *JobApplicationServiceImpl) GetJobApplicationById(id int) (entitypayloads.JobApplicationPayload, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := s.JobApplicationRepository.GetJobApplicationById(tx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *JobApplicationServiceImpl) SaveJobApplication(req entitypayloads.JobApplicationRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	results, err := s.JobApplicationRepository.SaveJobApplication(tx, req)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (s *JobApplicationServiceImpl) ChangeStatusJobApplication(id int) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := s.JobApplicationRepository.GetJobApplicationById(tx, id)

	if err != nil {
		// panic(exceptions.NewNotFoundError(err.Error()))
		return false, err
	}

	results, err := s.JobApplicationRepository.ChangeStatusJobApplication(tx, id)
	if err != nil {
		return results, err
	}
	return true, nil
}

func (s *JobApplicationServiceImpl) UpdateJobApplication(req entitypayloads.JobApplicationUpdate) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	results, err := s.JobApplicationRepository.UpdateJobApplication(tx, req)
	if err != nil {
		return results, err
	}

	return results, nil
}

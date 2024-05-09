package serviceimpl

import (
	exceptions "job-portal-project/api/exceptions"
	"job-portal-project/api/helper"
	entitypayloads "job-portal-project/api/payloads/entity-payloads"
	"job-portal-project/api/payloads/pagination"
	repository "job-portal-project/api/repositories"
	service "job-portal-project/api/services"
	"job-portal-project/api/utils"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type JobServiceImpl struct {
	JobRepository repository.JobRepository
	DB            *gorm.DB
	RedisClient   *redis.Client // Redis client
}

func StartJobService(JobRepository repository.JobRepository, db *gorm.DB, redisClient *redis.Client) service.JobService {
	return &JobServiceImpl{
		JobRepository: JobRepository,
		DB:            db,
		RedisClient:   redisClient,
	}
}

func (s *JobServiceImpl) GetJobList(filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := s.JobRepository.GetJobList(tx, filterCondition, pages)

	if err != nil {
		return get, err
	}

	return get, nil
}
func (s *JobServiceImpl) GetJobById(id int) (entitypayloads.JobPayload, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := s.JobRepository.GetJobById(tx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *JobServiceImpl) SaveJob(req entitypayloads.JobPayload) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	results, err := s.JobRepository.SaveJob(tx, req)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (s *JobServiceImpl) ChangeStatusJob(id int) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := s.JobRepository.GetJobById(tx, id)

	if err != nil {
		// panic(exceptions.NewNotFoundError(err.Error()))
		return false, err
	}

	results, err := s.JobRepository.ChangeStatusJob(tx, id)
	if err != nil {
		return results, err
	}
	return true, nil
}

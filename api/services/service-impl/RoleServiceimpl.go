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

type RoleServiceImpl struct {
	RoleRepository repository.RoleRepository
	DB             *gorm.DB
	// RedisClient   *redis.Client // Redis client
}

func StartRoleService(RoleRepository repository.RoleRepository, db *gorm.DB) service.RoleService {
	return &RoleServiceImpl{
		RoleRepository: RoleRepository,
		DB:             db,
	}
}

func (s *RoleServiceImpl) GetRoleList(filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	get, err := s.RoleRepository.GetRoleList(tx, filterCondition, pages)

	if err != nil {
		return get, err
	}

	return get, nil
}
func (s *RoleServiceImpl) GetRoleById(id int) (entitypayloads.RolePayload, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	result, err := s.RoleRepository.GetRoleById(tx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *RoleServiceImpl) SaveRole(req entitypayloads.RoleRequest) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	results, err := s.RoleRepository.SaveRole(tx, req)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (s *RoleServiceImpl) ChangeStatusRole(id int) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, err := s.RoleRepository.GetRoleById(tx, id)

	if err != nil {
		// panic(exceptions.NewNotFoundError(err.Error()))
		return false, err
	}

	results, err := s.RoleRepository.ChangeStatusRole(tx, id)
	if err != nil {
		return results, err
	}
	return true, nil
}

func (s *RoleServiceImpl) UpdateRole(req entitypayloads.RoleUpdate) (bool, *exceptions.BaseErrorResponse) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	results, err := s.RoleRepository.UpdateRole(tx, req)
	if err != nil {
		return results, err
	}

	return results, nil
}

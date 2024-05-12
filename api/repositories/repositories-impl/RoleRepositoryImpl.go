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

type RoleRepositoryImpl struct {
}

func StartRoleRepositoryImpl() repository.RoleRepository {
	return &RoleRepositoryImpl{}
}

func (*RoleRepositoryImpl) GetRoleById(tx *gorm.DB, Id int) (entitypayloads.RolePayload, *exceptions.BaseErrorResponse) {
	var RoleMapping entities.Role
	var RoleResponse entitypayloads.RolePayload

	rows, err := tx.
		Model(&RoleMapping).
		Where(entities.Role{RoleId: Id}).
		First(&RoleResponse).
		Rows()

	if err != nil {

		return RoleResponse, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}
	defer rows.Close()

	return RoleResponse, nil
}

func (r *RoleRepositoryImpl) SaveRole(tx *gorm.DB, req entitypayloads.RoleRequest) (bool, *exceptions.BaseErrorResponse) {
	entities := entities.Role{
		RoleName:        req.RoleName,
		RoleDescription: req.RoleDescription,
		ActiveStatus:    true,
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

func (r *RoleRepositoryImpl) ChangeStatusRole(tx *gorm.DB, Id int) (bool, *exceptions.BaseErrorResponse) {
	var entities entities.Role

	result := tx.Model(&entities).
		Where("role_id = ?", Id).
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

func (r *RoleRepositoryImpl) GetRoleList(tx *gorm.DB, filterCondition []utils.FilterCondition, pages pagination.Pagination) (pagination.Pagination, *exceptions.BaseErrorResponse) {
	RoleMapping := []entities.Role{}
	RoleResponse := []entitypayloads.RolePayload{}
	query := tx.
		Model(entities.Role{}).
		Scan(&RoleResponse)

	ApplyFilter := utils.ApplyFilter(query, filterCondition)

	err := ApplyFilter.
		Scopes(pagination.Paginate(&RoleMapping, &pages, ApplyFilter)).
		// Order("").
		Scan(&RoleResponse).
		Error

	if len(RoleResponse) == 0 {
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
	pages.Rows = RoleResponse

	return pages, nil
}

func (r *RoleRepositoryImpl) UpdateRole(tx *gorm.DB, req entitypayloads.RoleUpdate) (bool, *exceptions.BaseErrorResponse) {
	var model entities.Role
	if err := tx.Where("role_id = ?", req.RoleId).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusNotFound,
				Err:        errors.New("role not found"),
			}
		}
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	model.RoleName = req.RoleName
	model.RoleDescription = req.RoleDescription

	if err := tx.Save(&model).Error; err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
	}

	return true, nil
}

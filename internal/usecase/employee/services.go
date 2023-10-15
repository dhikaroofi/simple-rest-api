package employee

import (
	"context"
	"github.com/dhikaroofi/simple-rest-api/internal/common"
	"gorm.io/gorm"
	"strconv"
)

type EmployeeServices interface {
	Get(ctx context.Context, id string) (resp EmployeeResp, err error)
	GetList(ctx context.Context, queryCommon *common.QueryPagination) (resp ListEmployeeResp, err error)
	Create(ctx context.Context, req CreateOrUpdateEmployeeReq) (resp EmployeeResp, err error)
	Update(ctx context.Context, id string, req CreateOrUpdateEmployeeReq) (resp EmployeeResp, err error)
	Delete(ctx context.Context, id string) (err error)
}

type employeeServices struct {
	repo EmployeeRepo
}

func NewEmployeeServices(dbClient *gorm.DB) EmployeeServices {
	return &employeeServices{
		repo: NewEmployeeRepo(dbClient),
	}
}

func (s employeeServices) Get(ctx context.Context, id string) (resp EmployeeResp, err error) {
	result, err := s.repo.Get(ctx, id)
	if err != nil {
		return
	}

	resp = EmployeeResp{result}

	return
}

func (s employeeServices) GetList(ctx context.Context, queryCommon *common.QueryPagination) (resp ListEmployeeResp, err error) {
	result, err := s.repo.GetList(ctx, queryCommon)
	if err != nil {
		return
	}

	resp = ListEmployeeResp{
		List:       result,
		Pagination: common.GetPaginationFromQuery(queryCommon),
	}

	return
}

func (s employeeServices) Create(ctx context.Context, req CreateOrUpdateEmployeeReq) (resp EmployeeResp, err error) {
	ent := req.ConvertReqToEntity()
	if err = s.repo.Create(ctx, &ent); err != nil {
		return
	}

	resp = EmployeeResp{ent}
	return
}

func (s employeeServices) Update(ctx context.Context, id string, req CreateOrUpdateEmployeeReq) (resp EmployeeResp, err error) {
	ent := req.ConvertReqToEntity()

	if err = s.repo.CheckIfExist(ctx, id); err != nil {
		return
	}

	if err = s.repo.Update(ctx, id, &ent); err != nil {
		return
	}

	intID, _ := strconv.Atoi(id)
	ent.ID = intID

	resp = EmployeeResp{ent}
	return
}

func (s employeeServices) Delete(ctx context.Context, id string) (err error) {
	if err = s.repo.CheckIfExist(ctx, id); err != nil {
		return
	}

	if err = s.repo.Delete(ctx, id); err != nil {
		return
	}

	return
}

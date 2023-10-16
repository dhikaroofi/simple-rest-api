package employee

import (
	"context"
	"gorm.io/gorm"
	"strconv"

	"github.com/dhikaroofi/simple-rest-api/internal/common"
)

type ServicesInterfaces interface {
	Get(ctx context.Context, id string) (resp Resp, err error)
	GetList(ctx context.Context, queryCommon *common.QueryPagination) (resp ListEmployeeResp, err error)
	Create(ctx context.Context, req CreateOrUpdateEmployeeReq) (resp Resp, err error)
	Update(ctx context.Context, id string, req CreateOrUpdateEmployeeReq) (resp Resp, err error)
	Delete(ctx context.Context, id string) (err error)
}

type employeeServices struct {
	repo RepoInterfaces
}

func NewEmployeeServices(dbClient *gorm.DB) ServicesInterfaces {
	return &employeeServices{
		repo: NewEmployeeRepo(dbClient),
	}
}

func (s employeeServices) Get(ctx context.Context, id string) (resp Resp, err error) {
	result, err := s.repo.Get(ctx, id)
	if err != nil {
		return
	}

	resp = Resp{result}

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

func (s employeeServices) Create(ctx context.Context, req CreateOrUpdateEmployeeReq) (resp Resp, err error) {
	ent := req.ConvertReqToEntity()
	if err = s.repo.Create(ctx, &ent); err != nil {
		return
	}

	resp = Resp{ent}
	return
}

func (s employeeServices) Update(ctx context.Context, id string, req CreateOrUpdateEmployeeReq) (resp Resp, err error) {
	ent := req.ConvertReqToEntity()

	if err = s.repo.CheckIfExist(ctx, id); err != nil {
		return
	}

	if err = s.repo.Update(ctx, id, &ent); err != nil {
		return
	}

	intID, _ := strconv.Atoi(id)
	ent.ID = intID

	resp = Resp{ent}
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

package employee

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/dhikaroofi/simple-rest-api/internal/common"
	"github.com/dhikaroofi/simple-rest-api/pkg/customError"
)

type RepoInterfaces interface {
	Get(ctx context.Context, id string) (result Employee, err error)
	GetList(ctx context.Context, commonQuery *common.QueryPagination) (result []Employee, err error)
	Create(ctx context.Context, ent *Employee) (err error)
	Update(ctx context.Context, id string, ent *Employee) (err error)
	Delete(ctx context.Context, id string) (err error)
	CheckIfExist(ctx context.Context, id string) (err error)
}

type employeeRepo struct {
	dbClient *gorm.DB
	table    string
}

func NewEmployeeRepo(dbClient *gorm.DB) RepoInterfaces {
	return &employeeRepo{
		dbClient: dbClient,
		table:    "employees",
	}
}

func (r employeeRepo) GetList(ctx context.Context, commonQuery *common.QueryPagination) (result []Employee, err error) {
	exec := r.dbClient.Table(r.table).Select("id,first_name,last_name,email,TO_CHAR(hire_date, 'YYYY-MM-DD') AS hire_date")

	if commonQuery.SearchKeyword != "" {
		exec = exec.Where("first_name LIKE '%@val%' OR last_name LIKE '%@val%' OR email LIKE '%@val%'", sql.Named("val", commonQuery.SearchKeyword))
	}

	getCount := exec.Count(&commonQuery.TotalItems)
	if getCount.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	exec = exec.Scopes(paginate(commonQuery)).Find(&result)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	return
}

func (r employeeRepo) Get(ctx context.Context, id string) (result Employee, err error) {
	exec := r.dbClient.Table(r.table).Select("id,first_name,last_name,email,TO_CHAR(hire_date, 'YYYY-MM-DD') AS hire_date").Where("id=?", id).First(&result)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
			err = customError.ErrNotFound(r.table)
		}
		return
	}
	return
}

func (r employeeRepo) Create(ctx context.Context, ent *Employee) (err error) {
	exec := r.dbClient.Table(r.table).Create(ent)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to create employee"))
		return
	}

	return
}

func (r employeeRepo) Update(ctx context.Context, id string, ent *Employee) (err error) {
	exec := r.dbClient.Table(r.table).Where("id=?", id).Updates(ent)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to update employee"))
		return
	}

	return
}

func (r employeeRepo) Delete(ctx context.Context, id string) (err error) {
	ent := Employee{}
	exec := r.dbClient.Table(r.table).Where("id=?", id).Delete(&ent)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if exec.RowsAffected < 1 {
		err = customError.ErrQuery(fmt.Errorf("failed to delete employee"))
		return
	}

	return
}

func (r employeeRepo) CheckIfExist(ctx context.Context, id string) (err error) {
	var total int64
	exec := r.dbClient.Table(r.table).Where("id=?", id).Count(&total)
	if exec.Error != nil {
		err = customError.ErrQuery(exec.Error)
		return
	}

	if total < 1 {
		err = customError.ErrNotFound(r.table)
		return
	}

	return
}

func paginate(commonQuery *common.QueryPagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (commonQuery.Page - 1) * commonQuery.ItemsPerPage
		return db.Offset(offset).Limit(commonQuery.ItemsPerPage)
	}
}

package usecase

import "gorm.io/gorm"

type EmployeeRepo interface {
}

type employeeRepo struct {
	gorm *gorm.DB
}

func NewEmployeeRepo() EmployeeRepo {
	return &employeeRepo{}
}

package employee

import (
	"github.com/dhikaroofi/simple-rest-api/internal/common"
	"strings"
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	HireDate  string `json:"hire_date"`
}

type CreateOrUpdateEmployeeReq struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Email     string `json:"email" validate:"required,email"`
	HireDate  string `json:"hire_date" validate:"required,customDate"`
}

type EmployeeResp struct {
	Employee
}

type ListEmployeeResp struct {
	List []Employee `json:"list"`
	common.Pagination
}

func (r CreateOrUpdateEmployeeReq) ConvertReqToEntity() Employee {
	ent := Employee{}
	if val := strings.TrimSpace(r.FirstName); val != "" {
		ent.FirstName = val
	}

	if val := strings.TrimSpace(r.LastName); val != "" {
		ent.LastName = val
	}

	if val := strings.TrimSpace(r.Email); val != "" {
		ent.Email = val
	}

	if val := strings.TrimSpace(r.HireDate); val != "" {
		ent.HireDate = val
	}

	return ent
}

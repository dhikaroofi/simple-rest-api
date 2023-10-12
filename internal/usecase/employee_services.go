package usecase

type EmployeeServices interface {
	//Get(ctx context.Context)
	//GetList(ctx context.Context)
	//Create(ctx context.Context)
	//Update(ctx context.Context)
	//Delete(ctx context.Context)
}

type employeeServices struct {
	repo EmployeeRepo
}

func NewEmployeeServices() EmployeeServices {
	return &employeeServices{
		repo: NewEmployeeRepo(),
	}
}

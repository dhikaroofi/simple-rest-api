package restApi

import (
	"github.com/dhikaroofi/simple-rest-api/internal/presentation/restApi/common"
	"github.com/dhikaroofi/simple-rest-api/internal/usecase/employee"
	"github.com/gofiber/fiber/v2"
)

func (s *server) route() {
	s.app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(common.Response{
			Status:  200,
			Message: "hello world :)",
			Data:    struct{}{},
		})
	})

	routeGroupV1 := s.app.Group("/api/v1")

	routeGroupV1.Get("/employee/list", s.getListEmployee)
	routeGroupV1.Get("/employee/:id", s.getEmployee)
	routeGroupV1.Put("/employee/:id", s.actUpdate)
	routeGroupV1.Delete("/employee/remove/:id", s.actRemove)
	routeGroupV1.Post("/employee/create", s.actCreate)
}

func (s *server) getListEmployee(ctx *fiber.Ctx) error {
	resp, err := s.useCase.Employee.GetList(ctx.Context(), common.GetQueryPagination(ctx))
	if err != nil {
		return err
	}

	return common.ResponseOK(ctx, resp)
}

func (s *server) getEmployee(ctx *fiber.Ctx) error {
	id, err := common.GetValidIntID(ctx, s.validator)
	if err != nil {
		return err
	}

	resp, err := s.useCase.Employee.Get(ctx.Context(), id)
	if err != nil {
		return err
	}

	return common.ResponseOK(ctx, resp)
}

func (s *server) actCreate(ctx *fiber.Ctx) error {
	req := employee.CreateOrUpdateEmployeeReq{}

	if err := common.BindRequest(ctx, s.validator, &req); err != nil {
		return err
	}

	resp, err := s.useCase.Employee.Create(ctx.Context(), req)
	if err != nil {
		return err
	}

	return common.ResponseOK(ctx, resp)
}

func (s *server) actUpdate(ctx *fiber.Ctx) error {
	req := employee.CreateOrUpdateEmployeeReq{}

	if err := common.BindRequest(ctx, s.validator, &req); err != nil {
		return err
	}

	id, err := common.GetValidIntID(ctx, s.validator)
	if err != nil {
		return err
	}

	resp, err := s.useCase.Employee.Update(ctx.Context(), id, req)
	if err != nil {
		return err
	}

	return common.ResponseOK(ctx, resp)
}

func (s *server) actRemove(ctx *fiber.Ctx) error {
	id, err := common.GetValidIntID(ctx, s.validator)
	if err != nil {
		return err
	}

	err = s.useCase.Employee.Delete(ctx.Context(), id)
	if err != nil {
		return err
	}

	return common.ResponseOK(ctx, struct{}{})
}

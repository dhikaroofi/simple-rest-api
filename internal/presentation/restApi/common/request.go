package common

import (
	common2 "github.com/dhikaroofi/simple-rest-api/internal/common"
	"github.com/dhikaroofi/simple-rest-api/pkg/customError"
	validator2 "github.com/dhikaroofi/simple-rest-api/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func BindRequest(ctx *fiber.Ctx, validator *validator2.ValidationEngine, payload interface{}) error {
	if err := ctx.BodyParser(&payload); err != nil {
		err = customError.ErrBadRequest(err)
		return err
	}

	if err := Validate(validator, payload); err != nil {
		return err
	}

	return nil
}

func GetValidIntID(ctx *fiber.Ctx, validator *validator2.ValidationEngine) (id string, err error) {
	id = ctx.Params("id")

	payloadID := struct {
		ID string `json:"id" validate:"required,numeric"`
	}{ID: id}

	if err = Validate(validator, payloadID); err != nil {
		err = customError.ErrNotFound("")
		return
	}

	return
}

func GetQueryPagination(ctx *fiber.Ctx) *common2.QueryPagination {
	page := ctx.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}
	itemsPerPage := ctx.QueryInt("items_per_page", 10)
	if itemsPerPage < 1 {
		itemsPerPage = 10
	}

	return &common2.QueryPagination{
		Page:          page,
		ItemsPerPage:  itemsPerPage,
		SearchKeyword: ctx.Query("keyword"),
	}
}

package common

import (
	"github.com/dhikaroofi/simple-rest-api/pkg/customError"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`

	Error    string            `json:"error,omitempty"`
	ErrField map[string]string `json:"errorField,omitempty"`
}

func ResponseOK(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(Response{
		Status:  200,
		Message: "success",
		Data:    data,
	})
}

func ErrResponse(ctx *fiber.Ctx, err error) error {
	resp := Response{
		Status:  500,
		Message: customError.GeneralErrMessage,
		Error:   err.Error(),
		Data:    struct{}{},
	}

	if val, ok := err.(*customError.CustomError); ok {
		resp.Status = val.Code
		resp.Message = val.Message
		resp.Error = val.Error()
		if val.ErrField != nil {
			resp.ErrField = val.ErrField
		}
	}

	return ctx.JSON(resp)
}

package common

import (
	"fmt"
	"strings"

	"github.com/dhikaroofi/simple-rest-api/pkg/customError"
	validator2 "github.com/dhikaroofi/simple-rest-api/pkg/validator"

	"github.com/go-playground/validator/v10"
)

func Validate(validator *validator2.ValidationEngine, payload interface{}) error {
	if err := validator.Validator.Struct(payload); err != nil {
		return handleValidationErrors(err, validator)
	}
	return nil
}

func handleValidationErrors(err error, engine *validator2.ValidationEngine) error {
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok || len(validatorErrs) == 0 {
		return customError.ErrGeneral(fmt.Errorf("something went wrong in validation engine"))
	}

	errs := make(map[string]string, len(validatorErrs))
	for _, e := range validatorErrs {
		fieldName := strings.ToLower(e.Field())
		errMsg := e.Error()
		errMsg = e.Translate(engine.ENTranslator)
		errs[fieldName] = errMsg
	}

	return customError.ErrBadRequestFields(errs)
}

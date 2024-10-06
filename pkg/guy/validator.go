package guy

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ExtractValidationErrs(errs validator.ValidationErrors) error {
	var fields []string
	for _, e := range errs {
		fields = append(fields, fmt.Sprintf("%s is %s", e.Field(), e.ActualTag()))
	}
	return errors.New(strings.Join(fields, ", "))
}

func Validate[T any](req T) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		if err, ok := err.(validator.ValidationErrors); ok {
			return ExtractValidationErrs(err)
		}

		return err
	}

	return nil
}

package validator

import (
	v10 "github.com/go-playground/validator/v10"
)

var validator *v10.Validate

type V10Validator struct{}

func (v V10Validator) Struct(s interface{}) error {
	return validator.Struct(s)
}

func (v V10Validator) Initialize() {
	validator = v10.New()
}

package validator

import (
	"github.com/faruqisan/go-request-validation/validation"
	"gopkg.in/go-playground/validator.v9"
)

// Validator is an instance for validator
type Validator struct {
	validate *validator.Validate
}

// NewValidator create instance of validator
func NewValidator() validation.Validation {
	return &Validator{
		validator.New(),
	}
}

// ValidateStruct validate given struct
func (v *Validator) ValidateStruct(s interface{}) (err error) {

	err = v.validate.Struct(s)
	return
}

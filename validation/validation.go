package validation

// Validation Interface define contract for validator vendors
type Validation interface {
	ValidateStruct(s interface{}) error
}

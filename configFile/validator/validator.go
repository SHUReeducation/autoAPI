package validator

type RequireValidate interface {
	Validate() error
}

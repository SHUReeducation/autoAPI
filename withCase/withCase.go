package withCase

import "github.com/iancoleman/strcase"

type WithCase interface {
	PascalCase() string
	CamelCase() string
	SnakeCase() string
	KebabCase() string
}

type StringWithCase struct {
	s string
}

func (s StringWithCase) PascalCase() string {
	return strcase.ToCamel(s.s)
}

func (s StringWithCase) SnakeCase() string {
	return strcase.ToSnake(s.s)
}

func (s StringWithCase) CamelCase() string {
	return strcase.ToLowerCamel(s.s)
}

func (s StringWithCase) KebabCase() string {
	return strcase.ToKebab(s.s)
}

func New(s string) WithCase {
	return StringWithCase{s: s}
}

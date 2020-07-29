package withCase

import (
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v3"
)

type WithCase interface {
	PascalCase() string
	CamelCase() string
	SnakeCase() string
	KebabCase() string
}

type StringWithCase struct {
	string
}

func (s StringWithCase) PascalCase() string {
	return strcase.ToCamel(s.string)
}

func (s StringWithCase) SnakeCase() string {
	return strcase.ToSnake(s.string)
}

func (s StringWithCase) CamelCase() string {
	return strcase.ToLowerCamel(s.string)
}

func (s StringWithCase) KebabCase() string {
	return strcase.ToKebab(s.string)
}

func (s StringWithCase) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(s.string)
}

func (s *StringWithCase) UnmarshalYAML(node *yaml.Node) error {
	err := node.Decode(&s.string)
	return err
}

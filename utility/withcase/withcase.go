package withcase

import (
	"encoding/json"

	"github.com/iancoleman/strcase"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

type WithCase struct {
	string
}

func New(s string) WithCase {
	return WithCase{s}
}

func (s WithCase) PascalCase() string {
	return strcase.ToCamel(s.string)
}

func (s WithCase) SnakeCase() string {
	return strcase.ToSnake(s.string)
}

func (s WithCase) CamelCase() string {
	return strcase.ToLowerCamel(s.string)
}

func (s WithCase) KebabCase() string {
	return strcase.ToKebab(s.string)
}

func (s WithCase) MarshalYAML() (interface{}, error) {
	return yaml.Marshal(s.string)
}

func (s *WithCase) UnmarshalYAML(node *yaml.Node) error {
	err := node.Decode(&s.string)
	return err
}

func (s *WithCase) UnmarshalJSON(b []byte) error {
	var temp string
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	s.string = temp
	return nil
}

func (s *WithCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.string)
}

// MarshalTOML marshal `WithCase` to toml format value
func (s *WithCase) MarshalTOML() ([]byte, error) {
	return toml.Marshal(s.string)
}

// UnmarshalTOML unmarshal the value in toml format to `WithCase`
func (s *WithCase) UnmarshalTOML(b []byte) error {
	return toml.Unmarshal(b, s.string)
}

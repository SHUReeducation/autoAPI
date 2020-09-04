package field

import "autoAPI/utility/withcase"

type Field struct {
	Name withcase.WithCase `yaml:"name" json:"name" toml:"name"`
	Type string            `yaml:"type" json:"type" toml:"type"`
}

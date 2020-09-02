package field

import "autoAPI/utility/withCase"

type Field struct {
	Name withCase.WithCase `yaml:"name" json:"name" toml:"name"`
	Type string            `yaml:"type" json:"type" toml:"type"`
}

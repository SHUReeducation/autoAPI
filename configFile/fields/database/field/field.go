package field

import "autoAPI/utility/withCase"

type Field struct {
	Name withCase.WithCase `yaml:"name"`
	Type string            `yaml:"type"`
}

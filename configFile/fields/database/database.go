package database

import (
	"autoAPI/utility/withCase"
)

type Field struct {
	Name withCase.WithCase `yaml:"name"`
	Type string            `yaml:"type"`
}

type ComplexQueryResult struct {
	Array  bool    `yaml:"array"`
	Fields []Field `yaml:"fields"`
}

type ComplexQueryParam struct {
	// one and only one of OnThis and Name must be exists
	// only one OnThis can exist in a Complex
	OnThis *withCase.WithCase `yaml:"onThis"`
	Name   *withCase.WithCase `yaml:"name"`
	Type   string             `yaml:"type"`
}

type Complex struct {
	Name   withCase.WithCase   `yaml:"name"`
	SQL    string              `yaml:"sql"`
	Params []ComplexQueryParam `yaml:"params"`
	Result ComplexQueryResult  `yaml:"result"`
}

func (complex Complex) ForeignKey() *ComplexQueryParam {
	for _, param := range complex.Params {
		if param.OnThis != nil {
			return &param
		}
	}
	return nil
}

func (complex Complex) UseForeignKey() bool {
	for _, param := range complex.Params {
		if param.OnThis != nil {
			return true
		}
	}
	return false
}

type Table struct {
	Name    withCase.WithCase `yaml:"tablename"`
	Fields  []Field           `yaml:"fields"`
	Complex []Complex         `yaml:"complex"`
}

type Database struct {
	DBEngine string `yaml:"dbengine"`
	Table    Table  `yaml:",inline"`
}

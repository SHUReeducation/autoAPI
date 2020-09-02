package complex

import (
	"autoAPI/config/fields/database/field"
	"autoAPI/utility/withCase"
	"errors"
)

type QueryResult struct {
	Array  bool          `yaml:"array" json:"array" toml:"array"`
	Fields []field.Field `yaml:"fields" json:"fields" toml:"fields"`
}

type QueryParam struct {
	OnThis *withCase.WithCase `yaml:"onThis,omitempty" json:"onThis,omitempty" toml:"onThis,omitempty"`
	Name   *withCase.WithCase `yaml:"name,omitempty" json:"name,omitempty" toml:"name,omitempty"`
	Type   string             `yaml:"type" json:"type" toml:"type"`
}

type Complex struct {
	Name   withCase.WithCase `yaml:"name" json:"name" toml:"name"`
	SQL    string            `yaml:"sql" json:"sql" toml:"sql"`
	Params []QueryParam      `yaml:"params" json:"params" toml:"params"`
	Result QueryResult       `yaml:"result" json:"result" toml:"result"`
}

func (complex *Complex) Validate() error {
	alreadyMetOnThis := false
	for _, param := range complex.Params {
		if param.OnThis != nil {
			if alreadyMetOnThis {
				return errors.New("only one OnThis can exist in a Complex")
			} else {
				alreadyMetOnThis = true
			}
			if param.Name != nil {
				return errors.New("one and only one of OnThis and Name must be exists")
			}
		}
		if param.OnThis == nil && param.Name == nil {
			return errors.New("one and only one of OnThis and Name must be exists")
		}
	}
	return nil
}

func (complex Complex) ForeignKey() *QueryParam {
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

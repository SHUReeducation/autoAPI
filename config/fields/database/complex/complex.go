package complex

import (
	"autoAPI/config/fields/database/field"
	"autoAPI/utility/withCase"
	"errors"
)

type QueryResult struct {
	Array  bool          `yaml:"array" json:"array"`
	Fields []field.Field `yaml:"fields" json:"fields"`
}

type QueryParam struct {
	OnThis *withCase.WithCase `yaml:"onThis,omitempty" json:"onThis,omitempty"`
	Name   *withCase.WithCase `yaml:"name,omitempty" json:"name,omitempty"`
	Type   string             `yaml:"type" json:"type"`
}

type Complex struct {
	Name   withCase.WithCase `yaml:"name" json:"name"`
	SQL    string            `yaml:"sql" json:"sql"`
	Params []QueryParam      `yaml:"params" json:"params"`
	Result QueryResult       `yaml:"result" json:"result"`
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

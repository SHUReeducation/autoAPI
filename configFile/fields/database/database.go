package database

import (
	"autoAPI/configFile/fields/database/field"
	"autoAPI/dbAdapter"
	"autoAPI/dbAdapter/pgsql"
	"autoAPI/utility/withCase"
	"errors"
	_ "github.com/lib/pq"
)

type ComplexQueryResult struct {
	Array  bool          `yaml:"array" json:"array"`
	Fields []field.Field `yaml:"fields" json:"fields"`
}

type ComplexQueryParam struct {
	OnThis *withCase.WithCase `yaml:"onThis,omitempty" json:"onThis,omitempty"`
	Name   *withCase.WithCase `yaml:"name,omitempty" json:"name,omitempty"`
	Type   string             `yaml:"type" json:"type"`
}

type Complex struct {
	Name   withCase.WithCase   `yaml:"name" json:"name"`
	SQL    string              `yaml:"sql" json:"sql"`
	Params []ComplexQueryParam `yaml:"params" json:"params"`
	Result ComplexQueryResult  `yaml:"result" json:"result"`
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
	Name    *withCase.WithCase `yaml:"tablename" json:"tablename"`
	Fields  []field.Field      `yaml:"fields" json:"fields"`
	Complex []Complex          `yaml:"complex" json:"complex"`
}

func (t *Table) Validate() error {
	if t.Name == nil {
		return errors.New("table name must be given")
	}
	for _, c := range t.Complex {
		if err := c.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type Database struct {
	DBEngine *string `yaml:"dbengine" json:"dbengine"`
	Url      *string `yaml:"url" json:"url"`
	Table    *Table  `yaml:",inline" json:"table,inline"`
}

func FillWithDBAdapter(d *Database) error {
	if len(d.Table.Fields) == 0 && d.Url != nil {
		var adapter dbAdapter.DBAdapter
		switch *d.DBEngine {
		case "pgsql":
			adapter = pgsql.DBAdapter
		default:
			return errors.New("unsupported dbms")
		}
		fields, err := adapter.FillFields(*d.Url, *d.Table.Name)
		d.Table.Fields = fields
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) Validate() error {
	if d.DBEngine == nil {
		t := "pgsql"
		d.DBEngine = &t
	}
	if d.Table == nil {
		return errors.New("table information must be given")
	}
	if err := d.Table.Validate(); err != nil {
		return err
	}
	return FillWithDBAdapter(d)
}

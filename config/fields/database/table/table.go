package table

import (
	"autoAPI/config/fields/database/complex"
	"autoAPI/config/fields/database/field"
	"autoAPI/config/fields/database/sqlparser"
	"autoAPI/dbAdapter"
	"autoAPI/utility/withCase"
	"errors"
	"github.com/urfave/cli/v2"
)

type Table struct {
	Name    *withCase.WithCase `yaml:"tablename" json:"tablename"`
	Fields  []field.Field      `yaml:"fields" json:"fields"`
	Complex []complex.Complex  `yaml:"complex" json:"complex"`
}

func (table *Table) MergeWith(other *Table) {
	if other == nil {
		return
	}
	if table.Name == nil {
		table.Name = other.Name
	}
	if table.Fields == nil || len(table.Fields) == 0 {
		table.Fields = other.Fields
	}
	if table.Complex == nil || len(table.Complex) == 0 {
		table.Complex = other.Complex
	}
}

func (table *Table) Validate() error {
	if table.Name == nil {
		return errors.New("table name must be given")
	}
	for _, c := range table.Complex {
		if err := c.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func FromCommandLine(c *cli.Context) (*Table, error) {
	table := Table{}
	if name := c.String("tablename"); name != "" {
		n := withCase.New(name)
		table.Name = &n
	}
	return &table, nil
}

func (table *Table) MergeWithSQL(sqlFilePath string, dbms string) error {
	if table.Name == nil || len(table.Fields) == 0 {
		name, fields, err := sqlparser.ParseCreateTable(sqlFilePath, dbms)
		if err != nil {
			return err
		}
		table.Name = &name
		table.Fields = fields
	}
	return nil
}

func (table *Table) MergeWithDB(url string, dbEngine string) error {
	if len(table.Fields) == 0 {
		adapter, err := dbAdapter.New(dbEngine)
		if err != nil {
			return err
		}
		table.Fields, err = adapter.GetFields(url, *table.Name)
		return err
	}
	return nil
}

func (table *Table) MergeWithDefault() {
	for _, f := range table.Fields {
		if f.Name.CamelCase() == "id" {
			return
		}
	}
	table.Fields = append(
		[]field.Field{{Name: withCase.New("id"), Type: "bigserial"}},
		table.Fields...,
	)
}

func (table *Table) PrimaryKeyType() string {
	for _, f := range table.Fields {
		if f.Name.CamelCase() == "id" {
			return f.Type
		}
	}
	return ""
}

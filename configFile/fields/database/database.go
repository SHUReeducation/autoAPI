package database

import (
	"autoAPI/utility/withCase"
)

type Field struct {
	Name withCase.StringWithCase `yaml:"name"`
	Type string                  `yaml:"type"`
}

type Table struct {
	Name   withCase.StringWithCase `yaml:"tablename"`
	Fields []Field                 `yaml:"fields"`
}

type Database struct {
	DBEngine string `yaml:"dbengine"`
	Table    Table  `yaml:",inline"`
}

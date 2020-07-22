package table

import (
	"autoAPI/withCase"
)

type Field struct {
	Name       withCase.WithCase
	GoTypeName string
}

type Table struct {
	Name   withCase.WithCase
	Fields []Field
}

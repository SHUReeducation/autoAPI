package table

import (
	"autoAPI/typeTransform"
	"autoAPI/withCase"
	"autoAPI/yamlParser"
)

type Field struct {
	Name       withCase.WithCase
	GoTypeName string
}

type Table struct {
	Name   withCase.WithCase
	Fields []Field
}

func FromYaml(yamlFile yamlParser.YamlFile) Table {
	typeTransformer := typeTransform.GolangTypeTransformer{}
	table := Table{
		Name:   withCase.New(yamlFile.TableName),
		Fields: []Field{},
	}
	for _, field := range yamlFile.Fields {
		table.Fields = append(table.Fields, Field{
			Name:       withCase.New(field.Name),
			GoTypeName: typeTransformer.Transform(field.Type),
		})
	}
	return table
}

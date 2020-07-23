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

var importMap = map[string]string{
	"time.Time": "time",
}

func (f Field) ExtraImport() *string {
	result, contains := importMap[f.GoTypeName]
	if contains {
		return &result
	} else {
		return nil
	}
}

type Table struct {
	Name   withCase.WithCase
	Fields []Field
}

func (table Table) ExtraImport() []string {
	set := map[string]bool{}
	for _, field := range table.Fields {
		forField := field.ExtraImport()
		if forField != nil {
			set[*forField] = true
		}
	}
	var result []string
	for entry := range set {
		result = append(result, entry)
	}
	return result
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

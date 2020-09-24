package golang

import (
	apiIR "autoAPI/ir/api"
	"strings"
)

type Struct struct {
	Fields []Field
}

func lowStruct(s apiIR.Struct) Struct {
	var result Struct
	for _, f := range s.Fields {
		result.Fields = append(result.Fields, lowField(f))
	}
	return result
}

func (s Struct) PrimaryKey() Field {
	var result Field
	for _, f := range s.Fields {
		if f.Name.CamelCase() == "id" {
			result = f
		}
	}
	return result
}

func (s Struct) FieldsWithOutPrimaryKey() []Field {
	var result []Field
	for _, f := range s.Fields {
		if f.Name.CamelCase() != "id" {
			result = append(result, f)
		}
	}
	return result
}

func (s Struct) Imports() []string {
	var result []string
	for _, field := range s.Fields {
		importStr := field.Import()
		if importStr != "" {
			result = append(result, importStr)
		}
	}
	return result
}

func (s Struct) FieldNames() string {
	var fieldNames []string
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, field.Name.PascalCase())
	}
	return strings.Join(fieldNames, ", ")
}

func (s Struct) FieldNamesWithPrefix(prefix string) string {
	var fieldNames []string
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, prefix+field.Name.PascalCase())
	}
	return strings.Join(fieldNames, ", ")
}

func (s Struct) RepeatFieldTimes(content string) string {
	var fieldNames []string
	for range s.Fields {
		fieldNames = append(fieldNames, content)
	}
	return strings.Join(fieldNames, ", ")
}

func (s Struct) DBNamesString() string {
	var fieldNames []string
	for _, field := range s.Fields {
		fieldNames = append(fieldNames, field.Name.SnakeCase())
	}
	return strings.Join(fieldNames, ", ")
}

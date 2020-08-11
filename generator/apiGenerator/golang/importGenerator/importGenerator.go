package golangImportGenerator

import (
	"autoAPI/config"
	"autoAPI/config/fields/database/field"
	typeTransformer "autoAPI/generator/apiGenerator/golang/typeTransformer"
)

type importGenerator struct{}

var importMap = map[string]string{
	"time.Time": "time",
}

func importFor(f field.Field) *string {
	result, contains := importMap[typeTransformer.TypeTransformer.Transform(f.Type)]
	if contains {
		return &result
	} else {
		return nil
	}
}

func (i importGenerator) Generate(config config.Config) []string {
	set := map[string]bool{}
	for _, currentField := range config.Database.Table.Fields {
		forField := importFor(currentField)
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

var ImportGenerator importGenerator

func init() {
	ImportGenerator = importGenerator{}
}

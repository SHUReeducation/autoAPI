package golangImportGenerator

import (
	"autoAPI/configFile"
	"autoAPI/configFile/fields/database"
	typeTransformer "autoAPI/generator/apiGenerator/golang/typeTransformer"
)

type importGenerator struct{}

var importMap = map[string]string{
	"time.Time": "time",
}

func importFor(f database.Field) *string {
	result, contains := importMap[typeTransformer.TypeTransformer.Transform(f.Type)]
	if contains {
		return &result
	} else {
		return nil
	}
}

func (i importGenerator) Generate(configFile configFile.ConfigFile) []string {
	set := map[string]bool{}
	for _, field := range configFile.Database.Table.Fields {
		forField := importFor(field)
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

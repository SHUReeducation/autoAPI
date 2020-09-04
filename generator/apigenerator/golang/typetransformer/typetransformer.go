package typetransformer

import "strings"

type typeTransformer struct {
}

var transformMap = map[string]string{
	"smallint":    "int64",
	"integer":     "int64",
	"serial":      "int64",
	"bigint":      "int64",
	"bigserial":   "int64",
	"real":        "float64",
	"double":      "float64",
	"text":        "string",
	"date":        "time.Time",
	"time":        "time.Time",
	"timetz":      "time.Time",
	"timestamp":   "time.Time",
	"timestamptz": "time.Time",
	"boolean":     "bool",
	"bytea":       "[]byte",
}

func (typeTransformer) Transform(dataBaseType string) string {
	result, contains := transformMap[dataBaseType]

	switch {
	case contains:
		return result
	case strings.HasPrefix(dataBaseType, "char"), strings.HasPrefix(dataBaseType, "varchar"):
		return "string"
	}

	return dataBaseType
}

var TypeTransformer typeTransformer

func init() {
	TypeTransformer = typeTransformer{}
}

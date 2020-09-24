package api

import (
	"autoAPI/utility/withcase"
	"strings"
)

const (
	Int64 = iota
	Float64
	String
	// for languages which less shit than Golang, we need these
	Date
	Time
	DateTime

	Bool
	Binary
)

var trivialTypes = map[string]int{
	"int":         Int64,
	"smallint":    Int64,
	"integer":     Int64,
	"serial":      Int64,
	"bigint":      Int64,
	"bigserial":   Int64,
	"int4":        Int64,
	"int8":        Int64,
	"real":        Float64,
	"double":      Float64,
	"text":        String,
	"date":        Date,
	"time":        Time,
	"timetz":      Time,
	"timestamp":   Time,
	"ts":          Time,
	"timestamptz": Time,
	"tstz":        Time,
	"boolean":     Bool,
	"bytea":       Binary,
}

func lowType(dataBaseType string) interface{} {
	result, contains := trivialTypes[dataBaseType]
	switch {
	case contains:
		return result
	case strings.HasPrefix(dataBaseType, "char"), strings.HasPrefix(dataBaseType, "varchar"):
		return String
	case strings.HasSuffix(dataBaseType, "range"):
		return Struct{
			Fields: []Field{
				{Name: withcase.New("start"), Type: lowType(strings.ReplaceAll(dataBaseType, "range", ""))},
				{Name: withcase.New("end"), Type: lowType(strings.ReplaceAll(dataBaseType, "range", ""))},
			},
		}
	}
	return dataBaseType
}

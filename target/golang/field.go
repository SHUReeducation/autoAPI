package golang

import (
	apiIR "autoAPI/ir/api"
	"autoAPI/utility/withcase"
)

type Field struct {
	Name withcase.WithCase
	Type interface{} // interface{}: for trivial type, it's a string
	// 			 					 for struct type, it's a Struct
	// fuck golang's typesystem again
}

var nameMap = map[int]string{
	apiIR.Int64:    "int64",
	apiIR.Float64:  "float64",
	apiIR.String:   "string",
	apiIR.Date:     "time.Time",
	apiIR.Time:     "time.Time",
	apiIR.DateTime: "time.Time",
	apiIR.Bool:     "bool",
	apiIR.Binary:   "[]byte",
}

func lowField(field apiIR.Field) Field {
	if tn, ok := nameMap[field.Type.(int)]; ok {
		return Field{
			Name: field.Name,
			Type: tn,
		}
	} else {
		panic("todo: support complex types")
	}
}

var importMap = map[string]string{
	"time.Time": "time",
}

func (field Field) Import() string {
	return importMap[field.Type.(string)]
}

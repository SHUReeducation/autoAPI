package api

import (
	"autoAPI/config/fields/database"
	"autoAPI/config/fields/database/complex"
	"autoAPI/config/fields/database/field"
	"autoAPI/config/fields/database/table"
	"autoAPI/utility/withcase"
)

type Field struct {
	Name withcase.WithCase
	Type interface{} // interface{}: for trivial type, it's an enum, see ./type.go
	// 			 					 for struct type, it's a Struct
	// fuck golang's typesystem
}

type Struct struct {
	Fields []Field
}

func lowField(config field.Field) Field {
	return Field{
		Name: config.Name,
		Type: lowType(config.Type),
	}
}

func lowStruct(table table.Table) Struct {
	var result []Field
	for _, f := range table.Fields {
		lowed := lowField(f)
		if f.Inline {
			for _, subField := range lowed.Type.(Struct).Fields {
				result = append(result, subField)
			}
		} else {
			if _, ok := lowed.Type.(Struct); ok {
				panic("non-inlined complex type is not supported now!")
			}
			result = append(result, lowed)
		}
	}
	return Struct{result}
}

type API struct {
	Name     withcase.WithCase
	DBEngine string
	Model    Struct
	// todo: find a way to handle Complex better
	Complex []complex.Complex
}

func LowAPI(database database.Database) API {
	return API{
		Name:     *database.Table.Name,
		DBEngine: *database.DBEngine,
		Model:    lowStruct(*database.Table),
		Complex:  database.Table.Complex,
	}
}

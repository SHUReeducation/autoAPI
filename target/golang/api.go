package golang

import (
	"autoAPI/config/fields/database/complex"
	apiIR "autoAPI/ir/api"
	"autoAPI/utility/withcase"
)

type API struct {
	Name     withcase.WithCase
	DBEngine string
	Model    Struct
	Complex  []complex.Complex
}

func lowAPI(api apiIR.API) API {
	return API{
		Name:     api.Name,
		DBEngine: api.DBEngine,
		Model:    lowStruct(api.Model),
		Complex:  api.Complex,
	}
}

func (api API) GetDBEngineModURL() string {
	switch api.DBEngine {
	case "pgsql":
		return "github.com/lib/pq"
	case "mysql":
		return "github.com/go-sql-driver/mysql"
	default:
		panic("unsupported dbms")
	}
}

package dbAdapter

import (
	"autoAPI/configFile/fields/database/field"
	"autoAPI/utility/withCase"
)

type DBAdapter interface {
	FillFields(url string, tableName withCase.WithCase) ([]field.Field, error)
	// todo: use this to support multiple databases when generating code
	// RenderCreateSQL(configFile configFile.ConfigFile) string
	// RenderUpdateSQL(configFile configFile.ConfigFile) string
}

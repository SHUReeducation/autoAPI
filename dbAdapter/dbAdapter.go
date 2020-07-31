package dbAdapter

import (
	"autoAPI/configFile"
)

type DBAdapter interface {
	FillFields(config *configFile.ConfigFile) error
	// todo: use this to support multiple databases when generating code
	// RenderCreateSQL(configFile configFile.ConfigFile) string
	// RenderUpdateSQL(configFile configFile.ConfigFile) string
}

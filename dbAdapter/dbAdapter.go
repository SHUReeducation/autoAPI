package dbAdapter

import (
	"autoAPI/config/fields/database/field"
	"autoAPI/dbAdapter/pgsql"
	"autoAPI/utility/withCase"
	"errors"
)

type DBAdapter interface {
	GetFields(url string, tableName withCase.WithCase) ([]field.Field, error)
	// todo: use this to support multiple databases when generating code
	// RenderCreateSQL(config config.config) string
	// RenderUpdateSQL(config config.config) string
}

func New(dbms string) (DBAdapter, error) {
	switch dbms {
	case "pgsql":
		return pgsql.DBAdapter, nil
	default:
		return nil, errors.New("unsupported dbms")
	}
}

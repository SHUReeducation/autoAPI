package dbadapter

import (
	"errors"

	"autoAPI/config/fields/database/field"
	"autoAPI/dbadapter/pgsql"
	"autoAPI/utility/withcase"
)

type DBAdapter interface {
	GetFields(url string, tableName withcase.WithCase) ([]field.Field, error)
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

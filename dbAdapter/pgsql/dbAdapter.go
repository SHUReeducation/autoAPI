package pgsql

import (
	"autoAPI/configFile"
	"autoAPI/configFile/fields/database"
	"autoAPI/utility/withCase"
	"database/sql"
	_ "github.com/lib/pq"
)

type dbAdapter struct {
}

func (_ dbAdapter) FillFields(config *configFile.ConfigFile) error {
	db, err := sql.Open("postgres", *config.Database.Url)
	if err != nil {
		return err
	}
	rows, err := db.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = 'public'
  		AND table_name = $1;
	`, config.Database.Table.Name.CamelCase())
	if err != nil {
		return err
	}
	for rows.Next() {
		var field database.Field
		var nameStr, columnTypeStr string
		err = rows.Scan(&nameStr, &columnTypeStr)
		if err != nil {
			return err
		}
		field.Name = withCase.New(nameStr)
		field.Type = columnTypeStr
		if field.Name.CamelCase() != "id" {
			config.Database.Table.Fields = append(config.Database.Table.Fields, field)
		}
	}
	return nil
}

var DBAdapter dbAdapter

func init() {
	DBAdapter = dbAdapter{}
}

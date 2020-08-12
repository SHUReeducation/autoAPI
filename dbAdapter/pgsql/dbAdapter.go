package pgsql

import (
	"autoAPI/config/fields/database/field"
	"autoAPI/utility/withCase"
	"database/sql"
	_ "github.com/lib/pq"
)

type dbAdapter struct {
}

func (_ dbAdapter) GetFields(url string, tableName withCase.WithCase) ([]field.Field, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(`
		SELECT column_name, data_type
		FROM information_schema.columns
		WHERE table_schema = 'public'
  		AND table_name = $1;
	`, tableName.CamelCase())
	if err != nil {
		return nil, err
	}
	var result []field.Field
	for rows.Next() {
		var current field.Field
		var nameStr, columnTypeStr string
		err = rows.Scan(&nameStr, &columnTypeStr)
		if err != nil {
			return result, err
		}
		current.Name = withCase.New(nameStr)
		current.Type = columnTypeStr
		if current.Name.CamelCase() != "id" {
			result = append(result, current)
		}
	}
	return result, nil
}

var DBAdapter dbAdapter

func init() {
	DBAdapter = dbAdapter{}
}

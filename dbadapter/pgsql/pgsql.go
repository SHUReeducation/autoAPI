package pgsql

import (
	"database/sql"

	"autoAPI/config/fields/database/field"
	"autoAPI/utility/withcase"

	_ "github.com/lib/pq" // db driver
)

type dbAdapter struct {
}

func (dbAdapter) GetFields(url string, tableName withcase.WithCase) ([]field.Field, error) {
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
		current.Name = withcase.New(nameStr)
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

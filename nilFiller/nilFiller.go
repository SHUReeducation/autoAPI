package nilFiller

import (
	"autoAPI/configFile"
	"autoAPI/dbAdapter"
	"autoAPI/dbAdapter/pgsql"
	"errors"
)

func FillNil(config *configFile.ConfigFile) error {
	if len(config.Database.Table.Fields) == 0 && config.Database.Url != nil {
		var adapter dbAdapter.DBAdapter
		switch config.Database.DBEngine {
		case "pgsql":
			adapter = pgsql.DBAdapter
		default:
			return errors.New("unsupported dbms")
		}
		err := adapter.FillFields(config)
		if err != nil {
			return err
		}
	}
	return nil
}

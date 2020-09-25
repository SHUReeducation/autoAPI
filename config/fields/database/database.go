package database

import (
	"errors"
	"os"

	_ "github.com/lib/pq" // db driver
	"github.com/urfave/cli/v2"

	"autoAPI/config/fields/database/table"
)

type Database struct {
	DBEngine *string      `yaml:"dbengine" json:"dbengine" toml:"dbengine"`
	URL      *string      `yaml:"url" json:"url" toml:"url"`
	Table    *table.Table `yaml:",inline" json:",inline" toml:",inline"`
}

var Default = Database{}

func init() {
	dbEngine := "pgsql"
	Default.DBEngine = &dbEngine
}

func (database *Database) MergeWith(other *Database) {
	if other == nil {
		return
	}
	if database.DBEngine == nil {
		database.DBEngine = other.DBEngine
	}
	if database.URL == nil {
		database.URL = other.URL
	}
	if database.Table == nil {
		database.Table = other.Table
	} else {
		database.Table.MergeWith(other.Table)
	}
}

func (database *Database) MergeWithDefault() error {
	if database.DBEngine == nil || *database.DBEngine == "" {
		t := "pgsql"
		database.DBEngine = &t
	}
	database.Table.MergeWithDefault()
	return nil
}

func (database *Database) Validate() error {
	if database.Table == nil {
		return errors.New("table information must be given")
	}
	if err := database.Table.Validate(); err != nil {
		return err
	}
	if len(database.Table.Fields) == 0 {
		return errors.New("table fields must be given")
	}
	return nil
}

func FromCommandLine(c *cli.Context) (*Database, error) {
	var result Database
	if url := c.String("dburl"); url != "" {
		if url[:len("pgsql")] == "pgsql" || url[:len("postgres")] == "postgres" || url[:len("postgresql")] == "postgresql" {
			t := "pgsql"
			result.DBEngine = &t
		} else if url[:len("mysql")] == "mysql" {
			t := "mysql"
			result.DBEngine = &t
		}
		result.URL = &url
	}
	if dbengine := c.String("dbengine"); dbengine != "" {
		t := dbengine
		result.DBEngine = &t
	}
	var err error
	result.Table, err = table.FromCommandLine(c)
	return &result, err
}

func (database *Database) MergeWithEnv() error {
	if database.URL == nil {
		if url := os.Getenv("DB_ADDRESS"); url != "" {
			if url[:len("pgsql")] == "pgsql" || url[:len("postgres")] == "postgres" || url[:len("postgresql")] == "postgresql" {
				t := "pgsql"
				database.DBEngine = &t
			} else if url[:len("mysql")] == "mysql" {
				t := "mysql"
				database.DBEngine = &t
			}
			database.URL = &url
		}
	}
	return nil
}

func (database *Database) MergeWithSQL(sqlFilePath string) error {
	if database.DBEngine == nil || *database.DBEngine == "" {
		t := "pgsql"
		database.DBEngine = &t
	}
	return database.Table.MergeWithSQL(sqlFilePath, *database.DBEngine)
}

func (database *Database) MergeWithDB() error {
	if database.DBEngine != nil && database.URL != nil {
		return database.Table.MergeWithDB(*database.URL, *database.DBEngine)
	}
	return nil
}

func (database *Database) GetDBEngine() string {
	switch *database.DBEngine {
	case "pgsql":
		return "postgres"
	default:
		return *database.DBEngine
	}
}

package database

import (
	"autoAPI/config/fields/database/table"
	"errors"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"os"
)

type Database struct {
	DBEngine *string      `yaml:"dbengine" json:"dbengine" toml:"dbengine"`
	Url      *string      `yaml:"url" json:"url" toml:"url"`
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
	if database.Url == nil {
		database.Url = other.Url
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
		result.Url = &url
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
	if database.Url == nil {
		if url := os.Getenv("DB_ADDRESS"); url != "" {
			if url[:len("pgsql")] == "pgsql" || url[:len("postgres")] == "postgres" || url[:len("postgresql")] == "postgresql" {
				t := "pgsql"
				database.DBEngine = &t
			} else if url[:len("mysql")] == "mysql" {
				t := "mysql"
				database.DBEngine = &t
			}
			database.Url = &url
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
	if database.DBEngine != nil && database.Url != nil {
		return database.Table.MergeWithDB(*database.Url, *database.DBEngine)
	}
	return nil
}

func (database *Database) GetDBEngineModUrl() string {
	switch *database.DBEngine {
	case "pgsql":
		return "github.com/lib/pq"
	case "mysql":
		return "github.com/go-sql-driver/mysql"
	default:
		panic("unsupported dbms")
	}
}

func (database *Database) GetDBEngine() string {
	switch *database.DBEngine {
	case "pgsql":
		return "postgres"
	default:
		return *database.DBEngine
	}
}

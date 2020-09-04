package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"autoAPI/config/fields/cicd"
	"autoAPI/config/fields/database"
	"autoAPI/config/fields/docker"
)

type Config struct {
	Docker   *docker.Docker     `yaml:"docker" json:"docker" toml:"docker"`
	CICD     *cicd.CICD         `yaml:"cicd" json:"cicd" toml:"cicd"`
	Database *database.Database `yaml:"database" json:"database" toml:"database"`
}

func FromCommandLine(c *cli.Context) (*Config, error) {
	var result Config
	var err error
	result.Database, err = database.FromCommandLine(c)
	if err != nil {
		return nil, err
	}
	result.Docker, err = docker.FromCommandLine(c)
	if err != nil {
		return nil, err
	}
	result.CICD, err = cicd.FromCommandLine(c)
	return &result, err
}

func (c *Config) MergeWithEnv() error {
	if err := c.Database.MergeWithEnv(); err != nil {
		return err
	}
	if c.Docker != nil {
		return c.Docker.MergeWithEnv()
	}
	return nil
}

func (c *Config) MergeWithConfig(path string) error {
	fromFile, err := FromConfigFile(path)
	if err != nil {
		return err
	}
	c.MergeWith(fromFile)
	return nil
}

func (c *Config) MergeWithSQL(sqlFilePath string) error {
	return c.Database.MergeWithSQL(sqlFilePath)
}

func (c *Config) MergeWithDB() error {
	return c.Database.MergeWithDB()
}

func (c *Config) MergeWithDefault() error {
	if c.CICD != nil {
		if err := c.CICD.MergeWithDefault(); err != nil {
			return err
		}
	}
	if c.Database != nil {
		if err := c.Database.MergeWithDefault(); err != nil {
			return err
		}
	}
	return nil
}

func FromYaml(data []byte) (*Config, error) {
	var result Config
	err := yaml.Unmarshal(data, &result)
	return &result, err
}

func FromJSON(data []byte) (*Config, error) {
	var result Config
	err := json.Unmarshal(data, &result)
	return &result, err
}

// FromToml scan to config from toml format content
func FromToml(data []byte) (*Config, error) {
	var result Config
	err := toml.Unmarshal(data, &result)
	return &result, err
}

// FromConfigFile scan to config from file content
// only support json/yml/yaml and toml now
func FromConfigFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".json":
		return FromJSON(content)
	case ".yaml", ".yml":
		return FromYaml(content)
	case ".toml":
		return FromToml(content)
	default:
		return nil, errors.New("only support json/yml/yaml and toml now")
	}
}

func (c *Config) MergeWith(other *Config) {
	if c.Docker == nil {
		c.Docker = other.Docker
	} else {
		c.Docker.MergeWith(other.Docker)
	}
	if c.CICD == nil {
		c.CICD = other.CICD
	} else {
		c.CICD.MergeWith(other.CICD)
	}
	if c.Database == nil {
		c.Database = other.Database
	} else {
		c.Database.MergeWith(other.Database)
	}
}

func (c *Config) Validate() error {
	if (c.CICD != nil && c.CICD.K8s == nil && c.CICD.GithubAction == nil) || (c.Docker == nil || c.Docker.Username == nil || c.Docker.Tag == nil) {
		c.CICD = nil
	}
	err := c.Database.Validate()
	if err != nil {
		return err
	}
	return err
}

package config

import (
	"autoAPI/config/fields/k8s"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"autoAPI/config/fields/database"
	"autoAPI/config/fields/docker"
)

type Config struct {
	Docker       *docker.Docker     `yaml:"docker" json:"docker"`
	Database     *database.Database `yaml:"database" json:"database"`
	GitHubAction bool               `yaml:"GitHubAction" json:"GitHubAction"`
	K8s          *k8s.K8s           `yaml:"k8s" json:"k8s"`
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
	result.GitHubAction = c.Bool("ghaction")
	result.K8s, err = k8s.FromCommandLine(c)
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
	if c.GitHubAction == false {
		c.GitHubAction = other.GitHubAction
	}
	if c.K8s != nil {
		c.K8s.MergeWith(other.K8s)
	}
	if c.Database == nil {
		c.Database = other.Database
	} else {
		c.Database.MergeWith(other.Database)
	}
}

func (c *Config) Validate() error {
	if c.K8s != nil && c.K8s.Uri == nil && c.K8s.Host == nil {
		c.K8s = nil
	}
	err := c.Database.Validate()
	if c.K8s != nil && c.Database.URL == nil {
		err := fmt.Errorf("database url must be provided if want to generate k8s config")
		return err
	}
	if err != nil {
		return err
	}
	if c.K8s != nil && c.K8s.Uri == nil {
		uri := "/api/" + c.Database.Table.Name.KebabCase()
		c.K8s.Uri = &uri
	}
	return nil
}

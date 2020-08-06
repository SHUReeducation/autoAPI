package configFile

import (
	"autoAPI/configFile/fields/cicd"
	"autoAPI/configFile/fields/database"
	"autoAPI/configFile/fields/docker"
	"autoAPI/configFile/parser"
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type ConfigFile struct {
	Docker   *docker.Docker     `yaml:"docker" json:"docker"`
	CICD     *cicd.CICD         `yaml:"cicd" json:"cicd"`
	Database *database.Database `yaml:"database" json:"database"`
}

func FromYaml(data []byte) (ConfigFile, error) {
	var result ConfigFile
	err := yaml.Unmarshal(data, &result)
	if result.Database.CreateSql != nil &&
		(result.Database.Table.Fields == nil || result.Database.Table.Name == nil) {
		err = parser.FillTableInfo(*result.Database.CreateSql, result.Database.Table)
	}
	return result, err
}

func FromJson(data []byte) (ConfigFile, error) {
	var result ConfigFile
	err := json.Unmarshal(data, &result)
	return result, err
}

func LoadConfigFile(path string) (ConfigFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return ConfigFile{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return ConfigFile{}, err
	}
	switch ext := strings.ToLower(filepath.Ext(path)); ext {
	case ".json":
		return FromJson(content)
	case ".yaml", ".yml":
		return FromYaml(content)
	default:
		return ConfigFile{}, errors.New("only support json or yaml now")
	}
}

func (c *ConfigFile) Validate() error {
	err := c.Database.Validate()
	if err != nil {
		return err
	}
	if c.CICD != nil {
		err = c.CICD.Validate()
	}
	return err
}

package configFile

import (
	"autoAPI/configFile/fields/cicd"
	"autoAPI/configFile/fields/database"
	"autoAPI/configFile/fields/docker"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type ConfigFile struct {
	Docker   *docker.Docker     `yaml:"docker"`
	CICD     *cicd.CICD         `yaml:"cicd"`
	Database *database.Database `yaml:"database"`
}

func FromYaml(data []byte) (ConfigFile, error) {
	var result ConfigFile
	err := yaml.Unmarshal(data, &result)
	return result, err
}

func LoadYaml(path string) (ConfigFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return ConfigFile{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return ConfigFile{}, err
	}
	return FromYaml(content)
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

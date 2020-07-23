package yamlParser

import (
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type YamlFile struct {
	TableName string `yaml:"tablename"`
	Fields    []struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}
}

func New(data []byte) (YamlFile, error) {
	var result YamlFile
	err := yaml.Unmarshal(data, &result)
	return result, err
}

func Load(path string) (YamlFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return YamlFile{}, err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return YamlFile{}, err
	}
	return New(content)
}

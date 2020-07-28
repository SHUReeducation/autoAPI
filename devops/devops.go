package devops

import "autoAPI/yamlParser"

type MetaData struct {
	Version        string
	Dbengine       string
	DockerUsername string
}

func FromYaml(yamlFile yamlParser.YamlFile) MetaData {
	metaData := MetaData{
		Version:        yamlFile.Version,
		Dbengine:       yamlFile.DbEngine,
		DockerUsername: yamlFile.DockerUsername,
	}
	return metaData
}

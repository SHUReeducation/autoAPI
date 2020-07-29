package devops

import "autoAPI/yamlParser"

type MetaData struct {
	Version        string
	Dbengine       string
	GithubActions  string
	DockerUsername string
}

func FromYaml(yamlFile yamlParser.YamlFile) MetaData {
	metaData := MetaData{
		Version:        yamlFile.Version,
		Dbengine:       yamlFile.DbEngine,
		GithubActions:  yamlFile.GithubActions,
		DockerUsername: yamlFile.DockerUsername,
	}
	return metaData
}

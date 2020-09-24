package k8s

import (
	"autoAPI/config"
	"autoAPI/utility/withcase"
)

type K8s struct {
	Name           withcase.WithCase
	DockerUsername string
	DockerTag      string
	DBUrl          string
	Namespace      *string
	Host           *string // will not generate ingress when this is nil
	Uri            *string
}

func Low(config config.Config) K8s {
	return K8s{
		Name:           *config.Database.Table.Name,
		DockerUsername: *config.Docker.Username,
		DockerTag:      *config.Docker.Tag,
		DBUrl:          *config.Database.URL,
		Namespace:      config.K8s.Namespace,
		Host:           config.K8s.Host,
		Uri:            config.K8s.Uri,
	}
}

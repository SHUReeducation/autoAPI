package githubActions

import (
	"autoAPI/config"
	"autoAPI/utility/withcase"
)

type GitHubActions struct {
	DockerUsername string
	Tag            string
	Name           withcase.WithCase
}

func Low(config config.Config) GitHubActions {
	return GitHubActions{
		DockerUsername: *config.Docker.Username,
		Tag:            *config.Docker.Tag,
		Name:           *config.Database.Table.Name,
	}
}

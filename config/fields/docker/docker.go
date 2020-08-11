package docker

import (
	"github.com/urfave/cli/v2"
	"os"
)

type Docker struct {
	Username *string `yaml:"username" json:"username"`
	Tag      *string `yaml:"tag" json:"tag"`
}

func (docker *Docker) MergeWith(other *Docker) {
	if other == nil {
		return
	}
	if docker.Username == nil {
		docker.Username = other.Username
	}
	if docker.Tag == nil {
		docker.Tag = other.Tag
	}
}

func (docker *Docker) MergeWithEnv() error {
	if os.Getenv("DOCKER_USERNAME") != "" {
		username := os.Getenv("DOCKER_USERNAME")
		docker.Username = &username
	}
	if os.Getenv("DOCKER_TAG") != "" {
		tag := os.Getenv("DOCKER_TAG")
		docker.Tag = &tag
	}
	return nil
}

func FromCommandLine(c *cli.Context) (*Docker, error) {
	var result Docker
	if username := c.String("dockerusername"); username != "" {
		result.Username = &username
	}
	if tag := c.String("dockertag"); tag != "" {
		result.Tag = &tag
	}
	if result.Username == nil && result.Tag == nil {
		return nil, nil
	}
	return &result, nil
}

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
	if docker.Username == nil && os.Getenv("DOCKER_USERNAME") != "" {
		username := os.Getenv("DOCKER_USERNAME")
		docker.Username = &username
	}
	if docker.Tag == nil && os.Getenv("DOCKER_TAG") != "" {
		tag := os.Getenv("DOCKER_TAG")
		docker.Tag = &tag
	}
	if docker.Tag == nil && os.Getenv("GITHUB_RUN_NUMBER") != "" {
		tag := "ci-v" + os.Getenv("GITHUB_RUN_NUMBER")
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
	if c.Bool("nodocker") {
		return nil, nil
	}
	return &result, nil
}

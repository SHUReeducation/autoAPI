package ir

import (
	"autoAPI/config"
	"autoAPI/ir/api"
	"autoAPI/ir/dockerfile"
	"autoAPI/ir/githubActions"
	"autoAPI/ir/k8s"
)

type IR struct {
	API           api.API
	Dockerfile    *dockerfile.Dockerfile
	GitHubActions *githubActions.GitHubActions
	K8s           *k8s.K8s
}

func Low(config config.Config) IR {
	var ghAction *githubActions.GitHubActions
	var k *k8s.K8s
	action := githubActions.Low(config)
	ghAction = &action
	kube := k8s.Low(config)
	k = &kube
	var df *dockerfile.Dockerfile
	if config.Docker != nil {
		df = &dockerfile.Dockerfile{Name: *config.Database.Table.Name}
	}
	return IR{
		API:           api.LowAPI(*config.Database),
		Dockerfile:    df,
		GitHubActions: ghAction,
		K8s:           k,
	}
}

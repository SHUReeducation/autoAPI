package golang

import (
	"autoAPI/ir"
	"autoAPI/ir/dockerfile"
	"autoAPI/ir/githubActions"
	"autoAPI/ir/k8s"
)

type Target struct {
	// API is complex and needs to low from the ir
	API API
	// the other fields are simple, just reuse the ir
	Dockerfile    *dockerfile.Dockerfile
	GitHubActions *githubActions.GitHubActions
	K8s           *k8s.K8s
}

func Low(ir ir.IR) Target {
	return Target{
		API:           lowAPI(ir.API),
		Dockerfile:    ir.Dockerfile,
		GitHubActions: ir.GitHubActions,
		K8s:           ir.K8s,
	}
}

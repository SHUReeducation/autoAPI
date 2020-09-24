package generator

import (
	"autoAPI/ir/api"
	"autoAPI/ir/dockerfile"
	"autoAPI/ir/githubActions"
	"autoAPI/ir/k8s"
	"os"
	"path/filepath"

	ghTemplate "autoAPI/template/shared/githubActions"
	k8sTemplate "autoAPI/template/shared/k8s"
)

type Generator interface {
	GenerateAPI(apiConfig api.API, dirPath string) error
	GenerateDockerFile(dockerConfig dockerfile.Dockerfile, dirPath string) error
	GenerateGitHubActions(actions githubActions.GitHubActions, dirPath string) error
	GenerateK8s(k8s k8s.K8s, dirPath string) error
}

type Base struct {
}

func (b Base) GenerateAPI(_ api.API, _ string) error {
	panic("implement me")
}

func (b Base) GenerateDockerFile(_ dockerfile.Dockerfile, _ string) error {
	panic("implement me")
}

func (b Base) GenerateGitHubActions(actions githubActions.GitHubActions, dirPath string) error {
	githubActionDir := filepath.Join(dirPath, ".github")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionDir = filepath.Join(githubActionDir, "workflows")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionFile, err := os.Create(filepath.Join(githubActionDir, "dockerimage.yml"))
	if err != nil {
		return err
	}
	defer githubActionFile.Close()
	ghFileContent := ghTemplate.Render(actions)
	_, err = githubActionFile.WriteString(ghFileContent)
	return err
}

func (b Base) GenerateK8s(k8s k8s.K8s, dirPath string) error {
	kubernetesDeploymentFile, err := os.Create(filepath.Join(dirPath, k8s.Name.KebabCase()+".yaml"))
	if err != nil {
		return err
	}
	defer kubernetesDeploymentFile.Close()
	k8sFileContent := k8sTemplate.Render(k8s)
	_, err = kubernetesDeploymentFile.WriteString(k8sFileContent)
	return err
}

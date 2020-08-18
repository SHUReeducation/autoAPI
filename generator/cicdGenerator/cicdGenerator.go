package cicdGenerator

import (
	"autoAPI/config"
	"autoAPI/template/cicd"
	"os"
	"path/filepath"
)

type CICDGenerator struct{}

func renderKubernetesDeployment(config config.Config, dirPath string) error {
	kubernetesDeployment := cicd.KubernetesFile(config)
	kubernetesDeploymentFile, err := os.Create(filepath.Join(dirPath, config.Database.Table.Name.KebabCase()+".yaml"))
	if err != nil {
		return err
	}
	defer kubernetesDeploymentFile.Close()
	_, err = kubernetesDeploymentFile.WriteString(kubernetesDeployment)
	return err
}

func renderGitHubActions(config config.Config, dirPath string) error {
	githubActionDir := filepath.Join(dirPath, ".github")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionDir = filepath.Join(githubActionDir, "workflows")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionFileContent := cicd.GitHubActionDocker(config)
	githubActionFile, err := os.Create(filepath.Join(githubActionDir, "dockerimage.yml"))
	if err != nil {
		return err
	}
	defer githubActionFile.Close()
	_, err = githubActionFile.WriteString(githubActionFileContent)
	return err
}

func (generator CICDGenerator) Generate(config config.Config, dirPath string) error {
	if config.CICD.K8s == nil || *config.CICD.K8s {
		err := renderKubernetesDeployment(config, dirPath)
		if err != nil {
			return err
		}
	}
	if config.CICD.GithubAction == nil || *config.CICD.GithubAction {
		err := renderGitHubActions(config, dirPath)
		if err != nil {
			return err
		}
	}
	return nil
}

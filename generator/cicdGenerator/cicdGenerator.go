package cicdGenerator

import (
	"autoAPI/configFile"
	"autoAPI/template/cicd"
	"os"
	"path/filepath"
)

type DevopsGenerator struct{}

func renderKubernetesDeployment(configFile configFile.ConfigFile, dirPath string) error {

	kubernetesDeployment := cicd.KubernetesFile(configFile)
	kubernetesDeploymentFile, err := os.Create(filepath.Join(dirPath, configFile.Database.Table.Name.KebabCase()+".yaml"))
	if err != nil {
		return err
	}
	defer kubernetesDeploymentFile.Close()
	_, err = kubernetesDeploymentFile.WriteString(kubernetesDeployment)
	return err
}

func renderGitHubActions(configFile configFile.ConfigFile, dirPath string) error {
	githubActionDir := filepath.Join(dirPath, ".github")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionDir = filepath.Join(githubActionDir, "workflows")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionFileContent := cicd.GitHubActionDocker(configFile)
	githubActionFile, err := os.Create(filepath.Join(githubActionDir, "dockerimage.yml"))
	if err != nil {
		return err
	}
	defer githubActionFile.Close()
	_, err = githubActionFile.WriteString(githubActionFileContent)
	return err
}

func (generator DevopsGenerator) Generate(configFile configFile.ConfigFile, dirPath string) error {
	err := renderKubernetesDeployment(configFile, dirPath)
	if err != nil {
		return err
	}
	err = renderGitHubActions(configFile, dirPath)
	return err
}

package generator

import (
	"autoAPI/devops"
	"autoAPI/table"
	"autoAPI/template/general"
	"os"
	"path/filepath"
)

type DevopsGenerator struct{}

func renderKubernetesDeployment(metaData devops.MetaData, table table.Table, dirPath string) error {
	devopsDir := filepath.Join(dirPath, "devops")
	if err := os.Mkdir(devopsDir, 0755); err != nil {
		return err
	}
	kubernetesDeployment := general.KubernetesFile(metaData, table)
	kubernetesDeploymentFile, err := os.Create(filepath.Join(devopsDir, table.Name.KebabCase()+"-deployment.yaml"))
	if err != nil {
		return err
	}
	defer kubernetesDeploymentFile.Close()
	_, err = kubernetesDeploymentFile.WriteString(kubernetesDeployment)
	return err
}

func renderGitHubActions(metaData devops.MetaData, table table.Table, dirPath string) error {
	githubActionDir := filepath.Join(dirPath, ".github")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionDir = filepath.Join(githubActionDir, "workflow")
	if err := os.Mkdir(githubActionDir, 0755); err != nil {
		return err
	}
	githubActionFileContent := general.GitHubActionDocker(metaData, table)
	dbFile, err := os.Create(filepath.Join(githubActionDir, "dockerimage.yml"))
	if err != nil {
		return err
	}
	defer dbFile.Close()
	_, err = dbFile.WriteString(githubActionFileContent)
	return err
}

func (generator DevopsGenerator) Generate(metaData devops.MetaData, table table.Table, dirPath string) error {
	err := renderKubernetesDeployment(metaData, table, dirPath)
	if err != nil {
		return err
	}
	if metaData.GithubActions != "false" {
		err = renderGitHubActions(metaData, table, dirPath)
		return err
	}
	return nil
}

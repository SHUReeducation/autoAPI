package githubActionsGenerator

import (
	"autoAPI/config"
	"autoAPI/template/cicd"
	"os"
	"path/filepath"
)

type GitHubActionsGenerator struct{}

func (_ GitHubActionsGenerator) Generate(config config.Config, dirPath string) error {
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

package k8sGenerator

import (
	"autoAPI/config"
	"autoAPI/template/cicd"
	"os"
	"path/filepath"
)

type K8sGenerator struct{}

func (_ K8sGenerator) Generate(config config.Config, dirPath string) error {
	k8sDeployment := cicd.KubernetesFile(config)
	k8sDeploymentFile, err := os.Create(filepath.Join(dirPath, config.Database.Table.Name.KebabCase()+".yaml"))
	if err != nil {
		return err
	}
	defer k8sDeploymentFile.Close()
	_, err = k8sDeploymentFile.WriteString(k8sDeployment)
	return err
}

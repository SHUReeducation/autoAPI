package golang

import (
	"autoAPI/generator"
	"autoAPI/ir/dockerfile"
	api "autoAPI/target/golang"
	dockerfileTemplate "autoAPI/template/go/dockerfile"
	"autoAPI/template/go/goMod"
	handlerTemplate "autoAPI/template/go/handler"
	"autoAPI/template/go/infrastructure"
	"autoAPI/template/go/mainTemplate"
	mysqlModelTemplate "autoAPI/template/go/model/mysql"
	pgsqlModelTemplate "autoAPI/template/go/model/pgsql"
	"fmt"
	"os"
	"path/filepath"
)

type Generator struct {
	generator.Base
}

func (b Generator) GenerateAPI(apiConfig api.API, dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	if err = os.Mkdir(dirPath, 0755); err != nil {
		return err
	}
	if err = renderDB(apiConfig, dirPath); err != nil {
		return err
	}
	if err = renderModel(apiConfig, dirPath); err != nil {
		return err
	}
	if err = renderHandler(apiConfig, dirPath); err != nil {
		return err
	}
	if err = renderMain(apiConfig, dirPath); err != nil {
		return err
	}
	return renderGoMod(apiConfig, dirPath)
}

func renderMain(apiConfig api.API, dirPath string) error {
	mainFileContent := mainTemplate.Render(apiConfig)
	mainFile, err := os.Create(filepath.Join(dirPath, "main.go"))
	if err != nil {
		return err
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContent)
	return err
}

func renderHandler(apiConfig api.API, dirPath string) error {
	handlerDir := filepath.Join(dirPath, "handler")
	err := os.Mkdir(handlerDir, 0755)
	if err != nil {
		return err
	}
	handlerFileContent := handlerTemplate.Render(apiConfig)
	handlerFile, err := os.Create(filepath.Join(handlerDir, "handler.go"))
	if err != nil {
		return err
	}
	defer handlerFile.Close()
	_, err = handlerFile.WriteString(handlerFileContent)
	return err
}

func renderModel(apiConfig api.API, dirPath string) error {
	modelDir := filepath.Join(dirPath, "model")
	err := os.Mkdir(modelDir, 0755)
	if err != nil {
		return err
	}
	var modelFileContent string
	// todo: use inherit instead of if
	fmt.Println("    Use database engine", apiConfig.DBEngine)
	if apiConfig.DBEngine == "mysql" {
		modelFileContent = mysqlModelTemplate.Render(apiConfig)
	} else if apiConfig.DBEngine == "pgsql" {
		modelFileContent = pgsqlModelTemplate.Render(apiConfig)
	}
	modelFile, err := os.Create(filepath.Join(modelDir, "model.go"))
	if err != nil {
		return err
	}
	defer modelFile.Close()
	_, err = modelFile.WriteString(modelFileContent)
	return err
}

func renderDB(apiConfig api.API, dirPath string) error {
	infrastructureDir := filepath.Join(dirPath, "infrastructure")
	err := os.Mkdir(infrastructureDir, 0755)
	if err != nil {
		return err
	}
	dbFileContent := infrastructure.Render(apiConfig)
	dbFile, err := os.Create(filepath.Join(infrastructureDir, "db.go"))
	if err != nil {
		return err
	}
	defer dbFile.Close()
	_, err = dbFile.WriteString(dbFileContent)
	return err
}

func renderGoMod(apiConfig api.API, dirPath string) error {
	modFileContent := goMod.Render(apiConfig)
	modFile, err := os.Create(filepath.Join(dirPath, "go.mod"))
	if err != nil {
		return err
	}
	defer modFile.Close()
	_, err = modFile.WriteString(modFileContent)
	return err
}

func (b Generator) GenerateDockerFile(dockerConfig dockerfile.Dockerfile, dirPath string) error {
	fmt.Println("Generating Dockerfile")
	dockerFileContent := dockerfileTemplate.Render(dockerConfig)
	dockerfileFile, err := os.Create(filepath.Join(dirPath, "Dockerfile"))
	if err != nil {
		return err
	}
	defer dockerfileFile.Close()
	_, err = dockerfileFile.WriteString(dockerFileContent)
	return err
}

func (b Generator) Generate(targetConfig api.Target, dirPath string) error {
	fmt.Println("Generating API for Golang")
	if err := b.GenerateAPI(targetConfig.API, dirPath); err != nil {
		return err
	}
	if targetConfig.Dockerfile != nil {
		if err := b.GenerateDockerFile(*targetConfig.Dockerfile, dirPath); err != nil {
			return err
		}
	}
	if targetConfig.GitHubActions != nil {
		if err := b.GenerateGitHubActions(*targetConfig.GitHubActions, dirPath); err != nil {
			return err
		}
	}
	if targetConfig.K8s != nil {
		if err := b.GenerateK8s(*targetConfig.K8s, dirPath); err != nil {
			return err
		}
	}
	return nil
}

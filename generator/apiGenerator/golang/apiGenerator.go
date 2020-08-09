package golang

import (
	"autoAPI/configFile"
	"autoAPI/template/go/dockerfile"
	"autoAPI/template/go/goMod"
	"autoAPI/template/go/handler"
	"autoAPI/template/go/infrastructure"
	"autoAPI/template/go/mainTemplate"
	"autoAPI/template/go/model"
	"os"
	"path/filepath"
)

type APIGenerator struct{}

func renderDockerfile(configFile configFile.ConfigFile, dirPath string) error {
	dockerFileContent := dockerfile.Dockerfile(configFile)
	dockerfileFile, err := os.Create(filepath.Join(dirPath, "Dockerfile"))
	if err != nil {
		return err
	}
	defer dockerfileFile.Close()
	_, err = dockerfileFile.WriteString(dockerFileContent)
	return err
}

func renderGoMod(configFile configFile.ConfigFile, dirPath string) error {
	modFileContent := goMod.GoMod(configFile)
	modFile, err := os.Create(filepath.Join(dirPath, "go.mod"))
	if err != nil {
		return err
	}
	defer modFile.Close()
	_, err = modFile.WriteString(modFileContent)
	return err
}

func renderMain(configFile configFile.ConfigFile, dirPath string) error {
	mainFileContent := mainTemplate.MainTemplate(configFile)
	mainFile, err := os.Create(filepath.Join(dirPath, "main.go"))
	if err != nil {
		return err
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContent)
	return err
}

func renderHandler(configFile configFile.ConfigFile, dirPath string) error {
	handlerDir := filepath.Join(dirPath, "handler")
	err := os.Mkdir(handlerDir, 0755)
	if err != nil {
		return err
	}
	handlerFileContent := handler.Handler(configFile)
	handlerFile, err := os.Create(filepath.Join(handlerDir, "handler.go"))
	if err != nil {
		return err
	}
	defer handlerFile.Close()
	_, err = handlerFile.WriteString(handlerFileContent)
	return err
}

func renderModel(configFile configFile.ConfigFile, dirPath string) error {
	modelDir := filepath.Join(dirPath, "model")
	err := os.Mkdir(modelDir, 0755)
	if err != nil {
		return err
	}
	var modelFileContent string
	if configFile.Database.DBEngine != nil {
		if configFile.Database.GetDBEngine() == "mysql" {
			modelFileContent = model.ModelQ(configFile)
		} else {
			modelFileContent = model.Model(configFile)
		}
	}
	modelFile, err := os.Create(filepath.Join(modelDir, "model.go"))
	if err != nil {
		return err
	}
	defer modelFile.Close()
	_, err = modelFile.WriteString(modelFileContent)
	return err
}

func renderDB(configFile configFile.ConfigFile, dirPath string) error {
	infrastructureDir := filepath.Join(dirPath, "infrastructure")
	err := os.Mkdir(infrastructureDir, 0755)
	if err != nil {
		return err
	}
	dbFileContent := infrastructure.DB(configFile)
	dbFile, err := os.Create(filepath.Join(infrastructureDir, "db.go"))
	if err != nil {
		return err
	}
	defer dbFile.Close()
	_, err = dbFile.WriteString(dbFileContent)
	return err
}

func (generator APIGenerator) Generate(configFile configFile.ConfigFile, dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	if err = os.Mkdir(dirPath, 0755); err != nil {
		return err
	}
	if err = renderDB(configFile, dirPath); err != nil {
		return err
	}
	if err = renderModel(configFile, dirPath); err != nil {
		return err
	}
	if err = renderHandler(configFile, dirPath); err != nil {
		return err
	}
	if err = renderMain(configFile, dirPath); err != nil {
		return err
	}
	if err = renderGoMod(configFile, dirPath); err != nil {
		return err
	}
	if configFile.Docker != nil {
		err = renderDockerfile(configFile, dirPath)
	}
	return err
}

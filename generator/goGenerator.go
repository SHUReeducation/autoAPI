package generator

import (
	"autoAPI/table"
	"autoAPI/template/go/dockerfile"
	"autoAPI/template/go/goMod"
	"autoAPI/template/go/handler"
	"autoAPI/template/go/infrastructure"
	"autoAPI/template/go/mainTemplate"
	"autoAPI/template/go/model"
	"os"
	"path/filepath"
)

type GoGenerator struct{}

func renderDockerfile(table table.Table, dirPath string) error {
	dockerFileContent := dockerfile.Dockerfile(table)
	dockerfileFile, err := os.Create(filepath.Join(dirPath, "Dockerfile"))
	if err != nil {
		return err
	}
	defer dockerfileFile.Close()
	_, err = dockerfileFile.WriteString(dockerFileContent)
	return err
}

func renderGoMod(table table.Table, dirPath string) error {
	modFileContent := goMod.GoMod(table)
	modFile, err := os.Create(filepath.Join(dirPath, "go.mod"))
	if err != nil {
		return err
	}
	defer modFile.Close()
	_, err = modFile.WriteString(modFileContent)
	return err
}

func renderMain(table table.Table, dirPath string) error {
	mainFileContent := mainTemplate.MainTemplate(table)
	mainFile, err := os.Create(filepath.Join(dirPath, "main.go"))
	if err != nil {
		return err
	}
	defer mainFile.Close()
	_, err = mainFile.WriteString(mainFileContent)
	return err
}

func renderHandler(table table.Table, dirPath string) error {
	handlerDir := filepath.Join(dirPath, "handler")
	err := os.Mkdir(handlerDir, 0755)
	if err != nil {
		return err
	}
	handlerFileContent := handler.Handler(table)
	handlerFile, err := os.Create(filepath.Join(handlerDir, "handler.go"))
	if err != nil {
		return err
	}
	defer handlerFile.Close()
	_, err = handlerFile.WriteString(handlerFileContent)
	return err
}

func renderModel(table table.Table, dirPath string) error {
	modelDir := filepath.Join(dirPath, "model")
	err := os.Mkdir(modelDir, 0755)
	if err != nil {
		return err
	}
	modelFileContent := model.Model(table)
	modelFile, err := os.Create(filepath.Join(modelDir, "model.go"))
	if err != nil {
		return err
	}
	defer modelFile.Close()
	_, err = modelFile.WriteString(modelFileContent)
	return err
}

func renderDB(dirPath string) error {
	infrastructureDir := filepath.Join(dirPath, "infrastructure")
	err := os.Mkdir(infrastructureDir, 0755)
	if err != nil {
		return err
	}
	dbFileContent := infrastructure.DB()
	dbFile, err := os.Create(filepath.Join(infrastructureDir, "db.go"))
	if err != nil {
		return err
	}
	defer dbFile.Close()
	_, err = dbFile.WriteString(dbFileContent)
	return err
}

func (generator GoGenerator) Generate(table table.Table, dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	if err = os.Mkdir(dirPath, 0755); err != nil {
		return err
	}
	if err = renderDB(dirPath); err != nil {
		return err
	}
	if err = renderModel(table, dirPath); err != nil {
		return err
	}
	if err = renderHandler(table, dirPath); err != nil {
		return err
	}
	if err = renderMain(table, dirPath); err != nil {
		return err
	}
	if err = renderGoMod(table, dirPath); err != nil {
		return err
	}
	err = renderDockerfile(table, dirPath)
	return err
}

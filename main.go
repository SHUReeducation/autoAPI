//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./template/go/

package main

import (
	"autoAPI/table"
	"autoAPI/template/go/goMod"
	"autoAPI/template/go/handler"
	"autoAPI/template/go/infrastructure"
	"autoAPI/template/go/mainTemplate"
	"autoAPI/template/go/model"
	"autoAPI/withCase"
	"fmt"
	"os"
	"path/filepath"
)

func GenerateAt(table table.Table, dirPath string) error {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return err
	}
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		return err
	}
	infrastructureDir := filepath.Join(dirPath, "infrastructure")
	err = os.Mkdir(infrastructureDir, 0755)
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
	if err != nil {
		return err
	}
	modelDir := filepath.Join(dirPath, "model")
	err = os.Mkdir(modelDir, 0755)
	if err != nil {
		return err
	}
	modelFileContent := model.Model(table)
	modelFile, err := os.Create(filepath.Join(modelDir, "model.go"))
	defer modelFile.Close()
	if err != nil {
		return err
	}
	_, err = modelFile.WriteString(modelFileContent)
	if err != nil {
		return err
	}

	handlerDir := filepath.Join(dirPath, "handler")
	err = os.Mkdir(handlerDir, 0755)
	if err != nil {
		return err
	}
	handlerFileContent := handler.Handler(table)
	handlerFile, err := os.Create(filepath.Join(handlerDir, "handler.go"))
	defer handlerFile.Close()
	if err != nil {
		return err
	}
	_, err = handlerFile.WriteString(handlerFileContent)
	if err != nil {
		return err
	}

	mainFileContent := mainTemplate.MainTemplate(table)
	mainFile, err := os.Create(filepath.Join(dirPath, "main.go"))
	defer mainFile.Close()
	if err != nil {
		return err
	}
	_, err = mainFile.WriteString(mainFileContent)
	if err != nil {
		return err
	}

	modFileContent := goMod.GoMod(table)
	modFile, err := os.Create(filepath.Join(dirPath, "go.mod"))
	defer modFile.Close()
	if err != nil {
		return err
	}
	_, err = modFile.WriteString(modFileContent)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	m := table.Table{
		Name: withCase.New("student"),
		Fields: []table.Field{
			{Name: withCase.New("name"), GoTypeName: "string"},
			{Name: withCase.New("Age"), GoTypeName: "time.Time"},
		},
	}
	err := GenerateAt(m, "/tmp/student")
	fmt.Println(err)
}

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./template/

package main

import (
	"autoAPI/devops"
	"autoAPI/generator"
	"autoAPI/table"
	"autoAPI/yamlParser"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	err := (&cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Put the output code in `PATH`",
			},
		},
		Name:  "autoAPI",
		Usage: "Generate an CRUD api endpoint program automatically!",
		Action: func(c *cli.Context) error {
			f, err := yamlParser.Load(c.String("file"))
			if err != nil {
				return err
			}
			m := table.FromYaml(f)
			metaData := devops.FromYaml(f)
			gen := generator.GoGenerator{}
			devopsGen := generator.DevopsGenerator{}
			if err = gen.Generate(m, c.String("output")); err != nil {
				return err
			}
			err = devopsGen.Generate(metaData, m, c.String("output"))
			return err
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./template/

package main

import (
	"autoAPI/configFile"
	"autoAPI/generator/apiGenerator/golang"
	"autoAPI/generator/cicdGenerator"
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
			f, err := configFile.LoadYaml(c.String("file"))
			if err != nil {
				return err
			}
			gen := golang.APIGenerator{}
			cicdGen := cicdGenerator.CICDGenerator{}
			if err = gen.Generate(f, c.String("output")); err != nil {
				return err
			}
			// todo: See #33
			if f.CICD != nil {
				err = cicdGen.Generate(f, c.String("output"))
			}
			return err
		},
	}).Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

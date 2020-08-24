//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./template/

package main

import (
	"autoAPI/config"
	"autoAPI/generator/apiGenerator/golang"
	"autoAPI/generator/cicdGenerator"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	err := (&cli.App{
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "Load configuration from `FILE`",
			},
			&cli.PathFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Put the output code in `PATH`",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "force",
				Value: false,
				Usage: "If the output PATH dir exists, use '--force' to overwrite it",
			},
			&cli.StringFlag{
				Name:     "dockerusername",
				Aliases:  []string{"du"},
				Usage:    "Your DockerHub's username",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "dockertag",
				Aliases:  []string{"dt"},
				Usage:    "Tag of the generated docker image",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "nodocker",
				Aliases:  []string{"nd"},
				Usage:    "Force not generating docker",
				Required: false,
				Value:    false,
			},
			&cli.StringFlag{
				Name:     "dbengine",
				Aliases:  []string{"dbms"},
				Usage:    "The dbengine you want use",
				Required: false,
			},
		},
		Name:  "autoAPI",
		Usage: "Generate an CRUD api endpoint program automatically!",
		Action: func(c *cli.Context) error {
			var currentConfigure *config.Config
			var err error
			if currentConfigure, err = config.FromCommandLine(c); err != nil {
				return err
			}
			if err = currentConfigure.MergeWithEnv(); err != nil {
				return err
			}
			if configPath := c.String("file"); configPath != "" {
				if err = currentConfigure.MergeWithConfig(c.String("file")); err != nil {
					return err
				}
			}
			if sqlFilePath := c.String("sql"); sqlFilePath != "" {
				if err = currentConfigure.MergeWithSQL(c.String("sql")); err != nil {
					return err
				}
			}
			if err = currentConfigure.MergeWithDB(); err != nil {
				return err
			}
			if err = currentConfigure.MergeWithDefault(); err != nil {
				return err
			}
			err = currentConfigure.Validate()
			if err != nil {
				return err
			}
			if finfo, err := os.Stat(c.String("output")); finfo != nil && finfo.IsDir() {
				if err != nil {
					return err
				}
				if c.IsSet("force") {
					err = os.RemoveAll(c.String("output"))
					if err != nil {
						return err
					}
				} else {
					return errors.New("output dir already exists. Use '--force' to force remove it")
				}
			}
			gen := golang.APIGenerator{}
			cicdGen := cicdGenerator.CICDGenerator{}
			if err = gen.Generate(*currentConfigure, c.String("output")); err != nil {
				return err
			}
			if currentConfigure.CICD != nil {
				err = cicdGen.Generate(*currentConfigure, c.String("output"))
			}
			return err
		},
	}).Run(os.Args)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}
}

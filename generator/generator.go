package generator

import (
	"autoAPI/config"
)

type Generator interface {
	Generate(config config.Config, dirPath string) error
}

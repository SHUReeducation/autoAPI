package generator

import (
	"autoAPI/configFile"
)

type Generator interface {
	Generate(configFile configFile.ConfigFile, dirPath string) error
}

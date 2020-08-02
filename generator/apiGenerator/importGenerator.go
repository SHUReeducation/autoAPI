package apiGenerator

import "autoAPI/configFile"

type Generator interface {
	Generate(configFile configFile.ConfigFile) []string
}

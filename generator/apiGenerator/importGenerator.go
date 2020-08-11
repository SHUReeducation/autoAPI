package apiGenerator

import "autoAPI/config"

type Generator interface {
	Generate(config config.Config) []string
}

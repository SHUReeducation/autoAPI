package apigenerator

import "autoAPI/config"

type Generator interface {
	Generate(config config.Config) []string
}

type TypeTransformer interface {
	Transform(dataBaseType string) string
}

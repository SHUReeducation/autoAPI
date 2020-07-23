package generator

import "autoAPI/table"

type Generator interface {
	Generate(table table.Table, dirPath string) error
}

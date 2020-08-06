package parser

import (
	"autoAPI/configFile/fields/database"
	"autoAPI/configFile/fields/database/field"
	"autoAPI/utility/withCase"
	"errors"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/types"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"regexp"
	"strings"
)

func FillTableInfo(s string, tb *database.Table) error {
	parser := parser.New()
	s = fixSerial(s)
	stmt, err := parser.ParseOneStmt(s, "", "")
	if err != nil {
		return err
	}
	ctStmt, ok := stmt.(*ast.CreateTableStmt)
	if !ok {
		return errors.New("change type to CreateTableStmt fail")
	}
	for _, col := range ctStmt.Cols {
		tb.Fields = append(tb.Fields, field.Field{Name: withCase.New(col.Name.Name.L),
			Type: types.TypeStr(col.Tp.Tp)})
	}
	w := withCase.New(ctStmt.Table.Name.L)
	tb.Name = &w
	return nil
}

func fixSerial(s string) string {
	reg1 := regexp.MustCompile("([\b,\t\n]+)(serial)([\b,\t\n]+)")
	return reg1.ReplaceAllString(strings.ToLower(s), "${1}int${3}")
}

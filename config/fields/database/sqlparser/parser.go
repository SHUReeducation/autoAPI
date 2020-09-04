package sqlparser

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
	"unsafe"

	"autoAPI/config/fields/database/field"
	"autoAPI/utility/withcase"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/types"
	_ "github.com/pingcap/tidb/types/parser_driver" // register
)

// todo: handle multiple dbms
func ParseCreateTable(path string, dbms string) (withcase.WithCase, []field.Field, error) {
	s, err := getSQLString(path)
	if err != nil {
		return withcase.WithCase{}, nil, err
	}
	parser := parser.New()
	stmts, _, err := parser.Parse(s, "", "")
	if err != nil {
		return withcase.WithCase{}, nil, err
	}
	var name withcase.WithCase
	var fields []field.Field
	// todo: handle multiple create statements in one file
	for _, stmt := range stmts {
		ctStmt, ok := stmt.(*ast.CreateTableStmt)
		if !ok {
			return withcase.WithCase{}, nil, errors.New("cast type to CreateTableStmt fail")
		}
		for _, col := range ctStmt.Cols {
			fields = append(fields, field.Field{Name: withcase.New(col.Name.Name.L),
				Type: types.TypeStr(col.Tp.Tp)})
		}
		w := withcase.New(ctStmt.Table.Name.L)
		name = w
	}
	return name, fields, nil
}

func getSQLString(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	s := toString(b)
	s = fixSerial(s)
	s = fixBigSerial(s)
	return s, nil
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func fixSerial(s string) string {
	reg1 := regexp.MustCompile("([,\\s\t\n]+)(serial)([,\\s\t\n]+)")
	return reg1.ReplaceAllString(strings.ToLower(s), "${1}int${3}")
}

func fixBigSerial(s string) string {
	reg1 := regexp.MustCompile("([,\\s\t\n]+)(bigserial)([,\\s\t\n]+)")
	return reg1.ReplaceAllString(strings.ToLower(s), "${1}bigint${3}")
}

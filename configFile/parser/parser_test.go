package parser

import (
	"autoAPI/configFile/fields/database"
	"fmt"
	"testing"
)

func TestFixSerial(t *testing.T) {
	var serialTests = []struct {
		in       string
		expected string
	}{
		{", serial,", ", int,"},
		{" serial,", " int,"},
		{",serialxx", ",serialxx"},
		{"xxxserial", "xxxserial"},
		{" serial ", " int "},
	}
	for _, tt := range serialTests {
		actual := fixSerial(tt.in)
		if actual != tt.expected {
			t.Errorf("fixSerial(%s) = %s; expected %s", tt.in, actual, tt.expected)
		}
	}
}

func TestFixBigSerial(t *testing.T) {
	var serialTests = []struct {
		in       string
		expected string
	}{
		{", bigserial,", ", bigint,"},
		{" bigserial,", " bigint,"},
		{",bigserialxx", ",bigserialxx"},
		{"xxxbigserial", "xxxbigserial"},
		{" bigserial ", " bigint "},
	}
	for _, tt := range serialTests {
		actual := fixBigSerial(tt.in)
		if actual != tt.expected {
			t.Errorf("fixBigSerial(%s) = %s; expected %s", tt.in, actual, tt.expected)
		}
	}
}

func TestFillTableInfo(t *testing.T) {
	var tb database.Table
	err := FillTableInfo("../../example/studentCreateSql.sql", &tb)
	if tb.Name.CamelCase() != "student" {
		t.Errorf("FillTableInfo fail %v", err)
	}
	fmt.Print(tb)
}

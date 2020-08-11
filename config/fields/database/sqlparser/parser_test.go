package sqlparser

import (
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

func TestParseCreateTable(t *testing.T) {
	name, _, err := ParseCreateTable("../../example/studentCreateSql.sql", "pgsql")
	if err != nil {
		t.Errorf("ParseCreateTable fail, should parsed successfully, err=")
		t.Fatal(err)
	}
	if name.CamelCase() != "student" {
		t.Errorf("ParseCreateTable fail, expect name to be \"student\", found %v", name)
		t.FailNow()
	}
}

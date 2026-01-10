package cmd

import (
	"fmt"
	"testing"
)

var phdthesisExpectedFlags = map[string]string{
	"address": "l",
	"author":  "a",
	"key":     "k",
	"month":   "m",
	"note":    "n",
	"school":  "s",
	"title":   "t",
	"type":    "g",
	"year":    "y"}

func initPhDthesisPersistentFlagMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"phdthesis",
			"--braces",
			"--key", "test",
			"--author", "Jane Doe",
			"--school", "King's College London",
			"--year", "2002",
			"--type", "Thesis",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		BracesContains: []string{
			"@phdthesis",
			"{Doe, Jane}",
			"{King's College London}",
			"{2002}",
			"{Thesis}",
			"{Strand, London, WC2R 2LS}",
			"{jul}",
			"{Test Note"},
		CopyArgs: []string{
			"phdthesis",
			"--copy",
			"--key", "test",
			"--author", "Jane Doe",
			"--school", "King's College London",
			"--year", "2002",
			"--type", "Thesis",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@phdthesis",
			"Doe, Jane",
			"King's College London",
			"2002",
			"Thesis",
			"Strand, London, WC2R 2LS",
			"jul",
			"Test Note"},
		SaveArgs: []string{
			"phdthesis",
			"--save", path,
			"--key", "test",
			"--author", "Jane Doe",
			"--school", "King's College London",
			"--year", "2002",
			"--type", "Thesis",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		SaveContains: []string{
			"@phdthesis",
			"Doe, Jane",
			"King's College London",
			"2002",
			"Thesis",
			"Strand, London, WC2R 2LS",
			"jul",
			"Test Note"},
		TempDirPath: path,
	}
	return mock
}

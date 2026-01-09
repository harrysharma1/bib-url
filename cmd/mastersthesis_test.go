package cmd

import (
	"fmt"
	"testing"
)

var mastersthesisExpectedFlags = map[string]string{
	"address": "l",
	"author":  "a",
	"key":     "k",
	"month":   "m",
	"note":    "n",
	"school":  "s",
	"title":   "t",
	"type":    "g",
	"year":    "y"}

func initMastersthesisMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"mastersthesis",
			"--braces",
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--school", "King's College London",
			"--year", "2025",
			"--type", "Dissertation",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		BracesContains: []string{
			"@mastersthesis",
			"{Doe, John}",
			"{Test Title}",
			"{King's College London}",
			"{2025}",
			"{Dissertation}",
			"{Strand, London, WC2R 2LS}",
			"{jul}",
			"{Test Note}"},
		CopyArgs: []string{
			"mastersthesis",
			"--copy",
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--school", "King's College London",
			"--year", "2025",
			"--type", "Dissertation",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@mastersthesis",
			"Doe, John",
			"Test Title",
			"King's College London",
			"2025",
			"Dissertation",
			"Strand, London, WC2R 2LS",
			"jul",
			"Test Note"},
		SaveArgs: []string{
			"mastersthesis",
			"--save", path,
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--school", "King's College London",
			"--year", "2025",
			"--type", "Dissertation",
			"--address", "Strand, London, WC2R 2LS",
			"--month", "jul",
			"--note", "Test Note"},
		SaveContains: []string{
			"@mastersthesis",
			"Doe, John",
			"Test Title",
			"King's College London",
			"2025",
			"Dissertation",
			"Strand, London, WC2R 2LS",
			"jul",
			"Test Note"},
		TempDirPath: path,
	}
	return mock
}

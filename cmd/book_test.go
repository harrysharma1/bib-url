package cmd

import (
	"fmt"
	"testing"
)

var bookExpectedFlags = map[string]string{
	"address":   "l",
	"author":    "a",
	"isbn":      "i",
	"key":       "k",
	"publisher": "p",
	"title":     "t",
	"year":      "y"}

func initBookPersistentFlagMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"book",
			"--braces",
			"--key", "test",
			"--author", "Mary Shelley",
			"--publisher", "Lackington, Hughes, Harding, Mavor, & Jones",
			"--address", "Finsbury Square, London",
			"--year", "1818"},
		BracesContains: []string{
			"@book",
			"{Shelley, Mary}",
			"{Lackington, Hughes, Harding, Mavor, & Jones}",
			"{Finsbury Square, London}",
			"{1818}"},
		CopyArgs: []string{
			"book",
			"--copy",
			"--key", "test",
			"--author", "Mary Shelley",
			"--publisher", "Lackington, Hughes, Harding, Mavor, & Jones",
			"--address", "Finsbury Square, London",
			"--year", "1818"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@book",
			"Shelley, Mary",
			"Lackington, Hughes, Harding, Mavor, & Jones",
			"Finsbury Square, London",
			"1818"},
		SaveArgs: []string{
			"book",
			"--save", path,
			"--key", "test",
			"--author", "Mary Shelley",
			"--publisher", "Lackington, Hughes, Harding, Mavor, & Jones",
			"--address", "Finsbury Square, London",
			"--year", "1818"},
		SaveContains: []string{
			"@book",
			"Shelley, Mary",
			"Lackington, Hughes, Harding, Mavor, & Jones",
			"Finsbury Square, London",
			"1818"},
		TempDirPath: path,
	}

	return mock
}

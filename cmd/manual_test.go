package cmd

import (
	"fmt"
	"testing"
)

var manualExpectedFlags = map[string]string{
	"address":      "l",
	"author":       "a",
	"edition":      "e",
	"key":          "k",
	"month":        "m",
	"note":         "n",
	"organisation": "o",
	"title":        "t",
	"year":         "y"}

func initManualPersistentFlagMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"manual",
			"--braces",
			"--key", "test",
			"--year", "2002",
			"--author", "John Doe",
			"--organisation", "Test Organisation",
			"--address", "Test Avenue",
			"--edition", "2nd",
			"--month", "mar",
			"--note", "Test Note"},
		BracesContains: []string{
			"@manual",
			"{2002}",
			"{Doe, John}",
			"{Test Organisation}",
			"{Test Avenue}",
			"{2nd}",
			"{mar}",
			"{Test Note}"},
		CopyArgs: []string{
			"manual",
			"--copy",
			"--key", "test",
			"--year", "2002",
			"--author", "John Doe",
			"--organisation", "Test Organisation",
			"--address", "Test Avenue",
			"--edition", "2nd",
			"--month", "mar",
			"--note", "Test Note"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@manual",
			"2002",
			"Doe, John",
			"Test Organisation",
			"Test Avenue",
			"2nd",
			"mar",
			"Test Note"},
		SaveArgs: []string{
			"manual",
			"--save", path,
			"--key", "test",
			"--year", "2002",
			"--author", "John Doe",
			"--organisation", "Test Organisation",
			"--address", "Test Avenue",
			"--edition", "2nd",
			"--month", "mar",
			"--note", "Test Note"},
		SaveContains: []string{
			"@manual",
			"2002",
			"Doe, John",
			"Test Organisation",
			"Test Avenue",
			"2nd",
			"mar",
			"Test Note"},
		TempDirPath: path,
	}
	return mock
}

package cmd

import (
	"fmt"
	"testing"
)

var proceedingsExpectedFlags = map[string]string{
	"address":   "l",
	"editor":    "e",
	"key":       "k",
	"month":     "m",
	"number":    "i",
	"publisher": "p",
	"series":    "s",
	"title":     "t",
	"volume":    "v",
	"year":      "y"}

func initProceedingsPersistentFlagMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"proceedings",
			"--braces",
			"--title", "Test Title",
			"--year", "2002",
			"--editor", "Jane Doe",
			"--volume", "1",
			"--number", "12",
			"--series", "Test Series",
			"--address", "Test Avenue",
			"--month", "mar",
			"--publisher", "Test Publisher"},
		BracesContains: []string{
			"@proceedings",
			"{Test Title}",
			"{2002}",
			"{Doe, Jane}",
			"{1}",
			"{12}",
			"{Test Series}",
			"{Test Avenue}",
			"{mar}",
			"{Test Publisher}"},
		CopyArgs: []string{
			"proceedings",
			"--copy",
			"--title", "Test Title",
			"--year", "2002",
			"--editor", "Jane Doe",
			"--volume", "1",
			"--number", "12",
			"--series", "Test Series",
			"--address", "Test Avenue",
			"--month", "mar",
			"--publisher", "Test Publisher"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@proceedings",
			"Test Title",
			"2002",
			"Doe, Jane",
			"1",
			"12",
			"Test Series",
			"Test Avenue",
			"mar",
			"Test Publisher"},
		SaveArgs: []string{
			"proceedings",
			"--save", path,
			"--title", "Test Title",
			"--year", "2002",
			"--editor", "Jane Doe",
			"--volume", "1",
			"--number", "12",
			"--series", "Test Series",
			"--address", "Test Avenue",
			"--month", "mar",
			"--publisher", "Test Publisher"},
		SaveContains: []string{
			"@proceedings",
			"Test Title",
			"2002",
			"Doe, Jane",
			"1",
			"12",
			"Test Series",
			"Test Avenue",
			"mar",
			"Test Publisher"},
		TempDirPath: path,
	}
	return mock
}

package cmd

import (
	"fmt"
	"testing"
)

var techreportExpectedFlags = map[string]string{
	"address":     "l",
	"author":      "a",
	"institution": "i",
	"key":         "k",
	"month":       "m",
	"number":      "n",
	"title":       "t",
	"year":        "y"}

func initTechReportMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"techreport",
			"--braces",
			"--key", "test",
			"--title", "Test Title",
			"--author", "Jane Doe",
			"--institution", "Test Institution",
			"--address", "Test Avenue",
			"--number", "2",
			"--year", "2002",
			"--month", "sep"},
		BracesContains: []string{
			"@techreport",
			"{Test Title}",
			"{Doe, Jane}",
			"{Test Institution}",
			"{Test Avenue}",
			"{2}",
			"{2002}",
			"{sep}"},
		CopyArgs: []string{
			"techreport",
			"--copy",
			"--key", "test",
			"--title", "Test Title",
			"--author", "Jane Doe",
			"--institution", "Test Institution",
			"--address", "Test Avenue",
			"--number", "2",
			"--year", "2002",
			"--month", "sep"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@techreport",
			"Test Title",
			"Doe, Jane",
			"Test Institution",
			"Test Avenue",
			"2",
			"2002",
			"sep"},
		SaveArgs: []string{
			"techreport",
			"--save", path,
			"--key", "test",
			"--title", "Test Title",
			"--author", "Jane Doe",
			"--institution", "Test Institution",
			"--address", "Test Avenue",
			"--number", "2",
			"--year", "2002",
			"--month", "sep"},
		SaveContains: []string{
			"@techreport",
			"Test Title",
			"Doe, Jane",
			"Test Institution",
			"Test Avenue",
			"2",
			"2002",
			"sep"},
		TempDirPath: path,
	}
	return mock
}

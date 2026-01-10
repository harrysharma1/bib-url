package cmd

import (
	"fmt"
	"testing"
)

var miscExpectedFlags = map[string]string{
	"author":       "a",
	"howpublished": "p",
	"key":          "k",
	"note":         "n",
	"title":        "t",
	"year":         "y"}

func initMiscPersistentFlagMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"misc",
			"--braces",
			"--key", "test",
			"--title", "Test Title",
			"--author", "John Doe",
			"--howpublished", "https://example.com",
			"--year", "2002",
			"--note", "Test Note"},
		BracesContains: []string{
			"@misc",
			"{Test Title}",
			"{Doe, John}",
			"{https://example.com}",
			"{2002}",
			"{Test Note}"},
		CopyArgs: []string{
			"misc",
			"--copy",
			"--key", "test",
			"--title", "Test Title",
			"--author", "John Doe",
			"--howpublished", "https://example.com",
			"--year", "2002",
			"--note", "Test Note"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@misc",
			"Test Title",
			"Doe, John",
			"https://example.com",
			"2002",
			"Test Note"},
		SaveArgs: []string{
			"misc",
			"--save", path,
			"--key", "test",
			"--title", "Test Title",
			"--author", "John Doe",
			"--howpublished", "https://example.com",
			"--year", "2002",
			"--note", "Test Note"},
		SaveContains: []string{
			"@misc",
			"Test Title",
			"Doe, John",
			"https://example.com",
			"2002",
			"Test Note"},
		TempDirPath: path,
	}
	return mock

}

package cmd

import (
	"fmt"
	"testing"
)

var unpublishedExpectedFlags = map[string]string{
	"author":      "a",
	"institution": "i",
	"key":         "k",
	"title":       "t",
	"year":        "y"}

func initUnpublishedMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"unpublished",
			"--braces",
			"--key", "test",
			"--author", "Monkey D. Garp",
			"--title", "Hero of The Marines",
			"--institution", "World Govt.",
			"--year", "1446"},
		BracesContains: []string{
			"@unpublished",
			"{Garp, Monkey D.}",
			"{Hero of The Marines}",
			"{World Govt.}",
			"{1446}"},
		CopyArgs: []string{
			"unpublished",
			"--copy",
			"--key", "test",
			"--author", "Monkey D. Garp",
			"--title", "Hero of The Marines",
			"--institution", "World Govt.",
			"--year", "1446"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@unpublished",
			"Garp, Monkey D.",
			"Hero of The Marines",
			"World Govt.",
			"1446"},
		SaveArgs: []string{
			"unpublished",
			"--save", path,
			"--key", "test",
			"--author", "Monkey D. Garp",
			"--title", "Hero of The Marines",
			"--institution", "World Govt.",
			"--year", "1446"},
		SaveContains: []string{
			"@unpublished",
			"Garp, Monkey D.",
			"Hero of The Marines",
			"World Govt.",
			"1446"},
		TempDirPath: path,
	}
	return mock
}

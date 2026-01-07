package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.design/x/clipboard"
)

var articleExpectedFlags = map[string]string{
	"author":  "a",
	"doi":     "d",
	"journal": "j",
	"key":     "k",
	"month":   "m",
	"note":    "n",
	"number":  "i",
	"pages":   "p",
	"title":   "t",
	"volume":  "v",
	"year":    "y",
}

func initArticleMock(t *testing.T) PersistentFlagsMock {
	t.Helper()
	path := fmt.Sprintf("%s/reference.bib", t.TempDir())
	mock := PersistentFlagsMock{
		BracesArgs: []string{
			"article",
			"--braces",
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--journal", "Test Journal",
			"--year", "2002"},
		BracesContains: []string{
			"@article",
			"{Doe, John}",
			"{Test Title}",
			"{Test Journal}",
			"{2002}"},
		CopyArgs: []string{
			"article",
			"--copy",
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--journal", "Test Journal",
			"--year", "2002"},
		CopyContains: []string{
			"copied bibtex entry to cliplboard!!!",
			"@article",
			"Doe, John",
			"Test Title",
			"Test Journal",
			"2002"},
		SaveArgs: []string{
			"article",
			"--save", path,
			"--key", "test",
			"--author", "John Doe",
			"--title", "Test Title",
			"--journal", "Test Journal",
			"--year", "2002"},
		SaveContains: []string{
			"@article",
			"Doe, John",
			"Test Title",
			"Test Journal",
			"2002"},
		TempDirPath: path,
	}
	return mock
}

func TestArticlePersistentFlagsBraces(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs([]string{
		"article",
		"--braces",
		"--key", "test",
		"--author", "John Doe",
		"--title", "Test Title",
		"--journal", "Test Journal",
		"--year", "2002"})
	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err, "should not err check command flags in test")
	})
	assert.Contains(t, out, "@article")
	assert.Contains(t, out, "{Doe, John}", "braces should exist")
	assert.Contains(t, out, "{Test Title}", "braces should exist")
	assert.Contains(t, out, "{Test Journal}", "braces should exist")
	assert.Contains(t, out, "{2002}", "braces should exist")
	resetGlobalState(t)
}

func TestArticlePersistentFlagsCopy(t *testing.T) {
	resetGlobalState(t)
	cmd := rootCmd
	cmd.SetArgs([]string{
		"article",
		"--copy",
		"--key", "test",
		"--author", "John Doe",
		"--title", "Test Title",
		"--journal", "Test Journal",
		"--year", "2002"})
	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err, "should not err check command flags in test")
	})
	assert.Contains(t, out, "copied bibtex entry to cliplboard!!!")
	assert.Contains(t, out, "@article")
	assert.Contains(t, out, "Doe, John")
	assert.Contains(t, out, "Test Title")
	assert.Contains(t, out, "Test Journal")
	assert.Contains(t, out, "2002")
	hasClipboard := clipboard.Read(clipboard.FmtText)
	assert.Contains(t, out, string(hasClipboard), "clipboard should copy bibtex only not stdout")
	resetGlobalState(t)
}
func TestArticlePersistentFlagsSave(t *testing.T) {
	resetGlobalState(t)
	tmpDir := t.TempDir()
	path := fmt.Sprintf("%s/reference.bib", tmpDir)

	cmd := rootCmd
	cmd.SetArgs([]string{
		"article",
		"--save", path,
		"--key", "test",
		"--author", "John Doe",
		"--title", "Test Title",
		"--journal", "Test Journal",
		"--year", "2002"})
	out := captureStdout(t, func() {
		err := cmd.Execute()
		assert.Nil(t, err, "should not err check command flags in test")
	})
	assert.Contains(t, out, "@article")
	assert.Contains(t, out, "Doe, John")
	assert.Contains(t, out, "Test Title")
	assert.Contains(t, out, "Test Journal")
	assert.Contains(t, out, "2002")
	assertSaveFile(t, path, out)
	resetGlobalState(t)
}

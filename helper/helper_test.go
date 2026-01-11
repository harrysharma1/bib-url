package helper

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.design/x/clipboard"
)

// Test bibtex is copied into system clipboard
func TestCopy(t *testing.T) {
	wantString := "Copied String"
	Copy(wantString)
	haveString := string(clipboard.Read(clipboard.FmtText))
	assert.Equal(t, wantString, haveString, "text not copied to os clipboard")
}

// Test to make sure function returns error on invalid file type
func TestSaveInvalidFileType(t *testing.T) {
	saveInvalidFP := "reference.go"
	haveInvalidErr := Save(saveInvalidFP, "Test")
	assert.EqualError(t, haveInvalidErr, "file type not .bib did not write entry")
}

// Test to make sure function returns nil on valid file type
func TestSaveValidFileType(t *testing.T) {
	tmpDir := t.TempDir()
	saveValidFP := filepath.Join(tmpDir, "reference.bib")
	haveValidErr := Save(saveValidFP, "Test")
	assert.Nil(t, haveValidErr, "should save for valid filepath")
}

// Test to make sure function returns a specific error if golang fails to create/apppend to valid file type
func TestSaveValidFileTypeFailedOpeningFile(t *testing.T) {
	saveValidInvalidFP := "/reference.bib"
	haveValidOpenFileErr := Save(saveValidInvalidFP, "Test")
	assert.EqualError(t, haveValidOpenFileErr, "error opening file", "OPENING FILE")

}

// Struct for a purposfully failing error writer
type errorWriter struct{}

// The mock function of error writer object to test for specific fail
func (e errorWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("forced writing error")
}

// Test to make sure function returns a specific error if golang fails to write to valid file type
func TestSaveValidFileTypeFailedWritingToFile(t *testing.T) {
	err := SaveToWriter(errorWriter{}, "test")
	assert.NotNil(t, err, "should not be nil")
	assert.EqualError(t, err, "error writing to file", "this specific error should occur on failing write test")
}

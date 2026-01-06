package helper

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.design/x/clipboard"
)

func TestCopy(t *testing.T) {
	wantString := "Copied String"
	Copy(wantString)
	haveString := string(clipboard.Read(clipboard.FmtText))
	assert.Equal(t, wantString, haveString, "text not copied to os clipboard")
}

func TestSaveInvalidFileType(t *testing.T) {
	saveInvalidFP := "reference.go"
	haveInvalidErr := Save(saveInvalidFP, "Test")
	assert.EqualError(t, haveInvalidErr, "file type not .bib did not write entry")
}

func TestSaveValidFileType(t *testing.T) {
	tmpDir := t.TempDir()
	saveValidFP := filepath.Join(tmpDir, "reference.bib")
	haveValidErr := Save(saveValidFP, "Test")
	assert.Nil(t, haveValidErr, "should save for valid filepath")
}

func TestSaveValidFileTypeFailedOpeningFile(t *testing.T) {
	saveValidInvalidFP := "/reference.bib"
	haveValidOpenFileErr := Save(saveValidInvalidFP, "Test")
	assert.EqualError(t, haveValidOpenFileErr, "error opening file", "OPENING FILE")

}

type errorWriter struct{}

func (e errorWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("forced writing error")
}

func TestSaveValidFileTypeFailedWritingToFile(t *testing.T) {
	err := SaveToWriter(errorWriter{}, "test")
	assert.NotNil(t, err, "should not be ni")
	assert.EqualError(t, err, "error writing to file", "this specific error should occur on failing write test")
}

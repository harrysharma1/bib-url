package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

// Copy bibtex entry into OS clipboard
func Copy(bibtex string) {
	fmt.Println("copied bibtex entry to cliplboard!!!")
	clipboard.Write(clipboard.FmtText, []byte(bibtex))
}

// Save bibtex entry into .bib file (either by creating or appending to)
//
// Here the *os.File gets passed to SaveToWriter() function with set permissions
func Save(save string, bibtex string) error {
	if filepath.Ext(save) != ".bib" {
		return fmt.Errorf("file type not .bib did not write entry")
	}

	f, err := os.OpenFile(save, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file")
	}
	defer f.Close()

	return SaveToWriter(f, bibtex)
}

// Save bibtex string by using passed io.Writer object from Save() function
//
// This writes it as chunked bytes rather than using .WriteString() function, just because it is more robust
func SaveToWriter(w io.Writer, bibtex string) error {
	bibtex += "\n"
	_, err := w.Write([]byte(bibtex))
	if err != nil {
		return fmt.Errorf("error writing to file")
	}
	return nil
}

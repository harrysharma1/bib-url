package helper

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

func Copy(bibtex string) {
	fmt.Println("copied bibtex entry to cliplboard!!!")
	clipboard.Write(clipboard.FmtText, []byte(bibtex))
}

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

func SaveToWriter(w io.Writer, bibtex string) error {
	bibtex += "\n"
	_, err := w.Write([]byte(bibtex))
	if err != nil {
		return fmt.Errorf("error writing to file")
	}
	return nil
}

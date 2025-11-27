package helper

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.design/x/clipboard"
)

var (
	braces_open  = "{"
	braces_close = "}"
	speechmarks  = `"`
)

func Copy(bibtex string) {
	fmt.Println("copied bibtex entry to cliplboard!!!")
	clipboard.Write(clipboard.FmtText, []byte(bibtex))
}

func Save(save string, bibtex string) error {
	if filepath.Ext(save) == ".bib" {
		f, err := os.OpenFile(save, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("error opening file")
		}
		defer f.Close()
		bibtex += "\n"
		_, err = f.WriteString(bibtex)
		if err != nil {
			return fmt.Errorf("error writing to file")
		}
		fmt.Printf("saved bibtex entry to `%s`!!!\n", save)
		fmt.Println(bibtex[:len(bibtex)-1])
		return nil
	} else {
		return fmt.Errorf("error file type not .bib did not write entry")
	}

}

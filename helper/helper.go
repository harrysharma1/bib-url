package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.design/x/clipboard"
)

type Field struct {
	key string
	val string
}

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

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

func defaultIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func wrap(s string, braces bool) string {
	if braces {
		return "{" + s + "}"
	}
	if slices.Contains(months, s) {
		return s
	}
	return `"` + s + `"`

}

func getMaxKeyLength(fields []Field) int {
	maxKeyLength := 0

	for _, f := range fields {
		if len(f.key) > maxKeyLength {
			maxKeyLength = len(f.key)
		}
	}

	return maxKeyLength
}

func formatFields(fields []Field, braces bool) []string {
	formattedFields := make([]string, len(fields))
	maxKeyLength := getMaxKeyLength(fields)
	for i, f := range fields {
		padding := maxKeyLength - len(f.key)
		switch f.key {
		case "year":
			if _, err := strconv.Atoi(f.val); err == nil {
				if !braces {
					formattedFields[i] = fmt.Sprintf("\t%s%s = %s", f.key, strings.Repeat(" ", padding), f.val)
				} else {
					formattedFields[i] = fmt.Sprintf("\t%s%s = %s", f.key, strings.Repeat(" ", padding), wrap(f.val, braces))
				}
			} else {
				formattedFields[i] = fmt.Sprintf("\t%s%s = %s", f.key, strings.Repeat(" ", padding), wrap(f.val, braces))
			}
		default:
			formattedFields[i] = fmt.Sprintf("\t%s%s = %s", f.key, strings.Repeat(" ", padding), wrap(f.val, braces))
		}
	}
	return formattedFields

}

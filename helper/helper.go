package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.design/x/clipboard"
)

var (
	braces_open  = "{"
	braces_close = "}"
	speechmarks  = `"`
	months       = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
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

func defaultIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func parseFields(fields []string) map[string][]string {
	ret_val := make(map[string][]string)

	for _, val := range fields {
		parts := strings.SplitN(val, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(strings.ToLower(parts[0]))
		val := strings.TrimSpace(parts[1])
		ret_val[key] = append(ret_val[key], val)
	}
	return ret_val
}

func largestFieldInt(s []string) int {
	ret_val := 0
	for _, v := range s {
		if len(v) > ret_val {
			ret_val = len(v)
		}
	}
	return ret_val
}

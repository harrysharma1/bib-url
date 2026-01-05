package format

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Field struct {
	key string
	val string
}

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

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
func toBibtexName(name string) string {
	if strings.Contains(name, ",") {
		return name
	}
	parts := strings.Fields(name)
	if len(parts) < 2 {
		return name
	}
	first := parts[:len(parts)-1]
	last := parts[len(parts)-1]

	return last + ", " + strings.Join(first, " ")

}
func formatAuthors(authors []string) string {
	formattedAuthors := []string{}
	for _, author := range authors {
		formattedAuthors = append(formattedAuthors, toBibtexName(author))
	}
	return strings.Join(formattedAuthors, " and ")
}

func formatEditors(editors []string) string {
	formattedEditor := []string{}
	for _, editor := range editors {
		formattedEditor = append(formattedEditor, toBibtexName(editor))
	}
	return strings.Join(formattedEditor, " and ")
}

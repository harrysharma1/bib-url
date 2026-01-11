package format

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// Fields for bibtex formatting/generation
type Field struct {
	key string
	val string
}

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

// Returns a default value if an empty string is passed to this function
func defaultIfEmpty(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

// Returns a string wrapped in {} or "" or nothing based on braces
func wrap(s string, braces bool) string {
	if braces {
		return "{" + s + "}"
	}
	if slices.Contains(months, s) {
		return s
	}
	return `"` + s + `"`

}

// Returns length of largest key in a list of fields
func getMaxKeyLength(fields []Field) int {
	maxKeyLength := 0

	for _, f := range fields {
		if len(f.key) > maxKeyLength {
			maxKeyLength = len(f.key)
		}
	}

	return maxKeyLength
}

// Returns a list of formatted strings based on a list of Fields and braces
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

// Return formatted author string
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

// Return a single string delimited by `and` with formatted author names
func formatAuthors(authors []string) string {
	formattedAuthors := []string{}
	for _, author := range authors {
		formattedAuthors = append(formattedAuthors, toBibtexName(author))
	}
	return strings.Join(formattedAuthors, " and ")
}

// Return a single string delimited by `and` with formatted editor names
func formatEditors(editors []string) string {
	formattedEditor := []string{}
	for _, editor := range editors {
		formattedEditor = append(formattedEditor, toBibtexName(editor))
	}
	return strings.Join(formattedEditor, " and ")
}

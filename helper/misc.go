package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatMiscBibtex(
	miscCiteKey string,
	miscTitle string,
	miscAuthors []string,
	miscHowPublished string,
	miscYear string,
	miscNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@misc{")

	if miscCiteKey != "" {
		sb.WriteString(miscCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")
	fields := []string{}
	wrap := func(s string) string {
		if braces {
			return "{" + s + "}"
		}

		if _, err := strconv.Atoi(s); err == nil {
			return s
		}

		return `"` + s + `"`
	}

	// REQUIRED
	fields = append(fields, fmt.Sprintf("\ttitle        = %s", wrap(defaultIfEmpty(miscTitle, "<Title>"))))
	if len(miscAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap(strings.Join(miscAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap("<Lastname, Firstname>")))
	}

	fields = append(fields, fmt.Sprintf("\thowpublished = %s", wrap(defaultIfEmpty(miscHowPublished, "<How Published>"))))
	fields = append(fields, fmt.Sprintf("\tyear         = %s", wrap(defaultIfEmpty(miscYear, "<2002>"))))
	fields = append(fields, fmt.Sprintf("\tnote         = %s", wrap(defaultIfEmpty(miscNote, "<Note>"))))

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

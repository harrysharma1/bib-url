package format

import (
	"strings"

	"github.com/google/uuid"
)

// Formats passed values as an @misc bibtex response
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
	fields := []Field{}

	// REQUIRED
	fields = append(fields, Field{"title", defaultIfEmpty(miscTitle, "<Title>")})
	if len(miscAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(miscAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}

	fields = append(fields, Field{"howpublished", defaultIfEmpty(miscHowPublished, "<How Published>")})
	fields = append(fields, Field{"year", defaultIfEmpty(miscYear, "<2002>")})
	fields = append(fields, Field{"note", defaultIfEmpty(miscNote, "<Note>")})

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

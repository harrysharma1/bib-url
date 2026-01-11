package format

import (
	"strings"

	"github.com/google/uuid"
)

// Formats passed values as an @mastersthesis bibtex response
func FormatMastersThesisBibtex(
	mastersthesisCiteKey string,
	mastersthesisAuthors []string,
	mastersthesisTitle string,
	mastersthesisSchool string,
	mastersthesisYear string,
	mastersthesisType string,
	mastersthesisAddress string,
	mastersthesisMonth string,
	mastersthesisNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@mastersthesis{")

	if mastersthesisCiteKey != "" {
		sb.WriteString(mastersthesisCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")

	fields := []Field{}

	// REQUIRED
	if len(mastersthesisAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(mastersthesisAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(mastersthesisTitle, "<Title>")})
	fields = append(fields, Field{"school", defaultIfEmpty(mastersthesisSchool, "<School>")})
	fields = append(fields, Field{"year", defaultIfEmpty(mastersthesisYear, "<2002>")})
	// OPTIONAL
	if mastersthesisType != "" {
		fields = append(fields, Field{"type", mastersthesisType})
	}
	if mastersthesisAddress != "" {
		fields = append(fields, Field{"address", mastersthesisAddress})
	}
	if mastersthesisMonth != "" {
		fields = append(fields, Field{"month", mastersthesisMonth})
	}
	if mastersthesisNote != "" {
		fields = append(fields, Field{"note", mastersthesisNote})
	}

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

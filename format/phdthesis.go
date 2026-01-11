package format

import (
	"strings"

	"github.com/google/uuid"
)

// Formats passed values as an @phdthesis bibtex response
func FormatPhDThesisBibtex(
	phdthesisCiteKey string,
	phdthesisAuthors []string,
	phdthesisTitle string,
	phdthesisSchool string,
	phdthesisYear string,
	phdthesisType string,
	phdthesisAddress string,
	phdthesisMonth string,
	phdthesisNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@phdthesis{")

	if phdthesisCiteKey != "" {
		sb.WriteString(phdthesisCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")
	fields := []Field{}

	// REQUIRED
	if len(phdthesisAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(phdthesisAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})

	}
	fields = append(fields, Field{"title", defaultIfEmpty(phdthesisTitle, "<Title>")})
	fields = append(fields, Field{"school", defaultIfEmpty(phdthesisSchool, "<School>")})
	fields = append(fields, Field{"year", defaultIfEmpty(phdthesisYear, "<2002>")})

	// OPTIONAL
	if phdthesisType != "" {
		fields = append(fields, Field{"type", phdthesisType})
	}
	if phdthesisAddress != "" {
		fields = append(fields, Field{"address", phdthesisAddress})
	}
	if phdthesisMonth != "" {
		fields = append(fields, Field{"month", phdthesisMonth})
	}
	if phdthesisNote != "" {
		fields = append(fields, Field{"note", phdthesisNote})
	}
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

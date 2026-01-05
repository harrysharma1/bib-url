package format

import (
	"strings"

	"github.com/google/uuid"
)

func FormatUnpublishedBibtex(
	unpublishedCiteKey string,
	unpublishedAuthors []string,
	unpublishedtTitle string,
	unpublishedInstitution string,
	unpublishedYear string,
	braces bool) string {

	var sb strings.Builder
	sb.WriteString("@unpublished{")
	if unpublishedCiteKey != "" {
		sb.WriteString(unpublishedCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")
	fields := []Field{}

	// REQUIRED
	if len(unpublishedAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(unpublishedAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(unpublishedtTitle, "<Title>")})
	fields = append(fields, Field{"institution", defaultIfEmpty(unpublishedInstitution, "<Institution>")})
	fields = append(fields, Field{"year", defaultIfEmpty(unpublishedYear, "<2002>")})
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

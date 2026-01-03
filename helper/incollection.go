package helper

import (
	"strings"

	"github.com/google/uuid"
)

func FormatIncollectionBibtex(
	incollectionCiteKey string,
	incollectionAuthors []string,
	incollectionTitle string,
	incollectionBookTitle string,
	incollectionPublisher string,
	incollectionYear string,
	incollectionEditors []string,
	incollectionVolume string,
	incollectionNumber string,
	incollectionSeries string,
	incollectionPages string,
	incollectionAddress string,
	incollectionMonth string,
	braces bool) string {

	var sb strings.Builder
	sb.WriteString("@incollection{")

	if incollectionCiteKey != "" {
		sb.WriteString(incollectionCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")

	fields := []Field{}

	// REQUIRED
	if len(incollectionAuthors) > 0 {
		fields = append(fields, Field{"author", strings.Join(incollectionAuthors, " and ")})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(incollectionTitle, "<Title>")})
	fields = append(fields, Field{"booktitle", defaultIfEmpty(incollectionBookTitle, "<Book Title>")})
	fields = append(fields, Field{"publisher", defaultIfEmpty(incollectionPublisher, "<Publisher>")})
	fields = append(fields, Field{"year", defaultIfEmpty(incollectionYear, "<2002>")})

	// OPTIONAL
	if len(incollectionEditors) > 0 {
		fields = append(fields, Field{"editor", strings.Join(incollectionEditors, " and ")})
	}
	if incollectionVolume != "" {
		fields = append(fields, Field{"volume", incollectionVolume})
	}
	if incollectionNumber != "" {
		fields = append(fields, Field{"number", incollectionNumber})
	}
	if incollectionSeries != "" {
		fields = append(fields, Field{"series", incollectionSeries})
	}
	if incollectionPages != "" {
		fields = append(fields, Field{"pages", incollectionPages})
	}
	if incollectionAddress != "" {
		fields = append(fields, Field{"address", incollectionAddress})
	}
	if incollectionMonth != "" {
		fields = append(fields, Field{"month", incollectionMonth})
	}

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

package format

import (
	"strings"

	"github.com/google/uuid"
)

// Formats passed values as on @inbook bibtex response
func FormatInbookBibtex(
	inbookCiteKey string,
	inbookAuthors []string,
	inbookTitle string,
	inbookBookTitle string,
	inbookPublisher string,
	inbookYear string,
	inbookEditors []string,
	inbookVolume string,
	inbookNumber string,
	inbookSeries string,
	inbookAddress string,
	inbookEdition string,
	inboookMonth string,
	inbookPages string,
	inbookNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@inbook{")

	if inbookCiteKey != "" {
		sb.WriteString(inbookCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")

	fields := []Field{}

	// REQUIRED
	if len(inbookAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(inbookAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(inbookTitle, "<Title>")})
	fields = append(fields, Field{"booktitle", defaultIfEmpty(inbookBookTitle, "<Book Title>")})
	fields = append(fields, Field{"publisher", defaultIfEmpty(inbookPublisher, "<Publisher>")})
	fields = append(fields, Field{"year", defaultIfEmpty(inbookYear, "<2002>")})

	// OPTIONAL
	if len(inbookEditors) > 0 {
		fields = append(fields, Field{"editor", formatEditors(inbookEditors)})
	}
	if inbookVolume != "" {
		fields = append(fields, Field{"volume", inbookVolume})
	}
	if inbookNumber != "" {
		fields = append(fields, Field{"number", inbookNumber})
	}
	if inbookSeries != "" {
		fields = append(fields, Field{"series", inbookSeries})
	}
	if inbookAddress != "" {
		fields = append(fields, Field{"address", inbookAddress})
	}
	if inbookEdition != "" {
		fields = append(fields, Field{"edition", inbookEdition})
	}
	if inboookMonth != "" {
		fields = append(fields, Field{"month", inboookMonth})
	}
	if inbookPages != "" {
		fields = append(fields, Field{"pages", inbookPages})
	}
	if inbookNote != "" {
		fields = append(fields, Field{"note", inbookNote})
	}
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

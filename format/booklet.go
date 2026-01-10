package format

import (
	"strings"

	"github.com/google/uuid"
)

// Formats passed values as an @booklet bibtex respsonse
func FormatBookletBibtex(
	bookletCiteKey string,
	bookletTitle string,
	bookletAuthors []string,
	bookletHowPublished string,
	bookletAddress string,
	bookletYear string,
	bookletEditors []string,
	bookletVolume string,
	bookletNumber string,
	bookletSeries string,
	bookletOrganisation string,
	bookletMonth string,
	bookletNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@booklet{")

	if bookletCiteKey != "" {
		sb.WriteString(bookletCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")
	fields := []Field{}

	// REQUIRED
	fields = append(fields, Field{"title", defaultIfEmpty(bookletTitle, "<Title>")})
	if len(bookletAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(bookletAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"howpublished", defaultIfEmpty(bookletHowPublished, "<How Published>")})
	fields = append(fields, Field{"address", defaultIfEmpty(bookletAddress, "<Address>")})
	// DO NOT WRAP
	fields = append(fields, Field{"year", defaultIfEmpty(bookletYear, "<2002>")})

	// OPTIONAL
	if len(bookletEditors) > 0 {
		fields = append(fields, Field{"editor", formatEditors(bookletEditors)})
	}
	if bookletVolume != "" {
		fields = append(fields, Field{"volume", bookletVolume})
	}
	if bookletNumber != "" {
		fields = append(fields, Field{"number", bookletNumber})
	}
	if bookletSeries != "" {
		fields = append(fields, Field{"series", bookletSeries})
	}
	if bookletOrganisation != "" {
		fields = append(fields, Field{"organisation", bookletOrganisation})
	}
	if bookletMonth != "" {
		fields = append(fields, Field{"month", bookletMonth})
	}
	if bookletNote != "" {
		fields = append(fields, Field{"note", bookletNote})
	}

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")

	return sb.String()

}

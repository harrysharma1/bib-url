package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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

	fields := []string{}
	wrap := func(field string, s string) string {
		if braces {
			return "{" + s + "}"
		}
		if slices.Contains(months, s) {
			return s
		}
		if field == "year" {
			if _, err := strconv.Atoi(s); err == nil {
				return s
			}
		}
		return `"` + s + `"`

	}

	// REQUIRED
	if len(inbookAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap("author", strings.Join(inbookAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle     = %s", wrap("title", defaultIfEmpty(inbookTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tbooktitle = %s", wrap("booktitle", defaultIfEmpty(inbookBookTitle, "<Book Title>"))))
	fields = append(fields, fmt.Sprintf("\tpublisher = %s", wrap("publisher", defaultIfEmpty(inbookPublisher, "<Publisher>"))))
	fields = append(fields, fmt.Sprintf("\tyear      = %s", wrap("year", defaultIfEmpty(inbookYear, "<2002>"))))

	// OPTIONAL
	if len(inbookEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor    = %s", wrap("editor", strings.Join(inbookEditors, " and "))))
	}
	if inbookVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume    = %s", wrap("volume", inbookVolume)))
	}
	if inbookNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber    = %s", wrap("number", inbookNumber)))
	}
	if inbookSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries    = %s", wrap("series", inbookSeries)))
	}
	if inbookAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress   = %s", wrap("address", inbookAddress)))
	}
	if inbookEdition != "" {
		fields = append(fields, fmt.Sprintf("\tedition   = %s", wrap("edition", inbookEdition)))
	}
	if inboookMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth     = %s", wrap("month", inboookMonth)))
	}
	if inbookPages != "" {
		fields = append(fields, fmt.Sprintf("\tpages     = %s", wrap("pages", inbookPages)))
	}
	if inbookNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote      = %s", wrap("note", inbookNote)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

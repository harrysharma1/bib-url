package helper

import (
	"fmt"
	"slices"
	"strconv"
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
	if len(incollectionAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap("author", strings.Join(incollectionAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle     = %s", wrap("title", defaultIfEmpty(incollectionTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tbooktitle = %s", wrap("booktitle", defaultIfEmpty(incollectionBookTitle, "<Book Title>"))))
	fields = append(fields, fmt.Sprintf("\tpublisher = %s", wrap("publisher", defaultIfEmpty(incollectionPublisher, "<Publisher>"))))
	fields = append(fields, fmt.Sprintf("\tyear      = %s", wrap("year", defaultIfEmpty(incollectionYear, "<2002>"))))

	// OPTIONAL
	if len(incollectionEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor    = %s", wrap("editor", strings.Join(incollectionEditors, " and "))))
	}
	if incollectionVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume    = %s", wrap("volume", incollectionVolume)))
	}
	if incollectionNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber    = %s", wrap("number", incollectionNumber)))
	}
	if incollectionSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries    = %s", wrap("series", incollectionSeries)))
	}
	if incollectionPages != "" {
		fields = append(fields, fmt.Sprintf("\tpages     = %s", wrap("pages", incollectionPages)))
	}
	if incollectionAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress   = %s", wrap("address", incollectionAddress)))
	}
	if incollectionMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth     = %s", wrap("month", incollectionMonth)))
	}

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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
	fields = append(fields, fmt.Sprintf("\ttitle        = %s", wrap("title", defaultIfEmpty(bookletTitle, "<Title>"))))
	if len(bookletAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap("author,", strings.Join(bookletAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\thowpublished = %s", wrap("howpublished", defaultIfEmpty(bookletHowPublished, "<How Published>"))))
	fields = append(fields, fmt.Sprintf("\taddress      = %s", wrap("address", defaultIfEmpty(bookletAddress, "<Address>"))))
	fields = append(fields, fmt.Sprintf("\tyear         = %s", wrap("year", defaultIfEmpty(bookletYear, "2002"))))

	// OPTIONAL
	if len(bookletEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor       = %s", wrap("editor", strings.Join(bookletEditors, " and "))))
	}
	if bookletVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume       = %s", wrap("volume", bookletVolume)))
	}
	if bookletNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber       = %s", wrap("number", bookletNumber)))
	}
	if bookletSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries       = %s", wrap("series", bookletSeries)))
	}
	if bookletOrganisation != "" {
		fields = append(fields, fmt.Sprintf("\torganisation = %s", wrap("organisation", bookletOrganisation)))
	}
	if bookletMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth        = %s", wrap("month", bookletMonth)))
	}
	if bookletNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote         = %s", wrap("note", bookletNote)))
	}

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")

	return sb.String()

}

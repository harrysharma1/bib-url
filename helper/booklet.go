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

	wrap := func(s string) string {
		if braces {
			return "{" + s + "}"
		}
		if slices.Contains(months, s) {
			return s
		}
		if _, err := strconv.Atoi(s); err == nil {
			return s
		}
		return `"` + s + `"`

	}
	// REQUIRED
	fields = append(fields, fmt.Sprintf("\ttitle = %s", wrap(defaultIfEmpty(bookletTitle, "<Title>"))))
	if len(bookletAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor = %s", wrap(strings.Join(bookletAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor = %s", wrap("<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\thowpublished = %s", wrap(defaultIfEmpty(bookletHowPublished, "<How Published>"))))
	fields = append(fields, fmt.Sprintf("\taddress = %s", wrap(defaultIfEmpty(bookletAddress, "<Address>"))))
	fields = append(fields, fmt.Sprintf("\tyear = %s", wrap(defaultIfEmpty(bookletYear, "2002"))))

	// OPTIONAL
	if len(bookletEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor = %s", wrap(strings.Join(bookletEditors, " and "))))
	}
	if bookletVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume = %s", wrap(bookletVolume)))
	}
	if bookletNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber = %s", wrap(bookletNumber)))
	}
	if bookletSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries = %s", wrap(bookletSeries)))
	}
	if bookletOrganisation != "" {
		fields = append(fields, fmt.Sprintf("\torganisation = %s", wrap(bookletOrganisation)))
	}
	if bookletMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth = %s", wrap(bookletMonth)))
	}
	if bookletNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote = %s", wrap(bookletNote)))
	}

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")

	return sb.String()

}

package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatInproceedingsBibtex(
	inproceedingsCiteKey string,
	inproceedingsAuthors []string,
	inproceedingsTitle string,
	inproceedingsBooktTitle string,
	inproceedingsYear string,
	inproceedingsEditors []string,
	inproceedingsVolume string,
	inproceedingsNumber string,
	inproceedingsSeries string,
	inproceedingsPages string,
	inproceedingsAddress string,
	inprocceddingsMonth string,
	inproceedingsOrganisation string,
	inproceedingsPublisher string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@inproceedings{")
	if inproceedingsCiteKey != "" {
		sb.WriteString(inproceedingsCiteKey)
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
	if len(inproceedingsAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap("author", strings.Join(inproceedingsAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle        = %s", wrap("title", defaultIfEmpty(inproceedingsTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tbooktitle    = %s", wrap("booktitle", defaultIfEmpty(inproceedingsBooktTitle, "<Book Title>"))))
	fields = append(fields, fmt.Sprintf("\tyear         = %s", wrap("year", defaultIfEmpty(inproceedingsYear, "<2002>"))))

	// OPTIONAL
	if len(inproceedingsEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor       = %s", wrap("editor", strings.Join(inproceedingsEditors, " and "))))
	}
	if inproceedingsVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume       = %s", wrap("volume", inproceedingsVolume)))
	}
	if inproceedingsNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber       = %s", wrap("number", inproceedingsNumber)))
	}
	if inproceedingsSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries       = %s", wrap("series", inproceedingsSeries)))
	}
	if inproceedingsPages != "" {
		fields = append(fields, fmt.Sprintf("\tpages        = %s", wrap("pages", inproceedingsPages)))
	}
	if inproceedingsAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress      = %s", wrap("address", inproceedingsAddress)))
	}
	if inprocceddingsMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth        = %s", wrap("month", inprocceddingsMonth)))
	}
	if inproceedingsOrganisation != "" {
		fields = append(fields, fmt.Sprintf("\torganisation = %s", wrap("organisation", inproceedingsOrganisation)))
	}
	if inproceedingsPublisher != "" {
		fields = append(fields, fmt.Sprintf("\tpublisher    = %s", wrap("publisher", inproceedingsPublisher)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

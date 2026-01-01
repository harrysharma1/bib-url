package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatProceedingsBibtex(
	proceedingsCiteKey string,
	proceedingsTitle string,
	proceedingsYear string,
	proceedingsEditors []string,
	proceedingsVolume string,
	proceedingsNumber string,
	proceedingsSeries string,
	proceedingsAddress string,
	proceedingsMonth string,
	proceedingsPublisher string,
	braces bool) string {

	var sb strings.Builder

	sb.WriteString("@proceedings{")

	if proceedingsCiteKey != "" {
		sb.WriteString(proceedingsCiteKey)
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
	fields = append(fields, fmt.Sprintf("\ttitle     = %s", wrap("title", defaultIfEmpty(proceedingsTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tyear      = %s", wrap("year", defaultIfEmpty(proceedingsYear, "<2002>"))))
	// OPTIONAL
	if len(proceedingsEditors) > 0 {
		fields = append(fields, fmt.Sprintf("\teditor    = %s", wrap("editor", strings.Join(proceedingsEditors, " and "))))
	}
	if proceedingsVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume    = %s", wrap("volume", proceedingsVolume)))
	}
	if proceedingsNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber    = %s", wrap("number", proceedingsNumber)))
	}
	if proceedingsSeries != "" {
		fields = append(fields, fmt.Sprintf("\tseries    = %s", wrap("series", proceedingsSeries)))
	}
	if proceedingsAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress   = %s", wrap("address", proceedingsAddress)))
	}
	if proceedingsMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth     = %s", wrap("month", proceedingsMonth)))
	}
	if proceedingsPublisher != "" {
		fields = append(fields, fmt.Sprintf("\tpublisher = %s", wrap("publisher", proceedingsPublisher)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatTechReportBibtex(
	techreportCiteKey string,
	techreportTitle string,
	techreportAuthors []string,
	techreportInstitution string,
	techreportAddress string,
	techreportNumber string,
	techreportYear string,
	techreportMonth string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@techreport{")

	if techreportCiteKey != "" {
		sb.WriteString(techreportCiteKey)
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
	fields = append(fields, fmt.Sprintf("\ttitle       = %s", wrap("title", defaultIfEmpty(techreportTitle, "<Title>"))))
	if len(techreportAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor      = %s", wrap("author", strings.Join(techreportAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor      = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\tinstitution = %s", wrap("institution", defaultIfEmpty(techreportInstitution, "<Institution>"))))
	fields = append(fields, fmt.Sprintf("\taddress     = %s", wrap("address", defaultIfEmpty(techreportAddress, "<Address>"))))
	fields = append(fields, fmt.Sprintf("\tnumber      = %s", wrap("number", defaultIfEmpty(techreportNumber, "<50>"))))
	fields = append(fields, fmt.Sprintf("\tyear        = %s", wrap("year", defaultIfEmpty(techreportYear, "<2002>"))))
	fields = append(fields, fmt.Sprintf("\tmonth       = %s", wrap("month", defaultIfEmpty(techreportMonth, "<jul>"))))
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

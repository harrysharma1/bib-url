package helper

import (
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
	fields := []Field{}

	// REQUIRED
	fields = append(fields, Field{"title", defaultIfEmpty(techreportTitle, "<Title>")})
	if len(techreportAuthors) > 0 {
		fields = append(fields, Field{"author", strings.Join(techreportAuthors, " and ")})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"institution", defaultIfEmpty(techreportInstitution, "<Institution>")})
	fields = append(fields, Field{"address", defaultIfEmpty(techreportAddress, "<Address>")})
	fields = append(fields, Field{"number", defaultIfEmpty(techreportNumber, "<Number>")})
	fields = append(fields, Field{"year", defaultIfEmpty(techreportYear, "<2002>")})
	fields = append(fields, Field{"month", defaultIfEmpty(techreportMonth, "<jul>")})
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

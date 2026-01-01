package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatManualBibtex(
	manualCiteKey string,
	manualTitle string,
	manualYear string,
	manualAuthors []string,
	manualOrganisation string,
	manualAddress string,
	manualEdition string,
	manualMonth string,
	manualNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@manual{")

	if manualCiteKey != "" {
		sb.WriteString(manualCiteKey)
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
	fields = append(fields, fmt.Sprintf("\ttitle        = %s", wrap(defaultIfEmpty(manualTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tyear         = %s", wrap(defaultIfEmpty(manualYear, "<2002>"))))
	// OPTIONAL
	if len(manualAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor       = %s", wrap(strings.Join(manualAuthors, " and "))))
	}

	if manualOrganisation != "" {
		fields = append(fields, fmt.Sprintf("\torganisation = %s", wrap(manualOrganisation)))
	}

	if manualAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress      = %s", wrap(manualAddress)))
	}

	if manualEdition != "" {
		fields = append(fields, fmt.Sprintf("\tedition      = %s", wrap(manualEdition)))
	}

	if manualMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth        = %s", wrap(manualMonth)))
	}

	if manualNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote         = %s", wrap(manualNote)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

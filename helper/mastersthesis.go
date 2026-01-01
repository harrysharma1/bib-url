package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatMastersThesisBibtex(
	mastersthesisCiteKey string,
	mastersthesisAuthors []string,
	mastersthesisTitle string,
	mastersthesisSchool string,
	mastersthesisYear string,
	mastersthesisType string,
	mastersthesisAddress string,
	mastersthesisMonth string,
	mastersthesisNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@mastersthesis{")

	if mastersthesisCiteKey != "" {
		sb.WriteString(mastersthesisCiteKey)
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
	if len(mastersthesisAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor  = %s", wrap(strings.Join(mastersthesisAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor  = %s", wrap("<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle   = %s", wrap(defaultIfEmpty(mastersthesisTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tschool  = %s", wrap(defaultIfEmpty(mastersthesisSchool, "<School>"))))
	fields = append(fields, fmt.Sprintf("\tyear    = %s", wrap(defaultIfEmpty(mastersthesisYear, "<2002>"))))
	// OPTIONAL
	if mastersthesisType != "" {
		fields = append(fields, fmt.Sprintf("\ttype    = %s", wrap(mastersthesisType)))
	}
	if mastersthesisAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress = %s", wrap(mastersthesisAddress)))
	}
	if mastersthesisMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth   = %s", wrap(mastersthesisMonth)))
	}
	if mastersthesisNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote    = %s", wrap(mastersthesisNote)))
	}

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatPhDThesisBibtex(
	phdthesisCiteKey string,
	phdthesisAuthors []string,
	phdthesisTitle string,
	phdthesisSchool string,
	phdthesisYear string,
	phdthesisType string,
	phdthesisAddress string,
	phdthesisMonth string,
	phdthesisNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@phdthesis{")

	if phdthesisCiteKey != "" {
		sb.WriteString(phdthesisCiteKey)
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
	if len(phdthesisAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor  = %s", wrap(strings.Join(phdthesisAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor  = %s", wrap("<Lastname, Firstname>")))

	}
	fields = append(fields, fmt.Sprintf("\ttitle   = %s", wrap(defaultIfEmpty(phdthesisTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tschool  = %s", wrap(defaultIfEmpty(phdthesisSchool, "<School>"))))
	fields = append(fields, fmt.Sprintf("\tyear    = %s", wrap(defaultIfEmpty(phdthesisYear, "<2002>"))))
	// OPTIONAL
	if phdthesisType != "" {
		fields = append(fields, fmt.Sprintf("\ttype    = %s", wrap(phdthesisType)))
	}
	if phdthesisAddress != "" {
		fields = append(fields, fmt.Sprintf("\taddress = %s", wrap(phdthesisAddress)))
	}
	if phdthesisMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth   = %s", wrap(phdthesisMonth)))
	}
	if phdthesisNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote    = %s", wrap(phdthesisNote)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

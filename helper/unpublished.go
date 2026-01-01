package helper

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func FormatUnpublishedBibtex(
	unpublishedCiteKey string,
	unpublishedAuthors []string,
	unpublishedtTitle string,
	unpublishedInstitution string,
	unpublishedYear string,
	braces bool) string {

	var sb strings.Builder
	sb.WriteString("@unpublished{")
	if unpublishedCiteKey != "" {
		sb.WriteString(unpublishedCiteKey)
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
	if len(unpublishedAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor      = %s", wrap(strings.Join(unpublishedAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor      = %s", wrap("<Lastname, Firstname")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle       = %s", wrap(defaultIfEmpty(unpublishedtTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tinstitution = %s", wrap(defaultIfEmpty(unpublishedInstitution, "<Institution>"))))
	fields = append(fields, fmt.Sprintf("\tyear        = %s", wrap(defaultIfEmpty(unpublishedYear, "<Year>"))))
	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

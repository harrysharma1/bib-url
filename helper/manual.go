package helper

import (
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
	fields := []Field{}

	// REQUIRED
	fields = append(fields, Field{"title", defaultIfEmpty(manualTitle, "<Title>")})
	fields = append(fields, Field{"year", defaultIfEmpty(manualYear, "<2002>")})
	// OPTIONAL
	if len(manualAuthors) > 0 {
		fields = append(fields, Field{"author", strings.Join(manualAuthors, " and ")})
	}

	if manualOrganisation != "" {
		fields = append(fields, Field{"organisation", manualOrganisation})

	}

	if manualAddress != "" {
		fields = append(fields, Field{"address", manualAddress})
	}

	if manualEdition != "" {
		fields = append(fields, Field{"edition", manualEdition})
	}

	if manualMonth != "" {
		fields = append(fields, Field{"month", manualMonth})
	}

	if manualNote != "" {
		fields = append(fields, Field{"note", manualNote})
	}
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

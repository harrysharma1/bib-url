package format

import (
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

	fields := []Field{}

	// REQUIRED
	fields = append(fields, Field{"title", defaultIfEmpty(proceedingsTitle, "<Title>")})
	fields = append(fields, Field{"year", defaultIfEmpty(proceedingsYear, "<2002>")})
	// OPTIONAL
	if len(proceedingsEditors) > 0 {
		fields = append(fields, Field{"editor", formatEditors(proceedingsEditors)})
	}
	if proceedingsVolume != "" {
		fields = append(fields, Field{"volume", proceedingsVolume})
	}
	if proceedingsNumber != "" {
		fields = append(fields, Field{"number", proceedingsNumber})
	}
	if proceedingsSeries != "" {
		fields = append(fields, Field{"series", proceedingsSeries})
	}
	if proceedingsAddress != "" {
		fields = append(fields, Field{"address", proceedingsAddress})
	}
	if proceedingsMonth != "" {
		fields = append(fields, Field{"month", proceedingsMonth})
	}
	if proceedingsPublisher != "" {
		fields = append(fields, Field{"publisher", proceedingsPublisher})
	}
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

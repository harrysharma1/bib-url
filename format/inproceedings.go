package format

import (
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
	fields := []Field{}

	// REQUIRED
	if len(inproceedingsAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(inproceedingsAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(inproceedingsTitle, "<Title>")})
	fields = append(fields, Field{"booktitle", defaultIfEmpty(inproceedingsBooktTitle, "Book Title")})
	fields = append(fields, Field{"year", defaultIfEmpty(inproceedingsYear, "<2002>")})

	// OPTIONAL
	if len(inproceedingsEditors) > 0 {
		fields = append(fields, Field{"editor", formatEditors(inproceedingsEditors)})
	}
	if inproceedingsVolume != "" {
		fields = append(fields, Field{"volume", inproceedingsVolume})
	}
	if inproceedingsNumber != "" {
		fields = append(fields, Field{"number", inproceedingsNumber})
	}
	if inproceedingsSeries != "" {
		fields = append(fields, Field{"series", inproceedingsSeries})
	}
	if inproceedingsPages != "" {
		fields = append(fields, Field{"pages", inproceedingsPages})
	}
	if inproceedingsAddress != "" {
		fields = append(fields, Field{"address", inproceedingsAddress})
	}
	if inprocceddingsMonth != "" {
		fields = append(fields, Field{"month", inprocceddingsMonth})
	}
	if inproceedingsOrganisation != "" {
		fields = append(fields, Field{"organisation", inproceedingsOrganisation})
	}
	if inproceedingsPublisher != "" {
		fields = append(fields, Field{"publisher", inproceedingsPublisher})
	}
	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

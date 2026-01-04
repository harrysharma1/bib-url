package helper

import (
	"bibcli/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func FormatArticleBibtex(
	articleCiteKey string,
	articleAuthors []string,
	articleTitle string,
	articleJournal string,
	articleYear string,
	articleVolume string,
	articleNumber string,
	articlePages string,
	articleMonth string,
	articleNote string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@article{")

	if articleCiteKey != "" {
		sb.WriteString(articleCiteKey)
	} else {
		sb.WriteString(uuid.NewString())
	}
	sb.WriteString(",\n")

	fields := []Field{}

	// REQUIRED
	if len(articleAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(articleAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(articleTitle, "<Title>")})
	fields = append(fields, Field{"journal", (defaultIfEmpty(articleJournal, "<Journal>"))})
	fields = append(fields, Field{"year", defaultIfEmpty(articleYear, "<2002>")})

	// OPTIONAL
	if articleVolume != "" {
		fields = append(fields, Field{"volume", articleVolume})
	}

	if articleNumber != "" {
		fields = append(fields, Field{"number", articleNumber})
	}

	if articlePages != "" {
		fields = append(fields, Field{"pages", articlePages})
	}

	if articleMonth != "" {
		fields = append(fields, Field{"month", articleMonth})
	}

	if articleNote != "" {
		fields = append(fields, Field{"note", articleNote})
	}

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
	sb.WriteString("\n}")
	return sb.String()
}

func ArticleFromDOI(doi string) ([]string, string, string, string, string, string, error) {
	var (
		authors []string
		title   string
		journal string
		year    string
		volume  string
		number  string
	)

	baseUrl := `https://api.crossref.org/works`

	res, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, doi))
	if err != nil {
		return []string{}, "", "", "", "", "", err
	}

	if res.StatusCode == 404 {
		return []string{}, "", "", "", "", "", errors.New("work identified does not exist")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{}, "", "", "", "", "", err
	}

	var work models.Works
	json.Unmarshal(body, &work)
	for i := range len(work.Message.Author) {
		authors = append(authors, fmt.Sprintf("%s %s", work.Message.Author[i].Given, work.Message.Author[i].Family))
	}

	title += work.Message.Title[0]
	journal += work.Message.ContainerTitle[0]
	year += fmt.Sprintf("%d", work.Message.Created.DateParts[0][0])
	volume += work.Message.Volume
	number += work.Message.Issue

	return authors, title, journal, year, volume, number, nil
}

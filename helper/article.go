package helper

import (
	"bibcli/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
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
	if len(articleAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor   = %s", wrap("author", strings.Join(articleAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor   = %s", wrap("author", "<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle    = %s", wrap("title", defaultIfEmpty(articleTitle, "<Title>"))))

	fields = append(fields, fmt.Sprintf("\tjournal  = %s", wrap("journal", defaultIfEmpty(articleJournal, "<Journal>"))))

	fields = append(fields, fmt.Sprintf("\tyear     = %s", wrap("year", defaultIfEmpty(articleYear, "<2002>"))))

	// OPTIONAL
	if articleVolume != "" {
		fields = append(fields, fmt.Sprintf("\tvolume   = %s", wrap("volume", articleVolume)))
	}

	if articleNumber != "" {
		fields = append(fields, fmt.Sprintf("\tnumber   = %s", wrap("number", articleNumber)))
	}

	if articlePages != "" {
		fields = append(fields, fmt.Sprintf("\tpages    = %s", wrap("pages", articlePages)))
	}

	if articleMonth != "" {
		fields = append(fields, fmt.Sprintf("\tmonth    = %s", wrap("month", articleMonth)))
	}

	if articleNote != "" {
		fields = append(fields, fmt.Sprintf("\tnote     = %s", wrap("note", articleNote)))
	}
	sb.WriteString(strings.Join(fields, ",\n"))
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

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

func FormatBookBibtex(
	bookCiteKey string,
	bookAuthors []string,
	bookTitle string,
	bookPublisher string,
	bookAddress string,
	bookYear string,
	braces bool) string {
	var sb strings.Builder
	sb.WriteString("@book{")

	if bookCiteKey != "" {
		sb.WriteString(bookCiteKey)
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
	if len(bookAuthors) > 0 {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap(strings.Join(bookAuthors, " and "))))
	} else {
		fields = append(fields, fmt.Sprintf("\tauthor    = %s", wrap("<Lastname, Firstname>")))
	}
	fields = append(fields, fmt.Sprintf("\ttitle     = %s", wrap(defaultIfEmpty(bookTitle, "<Title>"))))
	fields = append(fields, fmt.Sprintf("\tpublisher = %s", wrap(defaultIfEmpty(bookPublisher, "<Publisher>"))))
	fields = append(fields, fmt.Sprintf("\taddress   = %s", wrap(defaultIfEmpty(bookAddress, "<Address>"))))
	fields = append(fields, fmt.Sprintf("\tyear      = %s", wrap(defaultIfEmpty(bookYear, "<Year>"))))

	sb.WriteString(strings.Join(fields, ",\n"))
	sb.WriteString("\n}")

	return sb.String()
}

func BookFromISBN(isbn string) ([]string, string, string, string, error) {
	var (
		authors   []string
		title     string
		publisher string
		year      string
	)
	openlibrary := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)

	res, err := http.Get(openlibrary)
	if err != nil {
		return []string{""}, "", "", "", err
	}

	if res.StatusCode == 404 {
		return []string{""}, "", "", "", errors.New("book not found")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []string{""}, "", "", "", err
	}

	var bookInfo models.BookInfo
	json.Unmarshal(body, &bookInfo)

	for i := range bookInfo.Authors {
		authorUrl := fmt.Sprintf("https://openlibrary.org%s.json", bookInfo.Authors[i].Key)
		r, err := http.Get(authorUrl)
		if err != nil {
			continue
		}
		b, err := io.ReadAll(r.Body)
		if err != nil {
			continue
		}
		var author models.Author
		json.Unmarshal(b, &author)
		authors = append(authors, author.Name)
	}

	title += bookInfo.FullTitle
	publisher += bookInfo.Publishers[0]
	if !strings.Contains(bookInfo.PublishDate, ",") {
		year += bookInfo.PublishDate
	} else {
		tmpYear := strings.Split(bookInfo.PublishDate, ",")[1]

		year += strings.TrimPrefix(tmpYear, " ")
	}
	return authors, title, publisher, year, nil

}

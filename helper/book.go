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
	fields := []Field{}

	// REQUIRED
	if len(bookAuthors) > 0 {
		fields = append(fields, Field{"author", formatAuthors(bookAuthors)})
	} else {
		fields = append(fields, Field{"author", "<Lastname, Firstname>"})
	}
	fields = append(fields, Field{"title", defaultIfEmpty(bookTitle, "<Title>")})
	fields = append(fields, Field{"publisher", defaultIfEmpty(bookPublisher, "<Publisher>")})
	fields = append(fields, Field{"address", defaultIfEmpty(bookAddress, "<Address>")})
	fields = append(fields, Field{"year", defaultIfEmpty(bookYear, "<2002>")})

	sb.WriteString(strings.Join(formatFields(fields, braces), ",\n"))
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

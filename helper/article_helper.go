package helper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"
)

func FormatArticleBibtex(articleCiteKey string, articleAuthors []string, articleTitle string, articleJournal string, articleYear int, articleVolume string, articleNumber string, articlePages string) string {
	var ret_string string
	ret_string += "@article{"

	if articleCiteKey == "uuid.uuid4()" {
		articleCiteKey = ""
	}

	// CitationArticle
	if articleCiteKey != "" {
		ret_string += articleCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(articleAuthors) > 0 {
		ret_string += `	author   = "`
		for i, articleAuthor := range articleAuthors {
			ret_string += articleAuthor
			if i < len(articleAuthors)-1 {
				ret_string += " and "
			}
		}
		ret_string += `"` + "\n"
	} else {
		ret_string += `	author   = "<Lastname>, <FirstName>"` + ",\n"
	}

	// title
	if articleTitle != "" {
		ret_string += `	title    = "` + articleTitle + `"` + ",\n"
	} else {
		ret_string += `	title    = "<Example Title: Please Change>"` + ",\n"
	}

	// journal
	if articleJournal != "" {
		ret_string += `	journal  = "` + articleJournal + `"` + ",\n"
	} else {
		ret_string += `	journal  = "<Example Journal>"` + ",\n"
	}

	// year
	if strconv.Itoa(articleYear) != "" {
		ret_string += `	year     = ` + strconv.Itoa(articleYear) + ",\n"
	} else {
		ret_string += `	year     = 2002` + ",\n"
	}

	// volume
	if articleVolume != "" {
		ret_string += `	volume   = "` + articleVolume + `"` + ",\n"
	} else {
		ret_string += `	volume   = "1"` + ",\n"
	}

	// number
	if articleNumber != "" {
		ret_string += `	number   = "` + articleNumber + `"` + ",\n"
	} else {
		ret_string += `	number   = "1"` + ",\n"
	}

	// pages
	if articlePages != "" {
		ret_string += `	pages    = "` + articlePages + `"` + "\n"
	} else {
		ret_string += `	pages    = "1--10"` + "\n"
	}
	ret_string += "}"
	return ret_string
}

func GetArticleDataFromUrl(articleUrl string) error {
	res, err := http.Get(articleUrl)
	if err != nil {
		return fmt.Errorf("error fetching %s", articleUrl)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("error status code %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("error scraping web page")
	}
	// author
	author, authorExists := doc.Find("meta[name='author']").Attr("content")
	if authorExists {
		fmt.Println(author)
	} else {
		fmt.Println("NOT FOUND")
	}

	// title
	title, titleExists := doc.Find("meta[property='og:title']").Attr("content")
	if titleExists {
		fmt.Println(title)
	} else {
		fmt.Println("NOT FOUND")
	}
	return nil
}

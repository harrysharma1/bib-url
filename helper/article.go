package helper

import (
	"bibcli/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func FormatArticleBibtex(articleCiteKey string, articleAuthors []string, articleTitle string, articleJournal string, articleYear string, articleVolume string, articleNumber string, articlePages string, braces bool) string {
	var ret_string = "@article{"

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
		if braces {
			ret_string += fmt.Sprintf("	author   = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author   = %s", speechmarks)
		}
		for i, articleAuthor := range articleAuthors {
			ret_string += articleAuthor
			if i < len(articleAuthors)-1 {
				ret_string += " and "
			}
		}
		if braces {
			ret_string += fmt.Sprintf("%s,\n", braces_close)
		} else {
			ret_string += fmt.Sprintf("%s,\n", speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	author   = %s<Lastname, FirstName>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author   = %s<Lastname, FirstName>%s,\n", speechmarks, speechmarks)
		}
	}

	// title
	if articleTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title    = %s%s%s,\n", braces_open, articleTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title    = %s%s%s,\n", speechmarks, articleTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title    = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title    = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// journal
	if articleJournal != "" {
		if braces {
			ret_string += fmt.Sprintf("	journal  = %s%s%s,\n", braces_open, articleJournal, braces_close)
		} else {
			ret_string += fmt.Sprintf("	journal  = %s%s%s,\n", speechmarks, articleJournal, speechmarks)
		}

	} else {
		if braces {
			ret_string += fmt.Sprintf("	journal  = %s<Example Journal: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	journal  = %s<Example Journal: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if articleYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year     = %s%s%s,\n", braces_open, articleYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year     = %s,\n", articleYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year     = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year     = <2002: Please Change>,\n"
		}
	}

	// volume
	if articleVolume != "" {
		if braces {
			ret_string += fmt.Sprintf("	volume   = %s%s%s,\n", braces_open, articleVolume, braces_close)
		} else {
			ret_string += fmt.Sprintf("	volume   = %s%s%s,\n", speechmarks, articleVolume, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	volume   = %s<1: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	volume   = %s<1: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// number
	if articleNumber != "" {
		if braces {
			ret_string += fmt.Sprintf("	number   = %s%s%s,\n", braces_open, articleNumber, braces_close)
		} else {
			ret_string += fmt.Sprintf("	number   = %s%s%s,\n", speechmarks, articleNumber, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	number   = %s<1: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	number   = %s<1: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// pages
	if articlePages != "" {
		if braces {
			ret_string += fmt.Sprintf("	pages    = %s%s%s\n", braces_open, articlePages, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages    = %s%s%s\n", speechmarks, articlePages, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	pages    = %s<1--10: Please Change>%s\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages    = %s<1--10: Please Change>%s\n", speechmarks, speechmarks)
		}
	}
	ret_string += "}"
	return ret_string
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

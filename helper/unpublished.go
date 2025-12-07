package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatUnpublishedBibtex(unpublishedCiteKey string, unpublishedAuthors []string, unpublishedTitle string, unpublishedYear string, braces bool) string {
	var ret_string = "@unpublished{"

	// CitationUnpublished
	if unpublishedCiteKey != "" {
		ret_string += unpublishedCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(unpublishedAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author   = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author   = %s", speechmarks)
		}
		for i, unpubunpublishedAuthor := range unpublishedAuthors {
			ret_string += unpubunpublishedAuthor
			if i < len(unpublishedAuthors)-1 {
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
	if unpublishedTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title    = %s%s%s,\n", braces_open, unpublishedTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title    = %s%s%s,\n", speechmarks, unpublishedTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title    = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title    = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if unpublishedYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year     = %s%s%s,\n", braces_open, unpublishedYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year     = %s,\n", unpublishedYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year     = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year     = <2002: Please Change>,\n"
		}
	}
	ret_string += "}"
	return ret_string
}

package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatBookletBibtex(bookletCiteKey string, bookletTitle string, bookletAuthors []string, bookletHowPublished string, bookletMonth string, bookletYear string, braces bool) string {
	var ret_string = "@booklet{"

	// CitationBooklet
	if bookletCiteKey != "" {
		ret_string += bookletCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// title
	if bookletTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", braces_open, bookletTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", speechmarks, bookletTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// author
	if len(bookletAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author       = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author       = %s", speechmarks)
		}
		for i, bookletAuthor := range bookletAuthors {
			ret_string += bookletAuthor
			if i < len(bookletAuthors)-1 {
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
			ret_string += fmt.Sprintf("	author       = %s<Lastname, Firstname>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author       = %s<Lastname, Firstname>%s,\n", speechmarks, speechmarks)
		}
	}

	// howpublished
	if bookletHowPublished != "" {
		if braces {
			ret_string += fmt.Sprintf("	howpublished = %s%s%s,\n", braces_open, bookletHowPublished, braces_close)
		} else {
			ret_string += fmt.Sprintf("	howpublished = %s%s%s,\n", speechmarks, bookletHowPublished, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	howpublished = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	howpublished = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// month
	if bookletMonth != "" {
		if braces {
			ret_string += fmt.Sprintf("	month        = %s%s%s,\n", braces_open, bookletMonth, braces_close)
		} else {
			ret_string += fmt.Sprintf(" month        = %s,\n", bookletMonth)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	month        = %s<sep: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	month        = <sep: Please Change>,\n"
		}
	}

	// year
	if bookletYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year         = %s%s%s\n", braces_open, bookletYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year         = %s%s%s\n", speechmarks, bookletYear, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year         = %s<2002: Please Change>%s\n", braces_open, braces_close)
		} else {
			ret_string += "	year         = <2002: Please Change>\n"
		}
	}

	ret_string += "}"
	return ret_string
}

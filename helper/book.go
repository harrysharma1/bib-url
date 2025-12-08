package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatBookBibtex(bookCiteKey string, bookAuthors []string, bookTitle string, bookPublisher string, bookAddress string, bookYear string, braces bool) string {
	var ret_string = "@book{"

	// CitationBook
	if bookCiteKey != "" {
		ret_string += bookCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(bookAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author    = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author    = %s", speechmarks)
		}
		for i, bookAuthor := range bookAuthors {
			ret_string += bookAuthor
			if i < len(bookAuthors)-1 {
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
			ret_string += fmt.Sprintf("	author    = %s<Lastname, Firstname>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author    = %s<Lastname, Firstname>%s,\n", speechmarks, speechmarks)
		}
	}

	// title
	if bookTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", braces_open, bookTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", speechmarks, bookTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// publisher
	if bookPublisher != "" {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", braces_open, bookPublisher, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", speechmarks, bookPublisher, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s<Example Publisher: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s<Example Publisher: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if bookAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", braces_open, bookAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", speechmarks, bookAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s<Example Address: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s<Example Address: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if bookYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s%s%s\n", braces_open, bookYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year      = %s\n", bookYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s<2002: Please Change>%s\n", braces_open, braces_close)
		} else {
			ret_string += "	year      = <2002: Please Change>\n"
		}
	}
	ret_string += "}"
	return ret_string

}

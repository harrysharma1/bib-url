package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatInbookBibtex(inbookCiteKey string, inbookAuthors []string, inbookTitle string, inbookBookTitle string, inbookYear string, inbookPublisher string, inbookAddress string, inbookPages string, braces bool) string {
	var ret_string = "@inbook{"

	// CitationInbook
	if inbookCiteKey != "" {
		ret_string += inbookCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(inbookAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author    = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author    = %s", speechmarks)
		}
		for i, inbookAuthor := range inbookAuthors {
			ret_string += inbookAuthor
			if i < len(inbookAuthors)-1 {
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
	if inbookTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", braces_open, inbookTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", speechmarks, inbookTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// booktitle
	if inbookBookTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", braces_open, inbookBookTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", speechmarks, inbookBookTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if inbookYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s%s%s\n", braces_open, inbookYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year      = %s\n", inbookYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year      = <2002: Please Change>,\n"
		}
	}

	// publisher
	if inbookPublisher != "" {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", braces_open, inbookPublisher, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", speechmarks, inbookPublisher, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if inbookAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", braces_open, inbookAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", speechmarks, inbookAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// pages
	if inbookPages != "" {
		if braces {
			ret_string += fmt.Sprintf("	pages     = %s%s%s\n", braces_open, inbookPages, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages     = %s%s%s\n", speechmarks, inbookPages, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	pages     = %s<Example How Published: Please Change>%s\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages     = %s<Example How Published: Please Change>%s\n", speechmarks, speechmarks)
		}
	}
	ret_string += "}"

	return ret_string
}

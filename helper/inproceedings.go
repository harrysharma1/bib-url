package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatInproceedingsBibtex(inproceedingsCiteKey string, inproceedingsAuthors []string, inproceedingsTitle string, inproceedingsBookTitle string, inproceedingsSeries string, inproceedingsYear string, inproceedingsPages string, inproceedingsPublisher string, inproceedingsAddress string, braces bool) string {
	var ret_string = "@inproceedings{"

	// CitationInproceedings
	if inproceedingsCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", inproceedingsCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(inproceedingsAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author    = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author    = %s", speechmarks)
		}
		for i, inproceedingsAuthor := range inproceedingsAuthors {
			ret_string += inproceedingsAuthor
			if i < len(inproceedingsAuthors)-1 {
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
	if inproceedingsTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", braces_open, inproceedingsTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", speechmarks, inproceedingsTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// booktitle
	if inproceedingsBookTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", braces_open, inproceedingsBookTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", speechmarks, inproceedingsBookTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Book Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Book Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// series
	if inproceedingsSeries != "" {
		if braces {
			ret_string += fmt.Sprintf("	series    = %s%s%s,\n", braces_open, inproceedingsSeries, braces_close)
		} else {
			ret_string += fmt.Sprintf("	series    = %s%s%s,\n", speechmarks, inproceedingsSeries, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	series    = %s<Example Series: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	series    = %s<Example Series: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if inproceedingsYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s%s%s,\n", braces_open, inproceedingsYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year      = %s,\n", inproceedingsYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s<Example Year: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year      = <Example Year: Please Change>,\n"
		}
	}

	// pages
	if inproceedingsPages != "" {
		if braces {
			ret_string += fmt.Sprintf("	pages     = %s%s%s,\n", braces_open, inproceedingsPages, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages     = %s%s%s,\n", speechmarks, inproceedingsPages, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	pages     = %s<Example Pages: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages     = %s<Example Pages: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// publisher
	if inproceedingsPublisher != "" {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", braces_open, inproceedingsPublisher, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", speechmarks, inproceedingsPublisher, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s<Example Publisher: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s<Example Publisher: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if inproceedingsAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s%s%s\n", braces_open, inproceedingsAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s%s%s\n", speechmarks, inproceedingsAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s<Example Address: Please Change>%s\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s<Example Address: Please Change>%s\n", speechmarks, speechmarks)
		}
	}

	ret_string += "}"
	return ret_string
}

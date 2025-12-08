package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatIncollectionBibtex(incollectionCiteKey string, incollectionAuthors []string, incollectionEditors []string, incollectionTitle string, incollectionBookTitle string, incollectionYear string, incollectionPublisher string, incollectionAddress string, incollectionPages string, braces bool) string {
	var ret_string = "@incollection{"

	// CitationArticle
	if incollectionCiteKey != "" {
		ret_string += incollectionCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(incollectionAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author    = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author    = %s", speechmarks)
		}
		for i, incollectionAuthor := range incollectionAuthors {
			ret_string += incollectionAuthor
			if i < len(incollectionAuthors)-1 {
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
			ret_string += fmt.Sprintf("	author    = %s<Lastname, FirstName>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author    = %s<Lastname, FirstName>%s,\n", speechmarks, speechmarks)
		}
	}

	// editor
	if len(incollectionEditors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	editor    = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	editor    = %s", speechmarks)
		}
		for i, incollectionEditor := range incollectionEditors {
			ret_string += incollectionEditor
			if i < len(incollectionEditors)-1 {
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
			ret_string += fmt.Sprintf("	editor    = %s<Lastname, FirstName>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	editor    = %s<Lastname, FirstName>%s,\n", speechmarks, speechmarks)
		}
	}

	// title
	if incollectionTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", braces_open, incollectionTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s%s%s,\n", speechmarks, incollectionTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title     = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// booktitle
	if incollectionBookTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", braces_open, incollectionBookTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s%s%s,\n", speechmarks, incollectionBookTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Book Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	booktitle = %s<Example Book Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if incollectionYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s%s%s,\n", braces_open, incollectionYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year      = %s,\n", incollectionYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year      = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year      = <2002: Please Change>,\n"
		}
	}

	// publisher
	if incollectionPublisher != "" {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", braces_open, incollectionPublisher, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s%s%s,\n", speechmarks, incollectionPublisher, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	publisher = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	publisher = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if incollectionAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", braces_open, incollectionAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s%s%s,\n", speechmarks, incollectionAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address   = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address   = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// pages
	if incollectionPages != "" {
		if braces {
			ret_string += fmt.Sprintf("	pages     = %s%s%s\n", braces_open, incollectionPages, braces_close)
		} else {
			ret_string += fmt.Sprintf("	pages     = %s%s%s\n", speechmarks, incollectionPages, speechmarks)
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

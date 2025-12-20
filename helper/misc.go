package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatMiscBibtex(miscCiteKey string, miscTitle string, miscAuthors []string, miscHowPublished string, miscYear string, miscNote string, braces bool) string {
	var ret_string = "@misc{"

	// CitationMisc
	if miscCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", miscCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// title
	if miscTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", braces_open, miscTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", speechmarks, miscTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// author
	if len(miscAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author       = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author       = %s", speechmarks)
		}
		for i, miscAuthor := range miscAuthors {
			ret_string += miscAuthor
			if i < len(miscAuthors)-1 {
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
	if miscHowPublished != "" {
		if braces {
			ret_string += fmt.Sprintf("	howpublished = %s%s%s,\n", braces_open, miscHowPublished, braces_close)
		} else {
			ret_string += fmt.Sprintf("	howpublished = %s%s%s,\n", speechmarks, miscHowPublished, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	howpublished = %s<Example How Published: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	howpublished = %s<Example How Published: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if miscYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year         = %s%s%s,\n", braces_open, miscYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year         = %s,\n", miscYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year         = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year         = <2002: Please Change>,\n"
		}
	}

	// note
	if miscNote != "" {
		if braces {
			ret_string += fmt.Sprintf("	note         = %s%s%s,\n", braces_open, miscNote, braces_close)
		} else {
			ret_string += fmt.Sprintf("	note         = %s%s%s,\n", speechmarks, miscNote, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	note         = %s<Example Note: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	note         = %s<Example NOte: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}
	ret_string += "}"
	return ret_string
}

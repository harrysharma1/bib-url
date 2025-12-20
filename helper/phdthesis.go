package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatPhDThesisBibtex(phdthesisCiteKey string, phdthesisAuthors []string, phdthesisTitle string, phdthesisSchool string, phdthesisAddress string, phdthesisYear string, phdthesisMonth string, braces bool) string {
	var ret_string = "@phdthesis{"

	// CitationPhDThesis
	if phdthesisCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", phdthesisCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// authors
	if len(phdthesisAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author  = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author  = %s", speechmarks)
		}
		for i, phdthesisAuthor := range phdthesisAuthors {
			ret_string += phdthesisAuthor
			if i < len(phdthesisAuthors)-1 {
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
			ret_string += fmt.Sprintf("	author  = %s<Lastname, Firstname>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author  = %s<Lastname, Firstname>%s,\n", speechmarks, speechmarks)
		}
	}

	// title
	if phdthesisTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title   = %s%s%s,\n", braces_open, phdthesisTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title   = %s%s%s,\n", speechmarks, phdthesisTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title   = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title   = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// school
	if phdthesisSchool != "" {
		if braces {
			ret_string += fmt.Sprintf("	school  = %s%s%s,\n", braces_open, phdthesisSchool, braces_close)
		} else {
			ret_string += fmt.Sprintf("	school  = %s%s%s,\n", speechmarks, phdthesisSchool, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	school  = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	school  = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if phdthesisAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address = %s%s%s,\n", braces_open, phdthesisAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address = %s%s%s,\n", speechmarks, phdthesisAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if phdthesisYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year    = %s%s%s\n", braces_open, phdthesisYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year    = %s\n", phdthesisYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year    = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year    = <2002: Please Change>,\n"
		}
	}

	// month
	if phdthesisMonth != "" {
		if braces {
			ret_string += fmt.Sprintf("	month   = %s%s%s\n", braces_open, phdthesisMonth, braces_close)
		} else {
			ret_string += fmt.Sprintf("	month   = %s\n", phdthesisMonth)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	month   = %s<month: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	month   = <sep: Please Change>,\n"
		}
	}
	ret_string += "}"
	return ret_string
}

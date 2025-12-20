package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatMastersThesisBibtex(mastersthesisCiteKey string, mastersthesisAuthors []string, mastersthesisTitle string, mastersthesisSchool string, mastersthesisYear string, mastersthesisAddress string, mastersthesisMonth string, braces bool) string {
	var ret_string = "@mastersthesis{"

	// CitationMastersThesis
	if mastersthesisCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", mastersthesisCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(mastersthesisAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author  = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author  = %s", speechmarks)
		}
		for i, mastersthesisAuthor := range mastersthesisAuthors {
			ret_string += mastersthesisAuthor
			if i < len(mastersthesisAuthors)-1 {
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
	if mastersthesisTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title   = %s%s%s,\n", braces_open, mastersthesisTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title   = %s%s%s,\n", speechmarks, mastersthesisTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title   = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title   = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// school
	if mastersthesisSchool != "" {
		if braces {
			ret_string += fmt.Sprintf("	school  = %s%s%s,\n", braces_open, mastersthesisSchool, braces_close)
		} else {
			ret_string += fmt.Sprintf("	school  = %s%s%s,\n", speechmarks, mastersthesisSchool, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	school  = %s<Example School: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	school  = %s<Example School: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if mastersthesisYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year    = %s%s%s,\n", braces_open, mastersthesisYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year    = %s,\n", mastersthesisYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year    = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year    = <2002: Please Change>,\n"
		}
	}

	// address
	if mastersthesisAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address = %s%s%s,\n", braces_open, mastersthesisAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address = %s%s%s,\n", speechmarks, mastersthesisAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address = %s<Example Address: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address = %s<Example Address: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// month
	if mastersthesisMonth != "" {
		if braces {
			ret_string += fmt.Sprintf("	month   = %s%s%s,\n", braces_open, mastersthesisMonth, braces_close)
		} else {
			ret_string += fmt.Sprintf("	month   = %s%s%s,\n", speechmarks, mastersthesisMonth, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	month   = %s<Example Month: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	month   = %s<Example Month: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}
	ret_string += "}"

	return ret_string
}

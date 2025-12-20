package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatTechReportBibtex(techreportCiteKey string, techreportTitle string, techreportAuthors []string, techreportInstitution string, techreportAddress string, techreportNumber string, techreportYear string, techreportMonth string, braces bool) string {
	var ret_string = "@techreport{"

	// CitationTechReport
	if techreportCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", techreportCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// title
	if techreportTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title       = %s%s%s,\n", braces_open, techreportTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title       = %s%s%s,\n", speechmarks, techreportTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title       = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title       = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// author
	if len(techreportAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author      = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author      = %s", speechmarks)
		}
		for i, techreportAuthor := range techreportAuthors {
			ret_string += techreportAuthor
			if i < len(techreportAuthors)-1 {
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
			ret_string += fmt.Sprintf("	author      = %s<Lastname, Firstname>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	author      = %s<Lastname, Firstname>%s,\n", speechmarks, speechmarks)
		}
	}

	// institution
	if techreportInstitution != "" {
		if braces {
			ret_string += fmt.Sprintf("	institution = %s%s%s,\n", braces_open, techreportInstitution, braces_close)
		} else {
			ret_string += fmt.Sprintf("	institution = %s%s%s,\n", speechmarks, techreportInstitution, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	institution = %s<Example Institution: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	institution = %s<Example Institution: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if techreportAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address     = %s%s%s,\n", braces_open, techreportAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address     = %s%s%s,\n", speechmarks, techreportAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address     = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address     = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// number
	if techreportNumber != "" {
		if braces {
			ret_string += fmt.Sprintf("	number      = %s%s%s,\n", braces_open, techreportNumber, braces_close)
		} else {
			ret_string += fmt.Sprintf("	number      = %s%s%s,\n", speechmarks, techreportNumber, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	number      = %s<Example Number: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	number      = %s<Example Number: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if techreportYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year        = %s%s%s,\n", braces_open, techreportYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year        = %s,\n", techreportYear)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	year        = %s<2002: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	year        = <2002: Please Change>,\n"
		}
	}

	// month
	if techreportMonth != "" {
		if braces {
			ret_string += fmt.Sprintf("	month       = %s%s%s,\n", braces_open, techreportMonth, braces_close)
		} else {
			ret_string += fmt.Sprintf("	month       = %s,\n", techreportMonth)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	month      = %s<sep: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += "	month       = <sep: Please Change>,\n"
		}
	}
	ret_string += "}"
	return ret_string
}

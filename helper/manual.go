package helper

import (
	"fmt"

	"github.com/google/uuid"
)

func FormatManualBibtex(manualCiteKey string, manualTitle string, manualAuthors []string, manualOrganisation string, manualAddress string, manualYear string, braces bool) string {
	var ret_string = "@manual{"

	// CitationManual
	if manualCiteKey != "" {
		ret_string += fmt.Sprintf("%s,\n", manualCiteKey)
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// title
	if manualTitle != "" {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", braces_open, manualTitle, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s%s%s,\n", speechmarks, manualTitle, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	title        = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// author
	if len(manualAuthors) > 0 {
		if braces {
			ret_string += fmt.Sprintf("	author       = %s", braces_open)
		} else {
			ret_string += fmt.Sprintf("	author       = %s", speechmarks)
		}
		for i, manualAuthor := range manualAuthors {
			ret_string += manualAuthor
			if i < len(manualAuthors)-1 {
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

	// organisation
	if manualOrganisation != "" {
		if braces {
			ret_string += fmt.Sprintf("	organisation = %s%s%s,\n", braces_open, manualOrganisation, braces_close)
		} else {
			ret_string += fmt.Sprintf("	organisation = %s%s%s,\n", speechmarks, manualOrganisation, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	organisation = %s<Example Title: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	organisation = %s<Example Title: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// address
	if manualAddress != "" {
		if braces {
			ret_string += fmt.Sprintf("	address      = %s%s%s,\n", braces_open, manualAddress, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address      = %s%s%s,\n", speechmarks, manualAddress, speechmarks)
		}
	} else {
		if braces {
			ret_string += fmt.Sprintf("	address      = %s<Example Address: Please Change>%s,\n", braces_open, braces_close)
		} else {
			ret_string += fmt.Sprintf("	address      = %s<Example Address: Please Change>%s,\n", speechmarks, speechmarks)
		}
	}

	// year
	if manualYear != "" {
		if braces {
			ret_string += fmt.Sprintf("	year         = %s%s%s\n", braces_open, manualYear, braces_close)
		} else {
			ret_string += fmt.Sprintf("	year         = %s\n", manualYear)
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

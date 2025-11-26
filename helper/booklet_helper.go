package helper

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func FormatBookletBibtex(bookletCiteKey string, bookletTitle string, bookletAuthors []string, bookletHowPublished string, bookletMonth string, bookletYear int) string {
	var ret_string = "@booklet{"

	if bookletCiteKey == "uuid.uuid4()" {
		bookletCiteKey = ""
	}

	// CitationBooklet
	if bookletCiteKey != "" {
		ret_string += bookletCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// title
	if bookletTitle != "" {
		ret_string += `	title        = "` + bookletTitle + `"` + ",\n"
	} else {
		ret_string += `	title        = "<Example Title: Please Change>"` + ",\n"
	}

	// author
	if len(bookletAuthors) > 0 {
		ret_string += `	author       = "`
		for i, bookletAuthor := range bookletAuthors {
			ret_string += bookletAuthor
			if i < len(bookletAuthors)-1 {
				ret_string += " and "
			}
		}
		ret_string += `"` + "\n"
	} else {
		ret_string += `	author       = "<Lastname>, <FirstName>"` + ",\n"
	}

	// howpublished
	if bookletHowPublished != "" {
		ret_string += `	howpublished = "` + bookletHowPublished + `"` + ",\n"
	} else {
		ret_string += `	howpublished = "<Example How Published: Please Change>"` + ",\n"
	}

	// month
	if bookletMonth != "" {
		ret_string += `	month        = ` + bookletMonth + ",\n"
	} else {
		ret_string += `	month        = sep`
	}

	// year
	if strconv.Itoa(bookletYear) != "" {
		ret_string += `	year         = ` + strconv.Itoa(bookletYear) + "\n"
	} else {
		ret_string += `	year         = 2002` + "\n"
	}

	ret_string += "}"
	return ret_string
}

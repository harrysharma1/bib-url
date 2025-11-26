package helper

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func FormatBookBibtex(bookCiteKey string, bookAuthors []string, bookTitle string, bookPublisher string, bookAddress string, bookYear int) string {
	var ret_string = "@book{"

	if bookCiteKey == "uuid.uuid4()" {
		bookCiteKey = ""
	}

	// CitationBook
	if bookCiteKey != "" {
		ret_string += bookCiteKey + ",\n"
	} else {
		key := uuid.NewString()
		ret_string += fmt.Sprintf("%s,\n", key)
	}

	// author
	if len(bookAuthors) > 0 {
		ret_string += `	author    = "`
		for i, bookAuthor := range bookAuthors {
			ret_string += bookAuthor
			if i < len(bookAuthors)-1 {
				ret_string += " and "
			}
		}
		ret_string += `"` + "\n"
	} else {
		ret_string += `	author    = "<Lastname>, <FirstName>"` + ",\n"
	}

	// title
	if bookTitle != "" {
		ret_string += `	title     = "` + bookTitle + `"` + ",\n"
	} else {
		ret_string += `	title     = "<Example Title: Please Change>"` + ",\n"
	}

	// publisher
	if bookPublisher != "" {
		ret_string += `	publisher = "` + bookPublisher + `"` + ",\n"
	} else {
		ret_string += ` publisher = "<Example Publisher: Please Change>"` + ",\n"
	}

	// address
	if bookAddress != "" {
		ret_string += `	address   = "` + bookAddress + `"` + ",\n"
	} else {
		ret_string += `	address   = "<Example Address: Please Change" ` + ",\n"
	}

	// year
	if strconv.Itoa(bookYear) != "" {
		ret_string += `	year      = ` + strconv.Itoa(bookYear) + "\n"
	} else {
		ret_string += `	year      = 2002` + "\n"
	}
	ret_string += "}"
	return ret_string

}

package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestInproceedingsRequiredValues(t *testing.T) {
	want := `@inproceedings{test,
	author       = "John Doe",
	title        = "Test Title",
	booktitle    = "Test Booktitle",
	year         = 2002
}`
	have := FormatInproceedingsBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		false)
	assert.Equal(t, want, have, "should be equal")
}
func TestInproceedingsRequiredValuesBraces(t *testing.T) {
	want := `@inproceedings{test,
	author       = {John Doe},
	title        = {Test Title},
	booktitle    = {Test Booktitle},
	year         = {2002}
}`
	have := FormatInproceedingsBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		true)
	assert.Equal(t, want, have, "should be equal")

}

func TestInproceedingsOptionalValues(t *testing.T) {
	want := `@inproceedings{test,
	author       = "John Doe",
	title        = "Test Title",
	booktitle    = "Test Booktitle",
	year         = 2002,
	editor       = "Jane Doe and John Doe",
	volume       = "1",
	number       = "50",
	series       = "Test Series",
	pages        = "1--10",
	address      = "Test Avenue",
	month        = jul,
	organisation = "Test Organisation",
	publisher    = "Test Publisher"
}`
	have := FormatInproceedingsBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"2002",
		[]string{"Jane Doe", "John Doe"},
		"1",
		"50",
		"Test Series",
		"1--10",
		"Test Avenue",
		"jul",
		"Test Organisation",
		"Test Publisher",
		false)
	assert.Equal(t, want, have, "should be equal")
}

func TestInproceedingsOptionalValuesBraces(t *testing.T) {
	want := `@inproceedings{test,
	author       = {John Doe},
	title        = {Test Title},
	booktitle    = {Test Booktitle},
	year         = {2002},
	editor       = {Jane Doe and John Doe},
	volume       = {1},
	number       = {50},
	series       = {Test Series},
	pages        = {1--10},
	address      = {Test Avenue},
	month        = {jul},
	organisation = {Test Organisation},
	publisher    = {Test Publisher}
}`
	have := FormatInproceedingsBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"2002",
		[]string{"Jane Doe", "John Doe"},
		"1",
		"50",
		"Test Series",
		"1--10",
		"Test Avenue",
		"jul",
		"Test Organisation",
		"Test Publisher",
		true)
	assert.Equal(t, want, have, "should be equal")
}

// --------------------------------------------------//

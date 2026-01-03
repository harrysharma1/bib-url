package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestIncollectionRequiredValues(t *testing.T) {
	want := `@incollection{test,
	author    = "John Doe",
	title     = "Test Title",
	booktitle = "Test Booktitle",
	publisher = "Test Publisher",
	year      = 2002
}`
	have := FormatIncollectionBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		"",
		false)
	assert.Equal(t, want, have, "should be equal")
}

func TestIncollectionRequiredValuesBraces(t *testing.T) {
	want := `@incollection{test,
	author    = {John Doe},
	title     = {Test Title},
	booktitle = {Test Booktitle},
	publisher = {Test Publisher},
	year      = {2002}
}`
	have := FormatIncollectionBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		"",
		true)
	assert.Equal(t, want, have, "should be equal")
}

func TestIncollectionOptionalValues(t *testing.T) {
	want := `@incollection{test,
	author    = "John Doe",
	title     = "Test Title",
	booktitle = "Test Booktitle",
	publisher = "Test Publisher",
	year      = 2002,
	editor    = "Jane Doe",
	volume    = "1",
	number    = "50",
	series    = "Test Series",
	pages     = "1--10",
	address   = "Test Avenue",
	month     = jul
}`

	have := FormatIncollectionBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{"Jane Doe"},
		"1",
		"50",
		"Test Series",
		"1--10",
		"Test Avenue",
		"jul",
		false)
	assert.Equal(t, want, have, "should be equal")
}
func TestIncollectionOptionalValuesBraces(t *testing.T) {
	want := `@incollection{test,
	author    = {John Doe},
	title     = {Test Title},
	booktitle = {Test Booktitle},
	publisher = {Test Publisher},
	year      = {2002},
	editor    = {Jane Doe},
	volume    = {1},
	number    = {50},
	series    = {Test Series},
	pages     = {1--10},
	address   = {Test Avenue},
	month     = {jul}
}`

	have := FormatIncollectionBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{"Jane Doe"},
		"1",
		"50",
		"Test Series",
		"1--10",
		"Test Avenue",
		"jul",
		true)
	assert.Equal(t, want, have, "should be equal")

}

// --------------------------------------------------//

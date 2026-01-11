package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @incollection subcommand for required values only
func TestIncollectionRequiredValues(t *testing.T) {
	want := `@incollection{test,
	author    = "Doe, John",
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

// Test @incollection subcommand output for required values only (with braces)
func TestIncollectionRequiredValuesBraces(t *testing.T) {
	want := `@incollection{test,
	author    = {Doe, John},
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

// Test @incollection output for optional values
func TestIncollectionOptionalValues(t *testing.T) {
	want := `@incollection{test,
	author    = "Doe, John",
	title     = "Test Title",
	booktitle = "Test Booktitle",
	publisher = "Test Publisher",
	year      = 2002,
	editor    = "Doe, Jane",
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

// Test @incollection subcommand output for optional values (with braces)
func TestIncollectionOptionalValuesBraces(t *testing.T) {
	want := `@incollection{test,
	author    = {Doe, John},
	title     = {Test Title},
	booktitle = {Test Booktitle},
	publisher = {Test Publisher},
	year      = {2002},
	editor    = {Doe, Jane},
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

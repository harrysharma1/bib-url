package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @inbook subcommand output for required values only
func TestInbookRequiredValues(t *testing.T) {
	want := `@inbook{test,
	author    = "Doe, John",
	title     = "Test Title",
	booktitle = "Test Booktitle",
	publisher = "Test Publisher",
	year      = 2002
}`
	have := FormatInbookBibtex(
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
		"",
		"",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @inbook subcommand output for required values only (with braces)
func TestInbookRequiredValuesBraces(t *testing.T) {
	want := `@inbook{test,
	author    = {Doe, John},
	title     = {Test Title},
	booktitle = {Test Booktitle},
	publisher = {Test Publisher},
	year      = {2002}
}`
	have := FormatInbookBibtex(
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
		"",
		"",
		true)
	assert.Equal(t, want, have, "should be equal")
}

// Test @inbook subcommand output for optional values
func TestInbookOptionalValues(t *testing.T) {
	want := `@inbook{test,
	author    = "Doe, John",
	title     = "Test Title",
	booktitle = "Test Booktitle",
	publisher = "Test Publisher",
	year      = 2002,
	editor    = "Doe, Jane",
	volume    = "50",
	number    = "1",
	series    = "Test Series",
	address   = "Test Avenue",
	edition   = "1st",
	month     = jul,
	pages     = "10--100",
	note      = "Test Note"
}`

	have := FormatInbookBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{"Jane Doe"},
		"50",
		"1",
		"Test Series",
		"Test Avenue",
		"1st",
		"jul",
		"10--100",
		"Test Note",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @inbook subcommand output for optional values (with braces)
func TestInbookOptionalValuesBraces(t *testing.T) {
	want := `@inbook{test,
	author    = {Doe, John},
	title     = {Test Title},
	booktitle = {Test Booktitle},
	publisher = {Test Publisher},
	year      = {2002},
	editor    = {Doe, Jane},
	volume    = {50},
	number    = {1},
	series    = {Test Series},
	address   = {Test Avenue},
	edition   = {1st},
	month     = {jul},
	pages     = {10--100},
	note      = {Test Note}
}`

	have := FormatInbookBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test Booktitle",
		"Test Publisher",
		"2002",
		[]string{"Jane Doe"},
		"50",
		"1",
		"Test Series",
		"Test Avenue",
		"1st",
		"jul",
		"10--100",
		"Test Note",
		true)
	assert.Equal(t, want, have, "should be equal")

}

// --------------------------------------------------//

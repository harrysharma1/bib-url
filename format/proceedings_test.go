package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @proceedings subcommand output for required values only
func TestProceedingsRequiredValues(t *testing.T) {
	want := `@proceedings{test,
	title = "Test Title",
	year  = 2002
}`
	have := FormatProceedingsBibtex(
		"test",
		"Test Title",
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

// Test @proceedings subcommand output for required values only (with braces)
func TestProceedingsRequiredValuesBraces(t *testing.T) {
	want := `@proceedings{test,
	title = {Test Title},
	year  = {2002}
}`
	have := FormatProceedingsBibtex(
		"test",
		"Test Title",
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

// Test @proceedings subcommand output for optional values
func TestProceedingsOptionalValues(t *testing.T) {
	want := `@proceedings{test,
	title     = "Test Title",
	year      = 2002,
	editor    = "Doe, Jane",
	volume    = "1",
	number    = "50",
	series    = "Test Series",
	address   = "Test Avenue",
	month     = jul,
	publisher = "Test Publisher"
}`
	have := FormatProceedingsBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{"Jane Doe"},
		"1",
		"50",
		"Test Series",
		"Test Avenue",
		"jul",
		"Test Publisher",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @proceedings subcommand output for optional values (with braces)
func TestProceedingsOpionalValuesBraces(t *testing.T) {
	want := `@proceedings{test,
	title     = {Test Title},
	year      = {2002},
	editor    = {Doe, Jane},
	volume    = {1},
	number    = {50},
	series    = {Test Series},
	address   = {Test Avenue},
	month     = {jul},
	publisher = {Test Publisher}
}`
	have := FormatProceedingsBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{"Jane Doe"},
		"1",
		"50",
		"Test Series",
		"Test Avenue",
		"jul",
		"Test Publisher",
		true)
	assert.Equal(t, want, have, "should be equal")
}

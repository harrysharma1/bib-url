package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @manual subcommand output for required values only
func TestManualRequiredValues(t *testing.T) {
	want := `@manual{test,
	title = "Test Title",
	year  = 2002
}`
	have := FormatManualBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		false)
	assert.Equal(t, want, have, "shouuld be equal")
}

// Test @manual subcommand output for required values only (with braces)
func TestManualRequiredValuesBraces(t *testing.T) {
	want := `@manual{test,
	title = {Test Title},
	year  = {2002}
}`
	have := FormatManualBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{},
		"",
		"",
		"",
		"",
		"",
		true)
	assert.Equal(t, want, have, "should be equal")
}

// Test @manual subcommand output for optional values
func TestManualOptionalValues(t *testing.T) {
	want := `@manual{test,
	title        = "Test Title",
	year         = 2002,
	author       = "Doe, John",
	organisation = "Test Organisation",
	address      = "Test Avenue",
	edition      = "2nd",
	month        = jul,
	note         = "Test Note"
}`
	have := FormatManualBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{"John Doe"},
		"Test Organisation",
		"Test Avenue",
		"2nd",
		"jul",
		"Test Note",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @manual subcommand output for optional values (with braces)
func TestManualOptionalValuesBraces(t *testing.T) {
	want := `@manual{test,
	title        = {Test Title},
	year         = {2002},
	author       = {Doe, John},
	organisation = {Test Organisation},
	address      = {Test Avenue},
	edition      = {2nd},
	month        = {jul},
	note         = {Test Note}
}`
	have := FormatManualBibtex(
		"test",
		"Test Title",
		"2002",
		[]string{"John Doe"},
		"Test Organisation",
		"Test Avenue",
		"2nd",
		"jul",
		"Test Note",
		true)
	assert.Equal(t, want, have, "should be equal")
}

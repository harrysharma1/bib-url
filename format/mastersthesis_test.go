package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @mastersthesis subcommand output for required values only
func TestMastersthesisRequiredValues(t *testing.T) {
	want := `@mastersthesis{test,
	author = "Doe, John",
	title  = "Test Title",
	school = "Test School",
	year   = 2002
}`
	have := FormatMastersThesisBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test School",
		"2002",
		"",
		"",
		"",
		"",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @mastersthesis subcommand output for required values only (with braces)
func TestMastersthesisRequiredValuesBraces(t *testing.T) {
	want := `@mastersthesis{test,
	author = {Doe, John},
	title  = {Test Title},
	school = {Test School},
	year   = {2002}
}`
	have := FormatMastersThesisBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test School",
		"2002",
		"",
		"",
		"",
		"",
		true)
	assert.Equal(t, want, have, "should be equal")
}

// Test @mastersthesis subcommand ouptut for optional values
func TestMastersthesisOptionalValues(t *testing.T) {
	want := `@mastersthesis{test,
	author  = "Doe, John",
	title   = "Test Title",
	school  = "Test School",
	year    = 2002,
	type    = "Test Type",
	address = "Test Avenue",
	month   = jul,
	note    = "Test Note"
}`
	have := FormatMastersThesisBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test School",
		"2002",
		"Test Type",
		"Test Avenue",
		"jul",
		"Test Note",
		false)
	assert.Equal(t, want, have, "should be equal")
}

// Test @mastersthesis subcommand ouptut for optional values (with braces)
func TestMastersthesisOptionalValuesBraces(t *testing.T) {
	want := `@mastersthesis{test,
	author  = {Doe, John},
	title   = {Test Title},
	school  = {Test School},
	year    = {2002},
	type    = {Test Type},
	address = {Test Avenue},
	month   = {jul},
	note    = {Test Note}
}`
	have := FormatMastersThesisBibtex(
		"test",
		[]string{"John Doe"},
		"Test Title",
		"Test School",
		"2002",
		"Test Type",
		"Test Avenue",
		"jul",
		"Test Note",
		true)
	assert.Equal(t, want, have, "should be equal")
}

package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestPhDThesisRequiredValues(t *testing.T) {
	want := `@phdthesis{test,
	author = "Doe, John",
	title  = "Test Title",
	school = "Test School",
	year   = 2002
}`
	have := FormatPhDThesisBibtex(
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

func TestPhDThesisRequiredValuesBraces(t *testing.T) {
	want := `@phdthesis{test,
	author = {Doe, John},
	title  = {Test Title},
	school = {Test School},
	year   = {2002}
}`
	have := FormatPhDThesisBibtex(
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

func TestPhDThesisOptionalValues(t *testing.T) {
	want := `@phdthesis{test,
	author  = "Doe, John",
	title   = "Test Title",
	school  = "Test School",
	year    = 2002,
	type    = "Test Type",
	address = "Test Avenue",
	month   = jul,
	note    = "Test Note"
}`
	have := FormatPhDThesisBibtex(
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

func TestPhDThesisOptionalValuesBraces(t *testing.T) {
	want := `@phdthesis{test,
	author  = {Doe, John},
	title   = {Test Title},
	school  = {Test School},
	year    = {2002},
	type    = {Test Type},
	address = {Test Avenue},
	month   = {jul},
	note    = {Test Note}
}`
	have := FormatPhDThesisBibtex(
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

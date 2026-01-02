package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestBookletRequiredValues(t *testing.T) {
	want := `@booklet{test,
	title        = "Test Title",
	author       = "John Doe",
	howpublished = "Test How Published",
	address      = "Test Avenue",
	year         = 2002
}`
	have := FormatBookletBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		"Test How Published",
		"Test Avenue",
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
func TestBookletRequiredValuesBraces(t *testing.T) {
	want := `@booklet{test,
	title        = {Test Title},
	author       = {John Doe},
	howpublished = {Test How Published},
	address      = {Test Avenue},
	year         = {2002}
}`
	have := FormatBookletBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		"Test How Published",
		"Test Avenue",
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
func TestBookletOptionalValues(t *testing.T) {
	want := `@booklet{test,
	title        = "Test Title",
	author       = "John Doe",
	howpublished = "Test How Published",
	address      = "Test Avenue",
	year         = 2002,
	editor       = "Jane Doe",
	volume       = "50",
	number       = "1",
	series       = "Test Series",
	organisation = "Test Organisation",
	month        = jul,
	note         = "Test Note"
}`
	have := FormatBookletBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		"Test How Published",
		"Test Avenue",
		"2002",
		[]string{"Jane Doe"},
		"50",
		"1",
		"Test Series",
		"Test Organisation",
		"jul",
		"Test Note",
		false)
	assert.Equal(t, want, have, "should be equal")
}
func TestBookletOptionalValuesBraces(t *testing.T) {
	want := `@booklet{test,
	title        = {Test Title},
	author       = {John Doe},
	howpublished = {Test How Published},
	address      = {Test Avenue},
	year         = {2002},
	editor       = {Jane Doe},
	volume       = {50},
	number       = {1},
	series       = {Test Series},
	organisation = {Test Organisation},
	month        = {jul},
	note         = {Test Note}
}`
	have := FormatBookletBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		"Test How Published",
		"Test Avenue",
		"2002",
		[]string{"Jane Doe"},
		"50",
		"1",
		"Test Series",
		"Test Organisation",
		"jul",
		"Test Note",
		true)
	assert.Equal(t, want, have, "should be equal")
}

// --------------------------------------------------//

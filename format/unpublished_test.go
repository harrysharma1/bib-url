package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestUnpublishedRequiredValues(t *testing.T) {
	want := `@unpublished{test,
	author      = "Doe, Jane",
	title       = "Test Title",
	institution = "Test Institution",
	year        = 2002
}`
	have := FormatUnpublishedBibtex(
		"test",
		[]string{"Jane Doe"},
		"Test Title",
		"Test Institution",
		"2002",
		false)
	assert.Equal(t, want, have, "should be equal")
}

func TestUnpublishedRequiredValuesBraces(t *testing.T) {
	want := `@unpublished{test,
	author      = {Doe, Jane},
	title       = {Test Title},
	institution = {Test Institution},
	year        = {2002}
}`

	have := FormatUnpublishedBibtex(
		"test",
		[]string{"Jane Doe"},
		"Test Title",
		"Test Institution",
		"2002",
		true)
	assert.Equal(t, want, have, "should be equal")
}

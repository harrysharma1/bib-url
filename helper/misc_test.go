package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestMiscRequiredValues(t *testing.T) {
	want := `@misc{test,
	title        = "Test Title",
	author       = "Doe, John",
	howpublished = "\url{https://harrysharma.co.uk}",
	year         = 2002,
	note         = "Test Note"
}`

	have := FormatMiscBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		`\url{https://harrysharma.co.uk}`,
		"2002",
		"Test Note",
		false)
	assert.Equal(t, want, have, "should be equal")
}

func TestMiscRequiredValuesBraces(t *testing.T) {
	want := `@misc{test,
	title        = {Test Title},
	author       = {Doe, John},
	howpublished = {\url{https://harrysharma.co.uk}},
	year         = {2002},
	note         = {Test Note}
}`

	have := FormatMiscBibtex(
		"test",
		"Test Title",
		[]string{"John Doe"},
		`\url{https://harrysharma.co.uk}`,
		"2002",
		"Test Note",
		true)
	assert.Equal(t, want, have, "should be equal")
}

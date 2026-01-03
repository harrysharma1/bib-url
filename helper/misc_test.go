package helper

import (
	"testing"
)

// Formatting
// --------------------------------------------------//
func TestMiscExampleValue(t *testing.T) {
	//		want := `@misc{test,
	//		title        = "Test Title",
	//		author       = "John Doe",
	//		howpublished = "\url{https://harrysharma.co.uk}",
	//		year         = 2002,
	//		note         = "Test Note"
	//	}`
	//
	//	have := FormatMiscBibtex(
	//		"test",
	//		[]string{
	//			`title="Test Title"`,
	//			`author="John Doe"`,
	//			`howpublished="https://harrysharma"`,
	//			`year=2002`,
	//			`note="Test Note"`},
	//		false)
	//	assert.Equal(t, want, have, "should be equal")
}

func TestMiscExampleValueBraces(t *testing.T) {
	//		want := `@misc{test,
	//		title        = {Test Title},
	//		author       = {John Doe},
	//		howpublished = {https://harrysharma.co.uk},
	//		year         = {2002},
	//		note         = {Test Note}"
	//	}`
	//
	//	have := FormatMiscBibtex(
	//		"test",
	//		[]string{
	//			`title="Test Title"`,
	//			`author="John Doe"`,
	//			`howpublished="\url{https://harrysharma}"`,
	//			`year=2002`,
	//			`note="Test Note"`},
	//		true)
	//	assert.Equal(t, want, have, "should be equal")
}

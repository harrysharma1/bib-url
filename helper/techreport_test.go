package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestTechReportRequiredValues(t *testing.T) {
	want := `@techreport{test,
	title       = "Test Title",
	author      = "Doe, Jane",
	institution = "Test Institution",
	address     = "Test Address",
	number      = "50",
	year        = 2002,
	month       = jul
}`
	have := FormatTechReportBibtex(
		"test",
		"Test Title",
		[]string{"Jane Doe"},
		"Test Institution",
		"Test Address",
		"50",
		"2002",
		"jul",
		false)
	assert.Equal(t, want, have, "should be equal")
}

func TestTechReportRequiredValuesBraces(t *testing.T) {
	want := `@techreport{test,
	title       = {Test Title},
	author      = {Doe, Jane},
	institution = {Test Institution},
	address     = {Test Address},
	number      = {50},
	year        = {2002},
	month       = {jul}
}`
	have := FormatTechReportBibtex(
		"test",
		"Test Title",
		[]string{"Jane Doe"},
		"Test Institution",
		"Test Address",
		"50",
		"2002",
		"jul",
		true)
	assert.Equal(t, want, have, "should be equal")
}

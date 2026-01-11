package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

// Test @techreport subcommand output for required values only
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

// Test @techreport subcommand output for required values only (with braces)
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

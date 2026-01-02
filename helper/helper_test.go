package helper

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	emptyValResult := defaultIfEmpty("", "Test Default")
	fillValResult := defaultIfEmpty("Test Value", "Test Default")
	assert.Equal(t, "Test Default", emptyValResult, "should be default value")
	assert.Equal(t, "Test Value", fillValResult, "should be actual value")
}

func TestParseFields(t *testing.T) {
	fields := []string{`author="John Doe"`}
	haveMap := parseFields(fields)
	assert.Equal(t, "\"John Doe\"", strings.Join(haveMap["author"], " and "), "fields should parse by splitting k:v by =")
}

func TestLargestFieldInt(t *testing.T) {
	testS := []string{"Klein", "Moretti"}
	have := largestFieldInt(testS)
	want := 7
	assert.Equal(t, want, have, "should pick based on largest string length")
}

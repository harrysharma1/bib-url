package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	emptyValResult := defaultIfEmpty("", "Test Default")
	fillValResult := defaultIfEmpty("Test Value", "Test Default")
	assert.Equal(t, "Test Default", emptyValResult, "should be default value")
	assert.Equal(t, "Test Value", fillValResult, "should be actual value")
}

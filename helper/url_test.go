package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidUrl(t *testing.T) {
	invalidUrl := "abc"
	validUrl := "https://harrysharma.co.uk"

	assert.True(t, IsValidUrl(validUrl), "should be valid")
	assert.False(t, IsValidUrl(invalidUrl), "should be invalid")
}

package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing if isValidUrl() fails and passes on correct and invalid URLs
//
// For later use cases e.g. generalised webscraping
func TestIsValidUrl(t *testing.T) {
	invalidUrl := "abc"
	validUrl := "https://harrysharma.co.uk"

	assert.True(t, IsValidUrl(validUrl), "should be valid")
	assert.False(t, IsValidUrl(invalidUrl), "should be invalid")
}

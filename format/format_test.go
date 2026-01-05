package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	def := "DEFAULT VALUE"
	val := "ACTUAL VALUE"
	wantEmpty := def
	wantFull := val
	haveEmpty := defaultIfEmpty("", def)
	haveFull := defaultIfEmpty(val, def)

	assert.Equal(t, wantEmpty, haveEmpty, "function should return default value")
	assert.Equal(t, wantFull, haveFull, "function should return provided value")
}

func TestWrap(t *testing.T) {
	wantBraces := "{test}"
	wantMonth := "jul"
	wantMonthBraces := "{jul}"
	wantSpeech := `"test"`
	haveBraces := wrap("test", true)
	haveMonth := wrap("jul", false)
	haveMonthBraces := wrap("jul", true)
	haveSpeech := wrap("test", false)
	assert.Equal(t, wantBraces, haveBraces, "string should be returned with braces")
	assert.Equal(t, wantMonth, haveMonth, "string for month returns without wrapping")
	assert.Equal(t, wantMonthBraces, haveMonthBraces, "braces should wrap the month")
	assert.Equal(t, wantSpeech, haveSpeech, "quotation marks should be wrap string")

}

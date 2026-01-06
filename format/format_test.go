package format

import (
	"fmt"
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

func TestGetMaxKeyLength(t *testing.T) {
	sameSizeStrSlice := []Field{{"a", ""}, {"b", ""}, {"c", ""}}
	definitiveMaxSizeStrSlice := []Field{{"ab", ""}, {"c", ""}, {"d", ""}}
	wantSameSize := 1
	haveSameSize := getMaxKeyLength(sameSizeStrSlice)
	wantMaxSize := 2
	haveMaxSize := getMaxKeyLength(definitiveMaxSizeStrSlice)
	assert.Equal(t, wantSameSize, haveSameSize, "all keys equal in length")
	assert.Equal(t, wantMaxSize, haveMaxSize, "the key size should be 2")
}

func TestFormatFields(t *testing.T) {
	want := []string{
		fmt.Sprintf("\tauthor  = %s", wrap(toBibtexName("John Doe"), false)),
		fmt.Sprintf("\ttitle   = %s", wrap("Test Title", false)),
		fmt.Sprintf("\tjournal = %s", wrap("Test Journal", false)),
		fmt.Sprintf("\tyear    = %d", 2002),
		fmt.Sprintf("\tvolume  = %s", wrap("50", false)),
		fmt.Sprintf("\tnumber  = %s", wrap("1", false)),
		fmt.Sprintf("\tpages   = %s", wrap("1--10", false)),
		fmt.Sprintf("\tmonth   = %s", wrap("jul", false)),
		fmt.Sprintf("\tnote    = %s", wrap("Test Note", false))}
	fields := []Field{
		{"author", toBibtexName("John Doe")},
		{"title", "Test Title"},
		{"journal", "Test Journal"},
		{"year", "2002"},
		{"volume", "50"},
		{"number", "1"},
		{"pages", "1--10"},
		{"month", "jul"},
		{"note", "Test Note"},
	}
	have := formatFields(fields, false)
	assert.Equal(t, want, have, "formatted fields different")
}

func TestFormatFieldsBraces(t *testing.T) {
	want := []string{
		fmt.Sprintf("\tauthor  = %s", wrap(toBibtexName("John Doe"), true)),
		fmt.Sprintf("\ttitle   = %s", wrap("Test Title", true)),
		fmt.Sprintf("\tjournal = %s", wrap("Test Journal", true)),
		fmt.Sprintf("\tyear    = %s", wrap("2002", true)),
		fmt.Sprintf("\tvolume  = %s", wrap("50", true)),
		fmt.Sprintf("\tnumber  = %s", wrap("1", true)),
		fmt.Sprintf("\tpages   = %s", wrap("1--10", true)),
		fmt.Sprintf("\tmonth   = %s", wrap("jul", true)),
		fmt.Sprintf("\tnote    = %s", wrap("Test Note", true))}
	fields := []Field{
		{"author", toBibtexName("John Doe")},
		{"title", "Test Title"},
		{"journal", "Test Journal"},
		{"year", "2002"},
		{"volume", "50"},
		{"number", "1"},
		{"pages", "1--10"},
		{"month", "jul"},
		{"note", "Test Note"},
	}
	have := formatFields(fields, true)
	assert.Equal(t, want, have, "formatted fields different")
}

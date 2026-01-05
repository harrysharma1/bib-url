package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

func TestArticleRequiredValues(t *testing.T) {
	want := `@article{test,
	author  = "Doe, John",
	title   = "Test Title",
	journal = "Test Journal",
	year    = 2002
}`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "", "", "", "", "", false)
	assert.Equal(t, want, have, "should be equal")
}

func TestArticleRequiredValuesBraces(t *testing.T) {
	want := `@article{test,
	author  = {Doe, John},
	title   = {Test Title},
	journal = {Test Journal},
	year    = {2002}
}`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "", "", "", "", "", true)
	assert.Equal(t, want, have, "should be equal")
}

func TestArticleOptionalValues(t *testing.T) {
	want := `@article{test,
	author  = "Doe, John",
	title   = "Test Title",
	journal = "Test Journal",
	year    = 2002,
	volume  = "50",
	number  = "1",
	pages   = "1--10",
	month   = jul,
	note    = "Test Note"
}`

	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "50", "1", "1--10", "jul", "Test Note", false)
	assert.Equal(t, want, have, "should be equal")
}
func TestArticleOptionalValuesBraces(t *testing.T) {
	want := `@article{test,
	author  = {Doe, John},
	title   = {Test Title},
	journal = {Test Journal},
	year    = {2002},
	volume  = {50},
	number  = {1},
	pages   = {1--10},
	month   = {jul},
	note    = {Test Note}
}`

	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "50", "1", "1--10", "jul", "Test Note", true)
	assert.Equal(t, want, have, "should be equal")
}

//--------------------------------------------------//

// DOI
// --------------------------------------------------//
func TestArticleDOIValid(t *testing.T) {
	doi := "10.3390/electronics10212707"
	wantAuthors := []string{"Fatiha Djebbar"}
	wantTitle := "Securing IoT Data Using Steganography: A Practical Implementation Approach"
	wantJournal := "Electronics"
	wantYear := "2021"
	wantVolume := "10"
	wantNumber := "21"

	haveAuthors, haveTitle, haveJournal, haveYear, haveVolume, haveNumber, err := ArticleFromDOI(doi)

	assert.Nil(t, err, "error occurred in ArticleFromDOI()")

	// authors
	assert.Equal(t, len(wantAuthors), len(haveAuthors), "not same length")
	assert.Equal(t, wantAuthors, haveAuthors, "not in same order")

	// title
	assert.Equal(t, wantTitle, haveTitle, "titles not equal")
	// journal

	assert.Equal(t, wantJournal, haveJournal, "journals not equal")
	// year
	assert.Equal(t, wantYear, haveYear, "years are not equal")
	// volume
	assert.Equal(t, wantVolume, haveVolume, "volumes not same")
	// number
	assert.Equal(t, wantNumber, haveNumber, "numbers not the same")
}

func TestArticleDOIInvalid(t *testing.T) {
	doi := "cafbawiefkhapoucwoefhewkjfnfs"
	_, _, _, _, _, _, err := ArticleFromDOI(doi)
	if err != nil {
		if err.Error() != "work identified does not exist" {
			t.Error("returned error did not indicate that the article does not exist")
		}
	}
}

//--------------------------------------------------//

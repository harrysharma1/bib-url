package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//

func TestArticleRequiredValues(t *testing.T) {
	want := `@article{test,
	author  = "John Doe",
	title   = "Test Title",
	journal = "Test Journal",
	year    = 2002
}`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "", "", "", "", "", false)
	assert.Equal(t, want, have, "should be equal")
}

func TestArticleRequiredValuesBraces(t *testing.T) {
	want := `@article{test,
	author  = {John Doe},
	title   = {Test Title},
	journal = {Test Journal},
	year    = {2002}
}`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "Test Title", "Test Journal", "2002", "", "", "", "", "", true)
	assert.Equal(t, want, have, "should be equal")
}

func TestArticleOptionalValues(t *testing.T) {
	want := `@article{test,
	author  = "John Doe",
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
	author  = {John Doe},
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

	if err != nil {
		t.Errorf("error ocurred in ArticleFromDOI() function call: %e", err)
	}

	// authors
	if len(wantAuthors) == len(haveAuthors) {
		for i := range len(wantAuthors) {
			if wantAuthors[i] != haveAuthors[i] {
				t.Errorf("return author at position %d not the same:\nWant %s\nHave:%s", i, wantAuthors[i], haveAuthors[i])
			}
		}
	} else {
		t.Errorf("return authors length does not match, either change test to reflect this or inspect ArticleFromDOI() function for bugs:\nDefault Author Count: %d\nReturn Author Count: %d", len(wantAuthors), len(haveAuthors))
	}

	// title
	if wantTitle != haveTitle {
		t.Errorf("return title does not match default title:\nWant %s\nHave: %s", wantTitle, haveTitle)
	}

	// journal
	if wantJournal != haveJournal {
		t.Errorf("return journal does not match default journal:\nWant %s\nHave: %s", wantJournal, haveJournal)
	}

	// year
	if wantYear != haveYear {
		t.Errorf("return year does not match default year:\nWant %s\nHave: %s", wantYear, haveYear)
	}

	// volume
	if wantVolume != haveVolume {
		t.Errorf("return volume does not match default volume:\nWant %s\nHave: %s", wantVolume, haveVolume)
	}

	// number
	if wantNumber != haveNumber {
		t.Errorf("return volume does not match default volume:\nWant %s\nHave: %s", wantNumber, haveNumber)
	}

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

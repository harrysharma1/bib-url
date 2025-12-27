package helper

import (
	"strings"
	"testing"
)

// Formatting
// --------------------------------------------------//
// return value
func TestArticleFormatDefaultNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatDefaultBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Custom Values

// Author
func TestArticleFormatCustomAuthorNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "John Doe",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomAuthorBraces(t *testing.T) {
	want := `@article{test,
	author   = {John Doe},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"John Doe"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomAuthorsNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "John Doe and Jane Doe",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"John Doe", "Jane Doe"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomAuthorsBraces(t *testing.T) {
	want := `@article{test,
	author   = {John Doe and Jane Doe},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"John Doe", "Jane Doe"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Title
func TestArticleFormatCustomTitleNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "The independence of the continuum hypothesis",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "The independence of the continuum hypothesis", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomTitleBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {The independence of the continuum hypothesis},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "The independence of the continuum hypothesis", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Journal
func TestArticleFormatCustomJournalNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "Proceedings of the National Academy of Sciences",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "Proceedings of the National Academy of Sciences", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomJournalBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {Proceedings of the National Academy of Sciences},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "Proceedings of the National Academy of Sciences", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Year
func TestArticleFormatCustomYearNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = 1963,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "1963", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomYearBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {1963},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "1963", "<1: Please Change>", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Volume
func TestArticleFormatCustomVolumeNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "50",
	number   = "<1: Please Change>",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "50", "<1: Please Change>", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomVolumeBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {50},
	number   = {<1: Please Change>},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "50", "<1: Please Change>", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Number
func TestArticleFormatCustomNumberNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "6",
	pages    = "<1--10: Please Change>"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "6", "<1--10: Please Change>", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomNumberBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {6},
	pages    = {<1--10: Please Change>}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "6", "<1--10: Please Change>", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Pages
func TestArticleFormatCustomPagesNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "<Lastname, FirstName>",
	title    = "<Example Title: Please Change>",
	journal  = "<Example Journal: Please Change>",
	year     = <2002: Please Change>,
	volume   = "<1: Please Change>",
	number   = "<1: Please Change>",
	pages    = "1143--1148"
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "1143--1148", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomPagesBraces(t *testing.T) {
	want := `@article{test,
	author   = {<Lastname, FirstName>},
	title    = {<Example Title: Please Change>},
	journal  = {<Example Journal: Please Change>},
	year     = {<2002: Please Change>},
	volume   = {<1: Please Change>},
	number   = {<1: Please Change>},
	pages    = {1143--1148}
}
`
	have := FormatArticleBibtex("test", []string{"<Lastname, FirstName>"}, "<Example Title: Please Change>", "<Example Journal: Please Change>", "<2002: Please Change>", "<1: Please Change>", "<1: Please Change>", "1143--1148", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

// Fully custom values
func TestArticleFormatCustomNoBraces(t *testing.T) {
	want := `@article{test,
	author   = "P. J. Cohen",
	title    = "The independence of the continuum hypothesis",
	journal  = "Proceedings of the National Academy of Sciences",
	year     = 1963,
	volume   = "50",
	number   = "6",
	pages    = "1143--1148"
}
`
	have := FormatArticleBibtex("test", []string{"P. J. Cohen"}, "The independence of the continuum hypothesis", "Proceedings of the National Academy of Sciences", "1963", "50", "6", "1143--1148", false)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
}

func TestArticleFormatCustomBraces(t *testing.T) {
	want := `@article{test,
	author   = {P. J. Cohen},
	title    = {The independence of the continuum hypothesis},
	journal  = {Proceedings of the National Academy of Sciences},
	year     = {1963},
	volume   = {50},
	number   = {6},
	pages    = {1143--1148}
}
`
	have := FormatArticleBibtex("test", []string{"P. J. Cohen"}, "The independence of the continuum hypothesis", "Proceedings of the National Academy of Sciences", "1963", "50", "6", "1143--1148", true)
	if strings.TrimSpace(want) != strings.TrimSpace(have) {
		t.Errorf("return value incorrect, either change test to reflect this change or fix article formatting:\nWant\n%s\nHave:\n%s", want, have)
	}
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

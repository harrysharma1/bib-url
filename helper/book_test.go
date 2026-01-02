package helper

import "testing"

// Formatting
// --------------------------------------------------//
func TestBookRequiredValues(t *testing.T) {
	want := `@book{test,
	author    = "Frank Herbert",
	title     = "Dune",
	publisher = "Chilton Books",
	address   = "Sparkford, Yeovil, Somerset",
	year      = 1965
}`
	have := FormatBookBibtex("test", []string{"Frank Herbert"}, "Dune", "Chilton Books", "Sparkford, Yeovil, Somerset", "1965", false)
	if want != have {
		t.Errorf("formatting error:\nWant:\n%s\nHave:\n%s", want, have)
	}

}

func TestBookRequiredValuesBraces(t *testing.T) {
	want := `@book{test,
	author    = {Frank Herbert},
	title     = {Dune},
	publisher = {Chilton Books},
	address   = {Sparkford, Yeovil, Somerset},
	year      = {1965}
}`
	have := FormatBookBibtex("test", []string{"Frank Herbert"}, "Dune", "Chilton Books", "Sparkford, Yeovil, Somerset", "1965", true)
	if want != have {
		t.Errorf("formatting error:\nWant:\n%s\nHave:\n%s", want, have)
	}

}

// --------------------------------------------------//

// ISBN
// --------------------------------------------------//
func TestBookISBNValid(t *testing.T) {
	isbn := "0340960191"
	wantAuthors := []string{"Frank Herbert"}
	wantTitle := "Dune"
	wantPublisher := "Hodder Paperback"
	wantYear := "2015"
	haveAuthors, haveTitle, havePublisher, haveYear, err := BookFromISBN(isbn)
	if err != nil {
		t.Errorf("error ocurred in BookFromISBN() function call: %e", err)
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

	// publisher
	if wantPublisher != havePublisher {
		t.Errorf("return publisher does not match with default publisher:\nWant: %s\nHave: %s", wantPublisher, havePublisher)
	}

	// year
	if wantYear != haveYear {
		t.Errorf("return year does not match default year:\nWant: %s\nHave: %s", wantYear, haveYear)
	}
}

func TestBookISBNInvalid(t *testing.T) {
	isbn := "dueafbwieufhajcliheyfiuqss"
	_, _, _, _, err := BookFromISBN(isbn)
	if err != nil {
		if err.Error() != "book not found" {
			t.Error("returned error did not indicated that the book does not exist")
		}
	}
}

// --------------------------------------------------//

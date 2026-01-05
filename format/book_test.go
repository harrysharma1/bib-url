package format

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Formatting
// --------------------------------------------------//
func TestBookRequiredValues(t *testing.T) {
	want := `@book{test,
	author    = "Herbert, Frank",
	title     = "Dune",
	publisher = "Chilton Books",
	address   = "Sparkford, Yeovil, Somerset",
	year      = 1965
}`
	have := FormatBookBibtex("test", []string{"Frank Herbert"}, "Dune", "Chilton Books", "Sparkford, Yeovil, Somerset", "1965", false)
	assert.Equal(t, want, have, "should be equal")

}

func TestBookRequiredValuesBraces(t *testing.T) {
	want := `@book{test,
	author    = {Herbert, Frank},
	title     = {Dune},
	publisher = {Chilton Books},
	address   = {Sparkford, Yeovil, Somerset},
	year      = {1965}
}`
	have := FormatBookBibtex("test", []string{"Frank Herbert"}, "Dune", "Chilton Books", "Sparkford, Yeovil, Somerset", "1965", true)
	assert.Equal(t, want, have, "should be equal")

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
	assert.Nil(t, err, "BookFromISBN(isbn) failed")

	// authors
	assert.Equal(t, len(wantAuthors), len(haveAuthors), "author lengths do not match")
	if len(wantAuthors) == len(haveAuthors) {
		for i := range len(wantAuthors) {
			assert.Equal(t, wantAuthors[i], haveAuthors[i], fmt.Sprintf("author at position %d is not the same", i))
		}
	}

	// title
	assert.Equal(t, wantTitle, haveTitle, "title not the same")

	// publisher
	assert.Equal(t, wantPublisher, havePublisher, "publisher not the same")

	// year
	assert.Equal(t, wantYear, haveYear, "year not the same")
}

func TestBookISBNInvalid(t *testing.T) {
	isbn := "dueafbwieufhajcliheyfiuqss"
	_, _, _, _, err := BookFromISBN(isbn)
	assert.EqualError(t, err, "book not found", "incorrect error")
}

// --------------------------------------------------//

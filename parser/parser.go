package parser

import "strings"

func Parse(bibtex string) {
	strings.SplitN(bibtex, "{", 1)
}

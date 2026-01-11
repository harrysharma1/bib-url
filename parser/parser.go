package parser

import "strings"

// Parses bibtex from file
//
// TODO: for later conversions of bibtex files into APA/Chicago/Harvard styles
func Parse(bibtex string) {
	strings.SplitN(bibtex, "{", 1)
}

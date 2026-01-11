package models

// Struct of @techreport
type Techreport struct {
	CiteKey     string
	Title       string
	Authors     []string
	Institution string
	Address     string
	Number      string
	Year        string
	Month       string
}

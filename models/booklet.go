package models

type Booklet struct {
	CiteKey      string
	Title        string
	Authors      []string
	HowPublished string
	Address      string
	Year         string
	Editors      []string
	Volume       string
	Number       string
	Series       string
	Organisation string
	Month        string
	Note         string
}

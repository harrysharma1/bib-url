package models

type Article struct {
	Doi     string
	CiteKey string
	Authors []string
	Title   string
	Journal string
	Year    string
	Volume  string
	Number  string
	Pages   string
	Month   string
	Note    string
}

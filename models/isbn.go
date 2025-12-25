package models

type BookInfo struct {
	Publishers    []string `json:"publishers"`
	Subtitle      string   `json:"subtitle"`
	SourceRecords []string `json:"source_records"`
	Title         string   `json:"title"`
	NumberOfPages int      `json:"number_of_pages"`
	Isbn13        []string `json:"isbn_13"`
	Languages     []struct {
		Key string `json:"key"`
	} `json:"languages"`
	FullTitle         string   `json:"full_title"`
	LcClassifications []string `json:"lc_classifications"`
	PublishDate       string   `json:"publish_date"`
	Key               string   `json:"key"`
	Authors           []struct {
		Key string `json:"key"`
	} `json:"authors"`
	Works []struct {
		Key string `json:"key"`
	} `json:"works"`
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Subjects       []string `json:"subjects"`
	Covers         []int    `json:"covers"`
	Ocaid          string   `json:"ocaid"`
	OclcNumbers    []string `json:"oclc_numbers"`
	LatestRevision int      `json:"latest_revision"`
	Revision       int      `json:"revision"`
	Created        struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}

type Author struct {
	BirthDate string `json:"birth_date"`
	Type      struct {
		Key string `json:"key"`
	} `json:"type"`
	Key           string   `json:"key"`
	SourceRecords []string `json:"source_records"`
	Photos        []int    `json:"photos"`
	Name          string   `json:"name"`
	RemoteIds     struct {
		Wikidata string `json:"wikidata"`
		Isni     string `json:"isni"`
		Viaf     string `json:"viaf"`
	} `json:"remote_ids"`
	PersonalName   string `json:"personal_name"`
	Bio            string `json:"bio"`
	LatestRevision int    `json:"latest_revision"`
	Revision       int    `json:"revision"`
	Created        struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
}

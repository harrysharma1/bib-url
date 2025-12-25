package models

import "time"

// Generated from: https://transform.tools/json-to-go
type Works struct {
	Status         string `json:"status"`
	MessageType    string `json:"message-type"`
	MessageVersion string `json:"message-version"`
	Message        struct {
		Institution []struct {
			Name       string   `json:"name"`
			Place      []string `json:"place"`
			Department []string `json:"department"`
			Acronym    []string `json:"acronym"`
			ID         []struct {
				ID         string `json:"id"`
				IDType     string `json:"id-type"`
				AssertedBy string `json:"asserted-by"`
			} `json:"id"`
		} `json:"institution"`
		Indexed struct {
			DateParts [][]int `json:"date-parts"`
			Version   string  `json:"version"`
		} `json:"indexed"`
		Description string `json:"description"`
		Posted      struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"posted"`
		PublisherLocation string `json:"publisher-location"`
		UpdateTo          []struct {
			Label   string `json:"label"`
			DOI     string `json:"DOI"`
			Type    string `json:"type"`
			Source  string `json:"source"`
			Updated struct {
				DateParts [][]int   `json:"date-parts"`
				DateTime  time.Time `json:"date-time"`
				Timestamp int       `json:"timestamp"`
			} `json:"updated"`
			RecordID string `json:"record-id"`
		} `json:"update-to"`
		StandardsBody struct {
			Name    string `json:"name"`
			Acronym string `json:"acronym"`
		} `json:"standards-body"`
		EditionNumber  string `json:"edition-number"`
		GroupTitle     string `json:"group-title"`
		ReferenceCount int    `json:"reference-count"`
		Publisher      string `json:"publisher"`
		Issue          string `json:"issue"`
		IsbnType       []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"isbn-type"`
		License []struct {
			URL   string `json:"URL"`
			Start struct {
				DateParts [][]int   `json:"date-parts"`
				DateTime  time.Time `json:"date-time"`
				Timestamp int       `json:"timestamp"`
			} `json:"start"`
			DelayInDays    int    `json:"delay-in-days"`
			ContentVersion string `json:"content-version"`
		} `json:"license"`
		Funder []struct {
			Name          string   `json:"name"`
			DOI           string   `json:"DOI"`
			DoiAssertedBy string   `json:"doi-asserted-by"`
			Award         []string `json:"award"`
			ID            []struct {
				ID         string `json:"id"`
				IDType     string `json:"id-type"`
				AssertedBy string `json:"asserted-by"`
			} `json:"id"`
		} `json:"funder"`
		ContentDomain struct {
			Domain               []string `json:"domain"`
			CrossmarkRestriction bool     `json:"crossmark-restriction"`
		} `json:"content-domain"`
		Chair []struct {
			ORCID       string `json:"ORCID"`
			Suffix      string `json:"suffix"`
			Given       string `json:"given"`
			Family      string `json:"family"`
			Affiliation []struct {
				Name       string   `json:"name"`
				Place      []string `json:"place"`
				Department []string `json:"department"`
				Acronym    []string `json:"acronym"`
				ID         []struct {
					ID         string `json:"id"`
					IDType     string `json:"id-type"`
					AssertedBy string `json:"asserted-by"`
				} `json:"id"`
			} `json:"affiliation"`
			Name               string `json:"name"`
			AuthenticatedOrcid bool   `json:"authenticated-orcid"`
			Prefix             string `json:"prefix"`
			Sequence           string `json:"sequence"`
		} `json:"chair"`
		ShortContainerTitle []string `json:"short-container-title"`
		Accepted            struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"accepted"`
		SpecialNumbering string `json:"special-numbering"`
		ContentUpdated   struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"content-updated"`
		PublishedPrint struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"published-print"`
		Abstract string `json:"abstract"`
		DOI      string `json:"DOI"`
		Type     string `json:"type"`
		Created  struct {
			DateParts [][]int   `json:"date-parts"`
			DateTime  time.Time `json:"date-time"`
			Timestamp int       `json:"timestamp"`
		} `json:"created"`
		Approved struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"approved"`
		Page                string   `json:"page"`
		UpdatePolicy        string   `json:"update-policy"`
		Source              string   `json:"source"`
		IsReferencedByCount int      `json:"is-referenced-by-count"`
		Title               []string `json:"title"`
		Prefix              string   `json:"prefix"`
		Volume              string   `json:"volume"`
		ClinicalTrialNumber []struct {
			ClinicalTrialNumber string `json:"clinical-trial-number"`
			Registry            string `json:"registry"`
			Type                string `json:"type"`
		} `json:"clinical-trial-number"`
		Author []struct {
			ORCID       string `json:"ORCID"`
			Suffix      string `json:"suffix"`
			Given       string `json:"given"`
			Family      string `json:"family"`
			Affiliation []struct {
				Name       string   `json:"name"`
				Place      []string `json:"place"`
				Department []string `json:"department"`
				Acronym    []string `json:"acronym"`
				ID         []struct {
					ID         string `json:"id"`
					IDType     string `json:"id-type"`
					AssertedBy string `json:"asserted-by"`
				} `json:"id"`
			} `json:"affiliation"`
			Name               string `json:"name"`
			AuthenticatedOrcid bool   `json:"authenticated-orcid"`
			Prefix             string `json:"prefix"`
			Sequence           string `json:"sequence"`
		} `json:"author"`
		Member         string `json:"member"`
		ContentCreated struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"content-created"`
		PublishedOnline struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"published-online"`
		Reference []struct {
			Issn               string `json:"issn"`
			StandardsBody      string `json:"standards-body"`
			Issue              string `json:"issue"`
			Key                string `json:"key"`
			SeriesTitle        string `json:"series-title"`
			IsbnType           string `json:"isbn-type"`
			DoiAssertedBy      string `json:"doi-asserted-by"`
			FirstPage          string `json:"first-page"`
			DOI                string `json:"DOI"`
			Type               string `json:"type"`
			Isbn               string `json:"isbn"`
			Component          string `json:"component"`
			ArticleTitle       string `json:"article-title"`
			VolumeTitle        string `json:"volume-title"`
			Volume             string `json:"volume"`
			Author             string `json:"author"`
			StandardDesignator string `json:"standard-designator"`
			Year               string `json:"year"`
			Unstructured       string `json:"unstructured"`
			Edition            string `json:"edition"`
			JournalTitle       string `json:"journal-title"`
			IssnType           string `json:"issn-type"`
		} `json:"reference"`
		UpdatedBy []struct {
			Label   string `json:"label"`
			DOI     string `json:"DOI"`
			Type    string `json:"type"`
			Source  string `json:"source"`
			Updated struct {
				DateParts [][]int   `json:"date-parts"`
				DateTime  time.Time `json:"date-time"`
				Timestamp int       `json:"timestamp"`
			} `json:"updated"`
			RecordID string `json:"record-id"`
		} `json:"updated-by"`
		Event struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Start    struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"start"`
			End struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"end"`
		} `json:"event"`
		ContainerTitle []string `json:"container-title"`
		Review         struct {
			Type                       string `json:"type"`
			RunningNumber              string `json:"running-number"`
			RevisionRound              string `json:"revision-round"`
			Stage                      string `json:"stage"`
			CompetingInterestStatement string `json:"competing-interest-statement"`
			Recommendation             string `json:"recommendation"`
			Language                   string `json:"language"`
		} `json:"review"`
		Project []struct {
			AwardEnd []struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"award-end"`
			AwardPlannedStart []struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"award-planned-start"`
			AwardStart []struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"award-start"`
			LeadInvestigator []struct {
				ORCID       string `json:"ORCID"`
				Suffix      string `json:"suffix"`
				Given       string `json:"given"`
				Family      string `json:"family"`
				Affiliation []struct {
					ID []struct {
						ID         string `json:"id"`
						IDType     string `json:"id-type"`
						AssertedBy string `json:"asserted-by"`
					} `json:"id"`
					Name string `json:"name"`
				} `json:"affiliation"`
				Name      string `json:"name"`
				RoleStart struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-start"`
				AuthenticatedOrcid bool   `json:"authenticated-orcid"`
				Prefix             string `json:"prefix"`
				AlternateName      string `json:"alternate-name"`
				Sequence           string `json:"sequence"`
				RoleEnd            struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-end"`
			} `json:"lead-investigator"`
			AwardPlannedEnd []struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"award-planned-end"`
			Investigator []struct {
				ORCID       string `json:"ORCID"`
				Suffix      string `json:"suffix"`
				Given       string `json:"given"`
				Family      string `json:"family"`
				Affiliation []struct {
					ID []struct {
						ID         string `json:"id"`
						IDType     string `json:"id-type"`
						AssertedBy string `json:"asserted-by"`
					} `json:"id"`
					Name string `json:"name"`
				} `json:"affiliation"`
				Name      string `json:"name"`
				RoleStart struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-start"`
				AuthenticatedOrcid bool   `json:"authenticated-orcid"`
				Prefix             string `json:"prefix"`
				AlternateName      string `json:"alternate-name"`
				Sequence           string `json:"sequence"`
				RoleEnd            struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-end"`
			} `json:"investigator"`
			Funding []struct {
				Type        string `json:"type"`
				Scheme      string `json:"scheme"`
				AwardAmount struct {
					Amount     int    `json:"amount"`
					Currency   string `json:"currency"`
					Percentage int    `json:"percentage"`
				} `json:"award-amount"`
				Funder struct {
					Name          string   `json:"name"`
					DOI           string   `json:"DOI"`
					DoiAssertedBy string   `json:"doi-asserted-by"`
					Award         []string `json:"award"`
					ID            []struct {
						ID         string `json:"id"`
						IDType     string `json:"id-type"`
						AssertedBy string `json:"asserted-by"`
					} `json:"id"`
				} `json:"funder"`
			} `json:"funding"`
			ProjectTitle []struct {
				Title    string `json:"title"`
				Language string `json:"language"`
			} `json:"project-title"`
			AwardAmount struct {
				Amount     int    `json:"amount"`
				Currency   string `json:"currency"`
				Percentage int    `json:"percentage"`
			} `json:"award-amount"`
			CoLeadInvestigator []struct {
				ORCID       string `json:"ORCID"`
				Suffix      string `json:"suffix"`
				Given       string `json:"given"`
				Family      string `json:"family"`
				Affiliation []struct {
					ID []struct {
						ID         string `json:"id"`
						IDType     string `json:"id-type"`
						AssertedBy string `json:"asserted-by"`
					} `json:"id"`
					Name string `json:"name"`
				} `json:"affiliation"`
				Name      string `json:"name"`
				RoleStart struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-start"`
				AuthenticatedOrcid bool   `json:"authenticated-orcid"`
				Prefix             string `json:"prefix"`
				AlternateName      string `json:"alternate-name"`
				Sequence           string `json:"sequence"`
				RoleEnd            struct {
					DateParts [][]string `json:"date-parts"`
				} `json:"role-end"`
			} `json:"co-lead-investigator"`
			ProjectDescription []struct {
				Description string `json:"description"`
				Language    string `json:"language"`
			} `json:"project-description"`
		} `json:"project"`
		OriginalTitle []string `json:"original-title"`
		Status        struct {
			Type   string `json:"type"`
			Update struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"update"`
			StatusDescription []struct {
				Language    string `json:"language"`
				Description string `json:"description"`
			} `json:"status-description"`
		} `json:"status"`
		Language string `json:"language"`
		Link     []struct {
			URL                 string `json:"URL"`
			ContentType         string `json:"content-type"`
			ContentVersion      string `json:"content-version"`
			IntendedApplication string `json:"intended-application"`
		} `json:"link"`
		Deposited struct {
			DateParts [][]int   `json:"date-parts"`
			DateTime  time.Time `json:"date-time"`
			Timestamp int       `json:"timestamp"`
		} `json:"deposited"`
		Score    int      `json:"score"`
		Degree   []string `json:"degree"`
		Resource struct {
			Primary struct {
				URL string `json:"URL"`
			} `json:"primary"`
			Secondary []struct {
				URL   string `json:"URL"`
				Label string `json:"label"`
			} `json:"secondary"`
		} `json:"resource"`
		Subtitle   []string `json:"subtitle"`
		Translator []struct {
			ORCID       string `json:"ORCID"`
			Suffix      string `json:"suffix"`
			Given       string `json:"given"`
			Family      string `json:"family"`
			Affiliation []struct {
				Name       string   `json:"name"`
				Place      []string `json:"place"`
				Department []string `json:"department"`
				Acronym    []string `json:"acronym"`
				ID         []struct {
					ID         string `json:"id"`
					IDType     string `json:"id-type"`
					AssertedBy string `json:"asserted-by"`
				} `json:"id"`
			} `json:"affiliation"`
			Name               string `json:"name"`
			AuthenticatedOrcid bool   `json:"authenticated-orcid"`
			Prefix             string `json:"prefix"`
			Sequence           string `json:"sequence"`
		} `json:"translator"`
		FreeToRead struct {
			StartDate struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"start-date"`
			EndDate struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"end-date"`
		} `json:"free-to-read"`
		Editor []struct {
			ORCID       string `json:"ORCID"`
			Suffix      string `json:"suffix"`
			Given       string `json:"given"`
			Family      string `json:"family"`
			Affiliation []struct {
				Name       string   `json:"name"`
				Place      []string `json:"place"`
				Department []string `json:"department"`
				Acronym    []string `json:"acronym"`
				ID         []struct {
					ID         string `json:"id"`
					IDType     string `json:"id-type"`
					AssertedBy string `json:"asserted-by"`
				} `json:"id"`
			} `json:"affiliation"`
			Name               string `json:"name"`
			AuthenticatedOrcid bool   `json:"authenticated-orcid"`
			Prefix             string `json:"prefix"`
			Sequence           string `json:"sequence"`
		} `json:"editor"`
		ProceedingsSubject string   `json:"proceedings-subject"`
		ComponentNumber    string   `json:"component-number"`
		ShortTitle         []string `json:"short-title"`
		Issued             struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"issued"`
		ISBN            []string `json:"ISBN"`
		ReferencesCount int      `json:"references-count"`
		PartNumber      string   `json:"part-number"`
		Aliases         []string `json:"aliases"`
		IssueTitle      []string `json:"issue-title"`
		JournalIssue    struct {
			Issue           string `json:"issue"`
			PublishedOnline struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"published-online"`
			PublishedPrint struct {
				DateParts [][]string `json:"date-parts"`
			} `json:"published-print"`
		} `json:"journal-issue"`
		AlternativeID []string `json:"alternative-id"`
		Version       struct {
			Version            string `json:"version"`
			Language           string `json:"language"`
			VersionDescription []struct {
				Language    string `json:"language"`
				Description string `json:"description"`
			} `json:"version-description"`
		} `json:"version"`
		URL      string   `json:"URL"`
		Archive  []string `json:"archive"`
		Relation struct {
			AdditionalProp1 []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"additionalProp1"`
			AdditionalProp2 []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"additionalProp2"`
			AdditionalProp3 []struct {
				IDType     string `json:"id-type"`
				ID         string `json:"id"`
				AssertedBy string `json:"asserted-by"`
			} `json:"additionalProp3"`
		} `json:"relation"`
		ISSN     []string `json:"ISSN"`
		IssnType []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"issn-type"`
		Subject        []string `json:"subject"`
		PublishedOther struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"published-other"`
		Published struct {
			DateParts [][]string `json:"date-parts"`
		} `json:"published"`
		Assertion []struct {
			Group struct {
				Name  string `json:"name"`
				Label string `json:"label"`
			} `json:"group"`
			Explanation struct {
				URL string `json:"URL"`
			} `json:"explanation"`
			Name  string `json:"name"`
			Value string `json:"value"`
			URL   string `json:"URL"`
			Order int    `json:"order"`
			Label string `json:"label"`
		} `json:"assertion"`
		Subtype       string `json:"subtype"`
		ArticleNumber string `json:"article-number"`
	} `json:"message"`
}

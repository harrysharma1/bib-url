/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/helper"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	articleDOI     string
	articleCiteKey string
	articleAuthors []string
	articleTitle   string
	articleJournal string
	articleYear    string
	articleVolume  string
	articleNumber  string
	articlePages   string
)

// articleCmd represents the article command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "Generate @article bibtex entry",
	Long: `This will generate an @article bibtex entry e.g.
	
@article{CitekeyArticle,
	author   = "<Lastname, FirstName>",
	title    = "<Title>",
	journal  = "<Journal>",
	year     = <1999>,
	volume   = "<50>",
	number   = "<1>",
	pages    = "<1--10>",
	month	 = "<January>",
	note	 = "<Note>"
}

Required:
	- author
	- title
	- journal
	- year
Optional:
	- volume
	- number
	- pages
	- month
	- note
<>: indicates that it is a example value and should be changed.	
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if articleDOI != "" {
			doiAuthors, doiTitle, doiJournal, doiYear, doiVolume, doiNumber, err := helper.ArticleFromDOI(articleDOI)
			if err != nil {
				return err
			}

			if len(articleAuthors) == 0 {
				articleAuthors = doiAuthors
			} else {
				articleAuthors = append(articleAuthors, doiAuthors...)
			}

			if articleTitle == "" {
				articleTitle = doiTitle
			}

			if articleJournal == "" {
				articleJournal = doiJournal
			}

			if articleYear == "" {
				articleYear = doiYear
			}

			if doiVolume == "" {
				articleVolume = doiVolume
			}

			if doiNumber == "" {
				articleNumber = doiNumber
			}
		}
		fmt.Println(articleTitle)
		var bibtex = helper.FormatArticleBibtex(articleCiteKey, articleAuthors, articleTitle, articleJournal, articleYear, articleVolume, articleNumber, articlePages, braces)

		if copy {
			helper.Copy(bibtex)
		}

		if save != "" {
			if err = helper.Save(save, bibtex); err != nil {
				return err
			}
		}

		fmt.Println(bibtex)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(articleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// articleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// articleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	articleCmd.Flags().StringVarP(&articleDOI, "doi", "d", "", "doi identifier for online article to auto-generate entry")
	articleCmd.Flags().StringVarP(&articleCiteKey, "key", "k", "", "citation key for bibtex entry")
	articleCmd.Flags().StringArrayVarP(&articleAuthors, "author", "a", []string{}, "author name(s) for article")
	articleCmd.Flags().StringVarP(&articleTitle, "title", "t", "", "title of article")
	articleCmd.Flags().StringVarP(&articleJournal, "journal", "j", "", "journal the article was published in")
	articleCmd.Flags().StringVarP(&articleYear, "year", "y", "", "year the article was published in")
	articleCmd.Flags().StringVarP(&articleVolume, "volume", "v", "", "volume of journal")
	articleCmd.Flags().StringVarP(&articleNumber, "number", "n", "", "issue of journal")
	articleCmd.Flags().StringVarP(&articlePages, "pages", "p", "", "pages within the journal")
}

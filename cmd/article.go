/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/helper"
	"bibcli/models"
	"fmt"

	"github.com/spf13/cobra"
)

var article models.Article

// articleCmd represents the article command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "Generate @article bibtex entry",
	Long: `This will generate an @article bibtex entry e.g.
	
@article{CitekeyArticle,
	author   = "<Lastname, FirstName>",
	title    = "<Title>",
	journal  = "<Journal>",
	year     = <2002>,
	volume   = "<50>",
	number   = "<1>",
	pages    = "<1--10>",
	month	 = <jul>,
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

		if article.Doi != "" {
			doiAuthors, doiTitle, doiJournal, doiYear, doiVolume, doiNumber, err := helper.ArticleFromDOI(article.Doi)
			if err != nil {
				return err
			}

			if len(article.Authors) == 0 {
				article.Authors = doiAuthors
			} else {
				article.Authors = append(article.Authors, doiAuthors...)
			}

			if article.Title == "" {
				article.Title = doiTitle
			}

			if article.Journal == "" {
				article.Journal = doiJournal
			}

			if article.Year == "" {
				article.Year = doiYear
			}

			if doiVolume == "" {
				article.Volume = doiVolume
			}

			if doiNumber == "" {
				article.Number = doiNumber
			}
		}
		var bibtex = helper.FormatArticleBibtex(
			article.CiteKey,
			article.Authors,
			article.Title,
			article.Journal,
			article.Year,
			article.Volume,
			article.Number,
			article.Pages,
			article.Month,
			article.Note,
			braces)

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
	articleCmd.Flags().StringVarP(&article.Doi, "doi", "d", "", "doi identifier for online article to auto-generate entry")
	articleCmd.Flags().StringVarP(&article.Doi, "key", "k", "", "citation key for bibtex entry")
	articleCmd.Flags().StringArrayVarP(&article.Authors, "author", "a", []string{}, "author name(s) for article")
	articleCmd.Flags().StringVarP(&article.Title, "title", "t", "", "title of article")
	articleCmd.Flags().StringVarP(&article.Journal, "journal", "j", "", "journal the article was published in")
	articleCmd.Flags().StringVarP(&article.Year, "year", "y", "", "year the article was published in")
	articleCmd.Flags().StringVarP(&article.Volume, "volume", "v", "", "volume of journal")
	articleCmd.Flags().StringVarP(&article.Number, "number", "i", "", "issue of journal")
	articleCmd.Flags().StringVarP(&article.Pages, "pages", "p", "", "pages within the journal")
	articleCmd.Flags().StringVarP(&article.Number, "month", "m", "", "month the article was published")
	articleCmd.Flags().StringVarP(&article.Note, "note", "n", "", "additional notes for article")
}

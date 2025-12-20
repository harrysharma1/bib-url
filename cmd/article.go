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
	articleUrl     string
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
  		author   = "P. J. Cohen",
  		title    = "The independence of the continuum hypothesis",
  		journal  = "Proceedings of the National Academy of Sciences",
 	 	year     = 1963,
  		volume   = "50",
  		number   = "6",
  		pages    = "1143--1148",
	}
	

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if helper.IsValidUrl(articleUrl) {
			helper.ScrapeArticle(articleUrl)
		} else {
			return fmt.Errorf("%s is not a valid URL", articleUrl)
		}

		var bibtex = helper.FormatArticleBibtex(articleCiteKey, articleAuthors, articleTitle, articleJournal, articleYear, articleVolume, articleNumber, articlePages, braces)

		if copy {
			helper.Copy(bibtex)
		}
		if save != "" {
			if err := helper.Save(save, bibtex); err != nil {
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
	articleCmd.Flags().StringVarP(&articleUrl, "url", "u", "", "url for online article to auto-generate entry")
	articleCmd.Flags().StringVarP(&articleCiteKey, "key", "k", "", "citation key for bibtex entry")
	articleCmd.Flags().StringArrayVarP(&articleAuthors, "author", "a", []string{}, "author name(s) for article")
	articleCmd.Flags().StringVarP(&articleTitle, "title", "t", "", "title of article")
	articleCmd.Flags().StringVarP(&articleJournal, "journal", "j", "", "journal the article was published in")
	articleCmd.Flags().StringVarP(&articleYear, "year", "y", "", "year the article was published in")
	articleCmd.Flags().StringVarP(&articleVolume, "volume", "v", "", "volume of journal")
	articleCmd.Flags().StringVarP(&articleNumber, "number", "n", "", "issue of journal")
	articleCmd.Flags().StringVarP(&articlePages, "pages", "p", "", "pages within the journal")
}

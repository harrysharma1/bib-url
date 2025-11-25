/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/helper"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	articleUrl     string
	articleCiteKey string
	articleAuthors []string
	articleTitle   string
	articleJournal string
	articleYear    int
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
		var bibtex = helper.FormatArticleBibtex(articleCiteKey, articleAuthors, articleTitle, articleJournal, articleYear, articleVolume, articleNumber, articlePages)
		if articleUrl != "" {

		}
		if copy {
			fmt.Println("copied bibtex entry to cliplboard!!!")
			clipboard.Write(clipboard.FmtText, []byte(bibtex))
		}
		if save != "" {
			if filepath.Ext(save) == ".bib" {
				f, err := os.OpenFile(save, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					return fmt.Errorf("error opening file")
				}
				defer f.Close()
				bibtex += "\n"
				_, err = f.WriteString(bibtex)
				if err != nil {
					return fmt.Errorf("error writing to file")
				}
				fmt.Printf("saved bibtex entry to `%s`!!!\n", save)
				fmt.Println(bibtex[:len(bibtex)-1])
				return nil
			} else {
				return fmt.Errorf("error file type not .bib did not write entry")
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
	articleCmd.Flags().StringVarP(&articleUrl, "url", "u", "https://example.com", "url for online article to auto-generate entry")
	articleCmd.Flags().StringVarP(&articleCiteKey, "key", "k", "uuid.uuid4()", "citation key for bibtex entry")
	articleCmd.Flags().StringArrayVarP(&articleAuthors, "author", "a", []string{"P. J. Cohen"}, "author name(s) for article")
	articleCmd.Flags().StringVarP(&articleTitle, "title", "t", "Title", "title of article")
	articleCmd.Flags().StringVarP(&articleJournal, "journal", "j", "Journal", "journal the article was published in")
	articleCmd.Flags().IntVarP(&articleYear, "year", "y", 2002, "year the article was published in")
	articleCmd.Flags().StringVarP(&articleVolume, "volume", "v", "1", "volume of journal")
	articleCmd.Flags().StringVarP(&articleNumber, "number", "n", "1", "issue of journal")
	articleCmd.Flags().StringVarP(&articlePages, "pages", "p", "1--10", "pages within the journal")
}

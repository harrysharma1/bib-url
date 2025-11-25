/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	url     string
	citeKey string
	author  string
	title   string
	journal string
	year    int
	volume  string
	number  string
	pages   string
)

// articleCmd represents the article command
var articleCmd = &cobra.Command{
	Use:   "article",
	Short: "Generate article bibtex entry",
	Long: `This will generate an @article bibtex entry with the following:
	
	@article{CitekeyArticle
		author = "",
		title = "",
		journal = "",
		year = "",
		volume = "",
		number = "",
		pages = "",
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("article called")
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
	articleCmd.Flags().StringVarP(&url, "url", "u", "https://example.com", "url for online entry to autogenerate entry")
	articleCmd.Flags().StringVarP(&citeKey, "key", "k", "uuid.uuid4()", "citation key for bibtex entry")
	articleCmd.Flags().StringVarP(&author, "author", "a", "Surname, Forename", "author name(s) for article")
	articleCmd.Flags().StringVarP(&title, "title", "t", "Title", "title name for article")
	articleCmd.Flags().StringVarP(&journal, "journal", "j", "Journal", "journal the article was published in")
	articleCmd.Flags().IntVarP(&year, "year", "y", 2002, "year the article was published in")
	articleCmd.Flags().StringVarP(&volume, "volume", "v", "1", "volume of journal")
	articleCmd.Flags().StringVarP(&number, "number", "n", "1", "issue of journal")
	articleCmd.Flags().StringVarP(&pages, "pages", "p", "1--10", "pages within the journal")
}

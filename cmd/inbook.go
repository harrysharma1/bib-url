/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	inbookUrl       string
	inbookCiteKey   string
	inbookAuthors   []string
	inbookBookTitle string
	inbookYear      string
	inbookPublisher string
	inbookAddress   string
	inbookPages     string
)

// inbookCmd represents the inbook command
var inbookCmd = &cobra.Command{
	Use:   "inbook",
	Short: "Generate @inbook bibtex entry",
	Long: `This will generate an @inbook bibtex entry e.g.
	
	@inbook{CitekeyInbook,
  		author    = "Lisa A. Urry and Michael L. Cain and Steven A. Wasserman and Peter V. Minorsky and Jane B. Reece",
  		title     = "Photosynthesis",
  		booktitle = "Campbell Biology",
  		year      = "2016",
  		publisher = "Pearson",
  		address   = "New York, NY",
  		pages     = "187--221"
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inbook called")
	},
}

func init() {
	rootCmd.AddCommand(inbookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inbookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inbookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inbookCmd.Flags().StringVarP(&inbookUrl, "url", "u", "", "url for online inbook to auto-generate entry")
	inbookCmd.Flags().StringVarP(&inbookCiteKey, "key", "k", "", "citation key for bibtex entry")
	inbookCmd.Flags().StringArrayVarP(&inbookAuthors, "author", "a", []string{}, "author name(s) for inbook")
	inbookCmd.Flags().StringVarP(&inbookBookTitle, "title", "t", "", "title of inbook")
	inbookCmd.Flags().StringVarP(&inbookYear, "year", "y", "", "year the inbook was released")
	inbookCmd.Flags().StringVarP(&inbookPublisher, "publisher", "p", "", "who published the inbook")
	inbookCmd.Flags().StringVarP(&inbookAddress, "address", "l", "", "address of publisher")
	inbookCmd.Flags().StringVarP(&inbookPages, "pages", "f", "", "pages within the inbook")
}

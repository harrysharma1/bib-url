/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/format"
	"bibcli/helper"
	"bibcli/models"
	"fmt"

	"github.com/spf13/cobra"
)

var inbook models.Inbook

// inbookCmd represents the inbook command
var inbookCmd = &cobra.Command{
	Use:   "inbook",
	Short: "Generate @inbook bibtex entry",
	Long: `This will generate an @inbook bibtex entry e.g.
	
@inbook{CitekeyInbook,
	author    = "<Lastname, Firstname>",
	title     = "<Title>",
	booktitle = "<Book Title>",
	publisher = "<Publisher>",
	year 	  = <2002>,
	editor	  = "<Lastname, Firstname>",
	volume	  = "<50>",
	number	  = "<1>",
	series	  = "<Series>",
	address   = "<Address>",
	edition	  = "<2nd>",
	month	  = <jul>,
	pages     = "<10--100>",
	note	  = "<Note>",
}

Required:
	- author
	- title
	- booktitle
	- publisher
	- year
Optional:
	- editor
	- volume
	- number
	- series
	- address
	- edition
	- month
	- pages
	- note
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatInbookBibtex(
			inbook.CiteKey,
			inbook.Authors,
			inbook.Title,
			inbook.BookTitle,
			inbook.Publisher,
			inbook.Year,
			inbook.Editors,
			inbook.Volume,
			inbook.Number,
			inbook.Series,
			inbook.Address,
			inbook.Edition,
			inbook.Month,
			inbook.Pages,
			inbook.Note,
			braces)

		if copy {
			helper.Copy(bibtex)
		}

		if save != "" {
			helper.Save(save, bibtex)
		}
		fmt.Println(bibtex)
		return nil
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
	inbookCmd.Flags().StringVarP(&inbook.CiteKey, "key", "k", "", "citation key for bibtex entry")
	inbookCmd.Flags().StringArrayVarP(&inbook.Authors, "author", "a", []string{}, "author name(s) for inbook")
	inbookCmd.Flags().StringVarP(&inbook.Title, "title", "t", "", "title of inbook")
	inbookCmd.Flags().StringVarP(&inbook.BookTitle, "booktitle", "b", "", "book title of inbook")
	inbookCmd.Flags().StringVarP(&inbook.Publisher, "publisher", "p", "", "who published the inbook")
	inbookCmd.Flags().StringVarP(&inbook.Year, "year", "y", "", "year the inbook was released")
	inbookCmd.Flags().StringArrayVarP(&inbook.Editors, "editor", "e", []string{}, "editor name(s) for inbook")
	inbookCmd.Flags().StringVarP(&inbook.Volume, "volume", "v", "", "volume of inbook")
	inbookCmd.Flags().StringVarP(&inbook.Number, "number", "i", "", "number of inbook")
	inbookCmd.Flags().StringVarP(&inbook.Series, "series", "s", "", "series where inbook was published")
	inbookCmd.Flags().StringVarP(&inbook.Address, "address", "l", "", "address of publisher")
	inbookCmd.Flags().StringVarP(&inbook.Edition, "edition", "r", "", "edition of publication inbook was published in")
	inbookCmd.Flags().StringVarP(&inbook.Month, "month", "m", "", "month the inbook was released")
	inbookCmd.Flags().StringVarP(&inbook.Pages, "pages", "f", "", "pages within the inbook")
	inbookCmd.Flags().StringVarP(&inbook.Note, "note", "n", "", "additional notes for inbook")
}

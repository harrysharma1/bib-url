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
	inbookUrl       string
	inbookCiteKey   string
	inbookAuthors   []string
	inbookTitle     string
	inbookBookTitle string
	inbookPublisher string
	inbookYear      string
	inbookEditors   []string
	inbookVolume    string
	inbookNumber    string
	inbookSeries    string
	inbookAddress   string
	inbookEdition   string
	inboookMonth    string
	inbookPages     string
	inbookNote      string
)

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
		var bibtex = helper.FormatInbookBibtex(
			inbookCiteKey,
			inbookAuthors,
			inbookTitle,
			inbookBookTitle,
			inbookPublisher,
			inbookYear,
			inbookEditors,
			inbookVolume,
			inbookNumber,
			inbookSeries,
			inbookAddress,
			inbookEdition,
			inboookMonth,
			inbookPages,
			inbookNote,
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
	inbookCmd.Flags().StringVarP(&inbookUrl, "url", "u", "", "url for online inbook to auto-generate entry")
	inbookCmd.Flags().StringVarP(&inbookCiteKey, "key", "k", "", "citation key for bibtex entry")
	inbookCmd.Flags().StringArrayVarP(&inbookAuthors, "author", "a", []string{}, "author name(s) for inbook")
	inbookCmd.Flags().StringVarP(&inbookTitle, "title", "t", "", "title of inbook")
	inbookCmd.Flags().StringVarP(&inbookBookTitle, "booktitle", "b", "", "book title of inbook")
	inbookCmd.Flags().StringVarP(&inbookPublisher, "publisher", "p", "", "who published the inbook")
	inbookCmd.Flags().StringVarP(&inbookYear, "year", "y", "", "year the inbook was released")
	inbookCmd.Flags().StringArrayVarP(&inbookEditors, "editor", "e", []string{}, "editor name(s) for inbook")
	inbookCmd.Flags().StringVarP(&inbookVolume, "volume", "v", "", "volume of inbook")
	inbookCmd.Flags().StringVarP(&inbookNumber, "number", "i", "", "number of inbook")
	inbookCmd.Flags().StringVarP(&inbookSeries, "series", "s", "", "series where inbook was published")
	inbookCmd.Flags().StringVarP(&inbookAddress, "address", "l", "", "address of publisher")
	inbookCmd.Flags().StringVarP(&inbookEdition, "edition", "r", "", "edition of publication inbook was published in")
	inbookCmd.Flags().StringVarP(&inboookMonth, "month", "m", "", "month the inbook was released")
	inbookCmd.Flags().StringVarP(&inbookPages, "pages", "f", "", "pages within the inbook")
	inbookCmd.Flags().StringVarP(&inbookNote, "note", "n", "", "additional notes for inbook")
}

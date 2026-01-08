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

var inproceedings models.Inproceedings

// inproceedingsCmd represents the inproceedings command
var inproceedingsCmd = &cobra.Command{
	Use:   "inproceedings",
	Short: "Generate @inproceedings bibtex entry",
	Long: `This will generate an @inproceedings bibtex entry e.g.
	
@inproceedings{CitekeyInproceedings,
	author       = "<Lastname, Firstname",
	title        = "<Title>",
	booktitle    = "<Booktitle>",	
	year         = <2002>,
	editor       = "<Lastname, Firstname>",
	volume       = "<2>",
	number       = "<50>",
	series       = "<Series>",
	pages        = "1--10",
	address      = "<Address>",
	month        = <jul>,
	organisation = "<Organisation>",
	publisher    = "<Publisher>",
}

Required:
	- author
	- title
	- booktitle
	- year
Optional:
	- editor
	- volume
	- number
	- series
	- pages
	- address
	- month
	- organisation
	- publisher
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatInproceedingsBibtex(
			inproceedings.CiteKey,
			inproceedings.Authors,
			inproceedings.Title,
			inproceedings.BooktTitle,
			inproceedings.Year,
			inproceedings.Editors,
			inproceedings.Volume,
			inproceedings.Number,
			inproceedings.Series,
			inproceedings.Pages,
			inproceedings.Address,
			inproceedings.Month,
			inproceedings.Organisation,
			inproceedings.Publisher,
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
	rootCmd.AddCommand(inproceedingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inproceedingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inproceedingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.CiteKey, "key", "k", "", "citation key for bibtex entry")
	inproceedingsCmd.Flags().StringArrayVarP(&inproceedings.Authors, "author", "a", []string{}, "author name(s) for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Title, "title", "t", "", "title for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.BooktTitle, "booktitle", "b", "", "book title for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Year, "year", "y", "", "year the inproceedings was released")
	inproceedingsCmd.Flags().StringArrayVarP(&inproceedings.Editors, "editor", "e", []string{}, "editor name(s) for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Volume, "volume", "v", "", "volume of inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Number, "number", "i", "", "number of inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Series, "series", "s", "", "series of inproceedings.")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Pages, "pages", "f", "", "pages where citation is")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Address, "address", "l", "", "address of publisher")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Month, "month", "m", "", "month the inproceedings was released")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Organisation, "organisation", "o", "", "organisation that released inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedings.Publisher, "publisher", "p", "", "publisher of inproceedings")
}

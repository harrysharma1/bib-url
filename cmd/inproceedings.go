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
	inproceedingsUrl       string
	inproceedingsCiteKey   string
	inproceedingsAuthors   []string
	inproceedingsTitle     string
	inproceedingsBookTitle string
	inproceedingsSeries    string
	inproceedingsYear      string
	inproceedingsPages     string
	inproceedingsPublisher string
	inproceedingsAddress   string
)

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
	editor	     = "<Lastname, Firstname>",
	volume	     = "<2>",
	number	     = "<50>",
	series	     = "<Series>",
	pages        = "1--10",
	address   	 = "<Address>",
	month	  	 = <jul>,
	organisation = "<Organisation>",
	publisher    = "<Publusher>",
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
		var bibtex = helper.FormatInproceedingsBibtex(inproceedingsCiteKey, inproceedingsAuthors, inproceedingsTitle, inproceedingsBookTitle, inproceedingsSeries, inproceedingsYear, inproceedingsPages, inproceedingsPublisher, inproceedingsAddress, braces)

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
	inproceedingsCmd.Flags().StringVarP(&inproceedingsUrl, "url", "u", "", "url for online inproceedings to auto-generate entry")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsCiteKey, "key", "k", "", "citation key for bibtex entry")
	inproceedingsCmd.Flags().StringArrayVarP(&inproceedingsAuthors, "author", "a", []string{}, "author name(s) for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsTitle, "title", "t", "", "title for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsBookTitle, "booktitle", "b", "", "book title for inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsSeries, "series", "s", "", "series of inproceedings.")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsYear, "year", "y", "", "year the inproceedings was released")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsPages, "pages", "f", "", "pages where citation is")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsPublisher, "publisher", "p", "", "publisher of inproceedings")
	inproceedingsCmd.Flags().StringVarP(&inproceedingsAddress, "address", "l", "", "address of publisher")
}

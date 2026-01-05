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

var incollection models.Incollection

// incollectionCmd represents the incollection command
var incollectionCmd = &cobra.Command{
	Use:   "incollection",
	Short: "Generate @incollection bibtex entry",
	Long: `This will generate an @incollection bibtex entry e.g.
	
@incollection{CitekeyIncollection,
	author    = "<Lastname, Firstname>",
	title     = "<Title>",
	booktitle = "<Book Title>",
	publisher = "<Publisher>",
	year      = <2002>,
	editor    = "<Lastname, Firstname>",
	volume    = "<2>",
	number    = "<50>",
	series    = "<Series>",
	pages     = "<1--10>",
	address   = "<Address>",
	month     = <jul>
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
	- pages
	- address
	- month
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatIncollectionBibtex(
			incollection.CiteKey,
			incollection.Authors,
			incollection.Title,
			incollection.BookTitle,
			incollection.Publisher,
			incollection.Year,
			incollection.Editors,
			incollection.Volume,
			incollection.Number,
			incollection.Series,
			incollection.Pages,
			incollection.Address,
			incollection.Month,
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
	rootCmd.AddCommand(incollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incollectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incollectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	incollectionCmd.Flags().StringVarP(&incollection.CiteKey, "key", "k", "", "citation key for bibtex entry")
	incollectionCmd.Flags().StringArrayVarP(&incollection.Authors, "author", "a", []string{}, "author name(s) for incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Title, "title", "t", "", "title for incollection")
	incollectionCmd.Flags().StringVarP(&incollection.BookTitle, "booktitle", "b", "", "book title for incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Publisher, "publisher", "p", "", "publisher of incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Year, "year", "y", "", "year the incollection was released")
	incollectionCmd.Flags().StringArrayVarP(&incollection.Editors, "editor", "e", []string{}, "editor name(s) for incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Volume, "volume", "v", "", "volume of incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Number, "number", "i", "", "number of incollection")
	incollectionCmd.Flags().StringVarP(&incollection.Series, "series", "s", "", "series incollection was published in")
	incollectionCmd.Flags().StringVarP(&incollection.Pages, "pages", "f", "", "pages where citation is")
	incollectionCmd.Flags().StringVarP(&incollection.Address, "address", "l", "", "address of publisher")
	incollectionCmd.Flags().StringVarP(&incollection.Month, "month", "m", "", "month the incollection was released")
}

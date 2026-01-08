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

var booklet models.Booklet

// bookletCmd represents the booklet command
var bookletCmd = &cobra.Command{
	Use:   "booklet",
	Short: "Generate @booklet bibtex entry",
	Long: `This will generate an @booklet bibtex entry e.g.
	
@booklet{CitekeyBooklet,
	title        = "<Title>",
	author       = "<Lastname, Firstname>", 
	howpublished = "<How Published>",
	address 	 = "<Address>",
	year         = <2002>,
	editor		 = "<Lastname, Firstname>",
	volume		 = "<50>",
	number		 = "<1>",
	series		 = "<Series>",
	organisation = "<Organisation>",
	month		 = <jul>,
	note		 = "<Note>"
}

Required:
	- title
	- author
	- howpublished
	- address
	- year
Optional:
	- editor
	- volume
	- number
	- series
	- organisation
	- month
	- note
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatBookletBibtex(
			booklet.CiteKey,
			booklet.Title,
			booklet.Authors,
			booklet.Address,
			booklet.HowPublished,
			booklet.Year,
			booklet.Editors,
			booklet.Volume,
			booklet.Number,
			booklet.Series,
			booklet.Organisation,
			booklet.Month,
			booklet.Note,
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
	rootCmd.AddCommand(bookletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bookletCmd.Flags().StringVarP(&booklet.CiteKey, "key", "k", "", "citation key for bibtex entry")
	bookletCmd.Flags().StringVarP(&booklet.Title, "title", "t", "", "title of booklet")
	bookletCmd.Flags().StringArrayVarP(&booklet.Authors, "author", "a", []string{}, "author name(s) for booklet")
	bookletCmd.Flags().StringVarP(&booklet.HowPublished, "howpublished", "p", "", "how the booklet was published")
	bookletCmd.Flags().StringVarP(&booklet.Address, "address", "l", "", "address of publisher")
	bookletCmd.Flags().StringVarP(&booklet.Year, "year", "y", "", "year the booklet was released")
	bookletCmd.Flags().StringArrayVarP(&booklet.Editors, "editor", "e", []string{}, "editor name(s) for booklet")
	bookletCmd.Flags().StringVarP(&booklet.Volume, "volume", "v", "", "volume of booklet")
	bookletCmd.Flags().StringVarP(&booklet.Number, "number", "i", "", "number of booklet")
	bookletCmd.Flags().StringVarP(&booklet.Series, "series", "s", "", "series the booklet is in")
	bookletCmd.Flags().StringVarP(&booklet.Organisation, "organisation", "o", "", "organisation that published booklet")
	bookletCmd.Flags().StringVarP(&booklet.Month, "month", "m", "", "month the booklet was released")
	bookletCmd.Flags().StringVarP(&booklet.Note, "note", "n", "", "additional notes for booklet")
}

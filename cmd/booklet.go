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
	bookletUrl          string
	bookletCiteKey      string
	bookletTitle        string
	bookletAuthors      []string
	bookletHowPublished string
	bookletAddress      string
	bookletYear         string
	bookletEditors      []string
	bookletVolume       string
	bookletNumber       string
	bookletSeries       string
	bookletOrganisation string
	bookletMonth        string
	bookletNote         string
)

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
		var bibtex = helper.FormatBookletBibtex(bookletCiteKey, bookletTitle, bookletAuthors, bookAddress, bookletHowPublished, bookletYear, bookletEditors, bookletVolume, bookletNumber, bookletSeries, bookletOrganisation, bookletMonth, bookletNote, braces)

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
	bookletCmd.Flags().StringVarP(&bookletUrl, "url", "u", "", "url for online booklet to auto-generate entry")
	bookletCmd.Flags().StringVarP(&bookletCiteKey, "key", "k", "", "citation key for bibtex entry")
	bookletCmd.Flags().StringVarP(&bookletTitle, "title", "t", "", "title of booklet")
	bookletCmd.Flags().StringArrayVarP(&bookletAuthors, "author", "a", []string{}, "author name(s) for booklet")
	bookletCmd.Flags().StringVarP(&bookletHowPublished, "published", "p", "", "how the booklet was published")
	bookletCmd.Flags().StringVarP(&bookletYear, "year", "y", "", "year the booklet was released")
	bookletCmd.Flags().StringArrayVarP(&bookletEditors, "editor", "e", []string{}, "editor name(s) for booklet")
	bookletCmd.Flags().StringVarP(&bookletVolume, "volume", "v", "", "volume of booklet")
	bookletCmd.Flags().StringVarP(&bookletNumber, "number", "i", "", "number of booklet")
	bookletCmd.Flags().StringVarP(&bookletSeries, "series", "s", "", "series the booklet is in")
	bookletCmd.Flags().StringVarP(&bookletOrganisation, "organisation", "o", "", "organisation that published booklet")
	bookletCmd.Flags().StringVarP(&bookletMonth, "month", "m", "", "month the booklet was released")
	bookletCmd.Flags().StringVarP(&bookletNote, "note", "n", "", "additional notes for booklet")
}

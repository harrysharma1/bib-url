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
	bookletMonth        string
	bookletYear         string
)

// bookletCmd represents the booklet command
var bookletCmd = &cobra.Command{
	Use:   "booklet",
	Short: "Generate @booklet bibtex entry",
	Long: `This will generate an @booklet bibtex entry e.g.
	
	@booklet{CitekeyBooklet,
 		title        = "Canoe tours in {S}weden",
  		author       = "Maria Swetla", 
  		howpublished = "Distributed at the Stockholm Tourist Office",
  		month        = jul,
  		year         = 2015
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatBookletBibtex(bookletCiteKey, bookletTitle, bookletAuthors, bookletHowPublished, bookletMonth, bookletYear, braces)

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
	bookletCmd.Flags().StringArrayVarP(&bookletAuthors, "author", "a", []string{}, "author name(s) for book")
	bookletCmd.Flags().StringVarP(&bookletHowPublished, "published", "p", "", "how the booklet was published")
	bookletCmd.Flags().StringVarP(&bookletMonth, "month", "m", "", "month the booklet was released")
	bookletCmd.Flags().StringVarP(&bookletYear, "year", "y", "", "year the booklet was released")
}

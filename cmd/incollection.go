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
	incollectionUrl       string
	incollectionCiteKey   string
	incollectionAuthors   []string
	incollectionEditors   []string
	incollectionTitle     string
	incollectionBookTitle string
	incollectionYear      string
	incollectionPublisher string
	incollectionAddress   string
	incollectionPages     string
)

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
	year	  = <2002>,
	editor	  = "<Lastname, Firstname>",
	volume	  = "<2>",
	number	  = "<50>",
	series	  = "<Series>",
	pages     = "<1--10>",
	address	  = "<Address>",
	month	  = <jul>
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
		var bibtex = helper.FormatIncollectionBibtex(incollectionCiteKey, incollectionAuthors, incollectionEditors, incollectionTitle, incollectionBookTitle, incollectionYear, incollectionPublisher, incollectionAddress, incollectionPages, braces)

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
	incollectionCmd.Flags().StringVarP(&incollectionUrl, "url", "u", "", "url for online incollection to auto-generate entry")
	incollectionCmd.Flags().StringVarP(&incollectionCiteKey, "key", "k", "", "citation key for bibtex entry")
	incollectionCmd.Flags().StringArrayVarP(&incollectionAuthors, "author", "a", []string{}, "author name(s) for incollection")
	incollectionCmd.Flags().StringArrayVarP(&incollectionEditors, "editor", "e", []string{}, "editor name(s) for incollection")
	incollectionCmd.Flags().StringVarP(&incollectionTitle, "title", "t", "", "title for incollection")
	incollectionCmd.Flags().StringVarP(&incollectionBookTitle, "booktitle", "b", "", "book title for incollection")
	incollectionCmd.Flags().StringVarP(&incollectionYear, "year", "y", "", "year the incollection was released")
	incollectionCmd.Flags().StringVarP(&incollectionPublisher, "publisher", "p", "", "publisher of incollection")
	incollectionCmd.Flags().StringVarP(&incollectionAddress, "address", "l", "", "address of publisher")
	incollectionCmd.Flags().StringVarP(&incollectionPages, "pages", "f", "", "pages where citation is")
}

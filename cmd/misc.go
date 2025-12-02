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
	miscUrl          string
	miscCiteKey      string
	miscTitle        string
	miscAuthors      []string
	miscHowPublished string
	miscYear         string
	miscNote         string
)

// miscCmd represents the misc command
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Generate @misc bibtex entry",
	Long: `This will generate an @misc bibtex entry e.g.
	
	@misc{CitekeyMisc,
  		title        = "Pluto: The 'Other' Red Planet",
  		author       = "{NASA}",
  		howpublished = "\url{https://www.nasa.gov/nh/pluto-the-other-red-planet}",
  		year         = 2015,
  		note         = "Accessed: 2018-12-06"
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatMiscBibtex()

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
	rootCmd.AddCommand(miscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// miscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// miscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	miscCmd.Flags().StringVarP(&miscUrl, "url", "u", "", "url for online masters thesis to auto-generate entry")
	miscCmd.Flags().StringVarP(&miscCiteKey, "key", "k", "", "citation key for bibtex entry")
	miscCmd.Flags().StringVarP(&miscTitle, "title", "t", "", "title for misc entry")
	miscCmd.Flags().StringArrayVarP(&miscAuthors, "authors", "a", []string{}, "author name(s) for misc")
	miscCmd.Flags().StringVarP(&miscHowPublished, "published", "p", "", "how misc was published")
	miscCmd.Flags().StringVarP(&miscYear, "year", "y", "", "year misc entry was created")
	miscCmd.Flags().StringVarP(&miscNote, "note", "n", "", "note, usually when last accessed")
}

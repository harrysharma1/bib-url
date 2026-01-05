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

var misc models.Misc

// miscCmd represents the misc command
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Generate @misc bibtex entry",
	Long: `This will generate an @misc bibtex entry e.g.
	
@misc{CitekeyMisc,
	title        = "<Title>",
	author       = "<Lastname, Firstname>",
	howpublished = "<How published>",
	year         = <2002>,
	note         = "<Note>"
}

Required:
	- title
	- author
	- howpublished
	- year
	- note
Optional:
	- !*
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatMiscBibtex(
			misc.CiteKey,
			misc.Title,
			misc.Authors,
			misc.HowPublished,
			misc.Year,
			misc.Note,
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
	rootCmd.AddCommand(miscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// miscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// miscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	miscCmd.Flags().StringVarP(&misc.CiteKey, "key", "k", "", "citation key for bibtex entry")
	miscCmd.Flags().StringVarP(&misc.Title, "title", "t", "", "title of misc entry")
	miscCmd.Flags().StringArrayVarP(&misc.Authors, "author", "a", []string{}, "author name(s) of misc entry")
	miscCmd.Flags().StringVarP(&misc.HowPublished, "howpublished", "p", "", "how misc entry was published")
	miscCmd.Flags().StringVarP(&misc.Year, "year", "y", "", "year misc entry was published")
	miscCmd.Flags().StringVarP(&misc.Note, "note", "n", "", "additional notes for misc entry")
}

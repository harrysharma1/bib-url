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
	unpublishedUrl     string
	unpublishedCiteKey string
	unpublishedAuthors []string
	unpublishedTitle   string
	unpublishedYear    string
)

// unpublishedCmd represents the unpublished command
var unpublishedCmd = &cobra.Command{
	Use:   "unpublished",
	Short: "Generate @unpublished bibtex entry",
	Long: `This will generate an @unpublished bibtex entry e.g.
	
@unpublished{CitekeyUnpublished,
	author = "<Lastname, Firstname>",
	title  = "<Title>",
	institution = "<Institution>",
	year   = <2002>
}

Required:
	- author
	- title
	- instituion
	- year
Optional:
	- *
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatUnpublishedBibtex(unpublishedCiteKey, unpublishedAuthors, unpublishedTitle, unpublishedYear, braces)

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
	rootCmd.AddCommand(unpublishedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unpublishedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unpublishedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	unpublishedCmd.Flags().StringVarP(&unpublishedUrl, "url", "u", "", "url for online unpublished work to auto-generate entry")
	unpublishedCmd.Flags().StringVarP(&unpublishedCiteKey, "key", "k", "", "citation key for bibtex entry")
	unpublishedCmd.Flags().StringArrayVarP(&unpublishedAuthors, "author", "a", []string{}, "author name(s) for unpublished work")
	unpublishedCmd.Flags().StringVarP(&unpublishedTitle, "title", "t", "", "title of unpublished work")
	unpublishedCmd.Flags().StringVarP(&unpublishedYear, "year", "y", "", "year unpublished work was created")
}

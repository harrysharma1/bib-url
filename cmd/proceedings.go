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

var proceedings models.Proceedings

// proceedingsCmd represents the proceedings command
var proceedingsCmd = &cobra.Command{
	Use:   "proceedings",
	Short: "Generate @proceedings bibtex entry",
	Long: `This will generate an @proceedings entry e.g.
	
@proceedings{CitekeyProceedings,
	title     = "<Title>",
	year	  = <2002>,
	editor	  = "<Lastname, Firstname>",
	volume    = "<1>",
	number	  = "<50>",
	series	  = "<Series>",
	address   = "<Address>",
	month	  = <jul>,
	publisher = "<Publisher>"
}

Required:
	- title
	- year
Optional:
	- editor
	- volume
	- number
	- series
	- address
	- month
	- publisher
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatProceedingsBibtex(
			proceedings.CiteKey,
			proceedings.Title,
			proceedings.Year,
			proceedings.Editors,
			proceedings.Volume,
			proceedings.Number,
			proceedings.Series,
			proceedings.Address,
			proceedings.Month,
			proceedings.Publisher,
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
	rootCmd.AddCommand(proceedingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proceedingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proceedingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	proceedingsCmd.Flags().StringVarP(&proceedings.CiteKey, "key", "k", "", "citation key for bibtex entry")
	proceedingsCmd.Flags().StringVarP(&proceedings.Title, "title", "t", "", "title of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedings.Year, "year", "y", "", "year proceedings was published")
	proceedingsCmd.Flags().StringArrayVarP(&proceedings.Editors, "editor", "e", []string{}, "editor name(s) for proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedings.Volume, "volume", "v", "", "volume of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedings.Number, "number", "i", "", "number of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedings.Series, "series", "s", "", "series where proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedings.Address, "address", "l", "", "location of publisher")
	proceedingsCmd.Flags().StringVarP(&proceedings.Month, "month", "m", "", "month proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedings.Publisher, "publisher", "p", "", "the group that published proceedings")
}

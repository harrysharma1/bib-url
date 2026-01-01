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
	proceedingsUrl       string
	proceedingsCiteKey   string
	proceedingsTitle     string
	proceedingsYear      string
	proceedingsEditors   []string
	proceedingsVolume    string
	proceedingsNumber    string
	proceedingsSeries    string
	proceedingsAddress   string
	proceedingsMonth     string
	proceedingsPublisher string
)

// proceedingsCmd represents the proceedings command
var proceedingsCmd = &cobra.Command{
	Use:   "proceedings",
	Short: "Generate @proceedings bibtex entry",
	Long: `This will generate an @proceedings entry e.g.
	
@proceedings{CitekeyProceedings,
	title     = "<Title>",
	year	  = <2002>,
	editor	  = "<Lastname, Firstname>",
	volume    = "10867",
	number	  = "<50>",
	series	  = "<Series>,
	address   = "Cham, Switzerland",
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
		var bibtex = helper.FormatProceedingsBibtex(
			proceedingsCiteKey,
			proceedingsTitle,
			proceedingsYear,
			proceedingsEditors,
			proceedingsVolume,
			proceedingsNumber,
			proceedingsSeries,
			proceedingsAddress,
			proceedingsMonth,
			proceedingsPublisher,
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

	proceedingsCmd.Flags().StringVarP(&proceedingsUrl, "url", "u", "", "url for online proceedings to auto-generate entry")
	proceedingsCmd.Flags().StringVarP(&proceedingsCiteKey, "key", "k", "", "citation key for bibtex entry")
	proceedingsCmd.Flags().StringVarP(&proceedingsTitle, "title", "t", "", "title of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsYear, "year", "y", "", "year proceedings was published")
	proceedingsCmd.Flags().StringArrayVarP(&proceedingsEditors, "editors", "e", []string{}, "editor name(s) for proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsVolume, "volume", "v", "", "volume of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsNumber, "number", "i", "", "number of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsSeries, "series", "s", "", "series where proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedingsAddress, "address", "l", "", "location of publisher")
	proceedingsCmd.Flags().StringVarP(&proceedingsMonth, "month", "m", "", "month proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedingsPublisher, "publisher", "p", "", "the group that published proceedings")
}

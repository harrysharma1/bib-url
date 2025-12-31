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
	manualUrl          string
	manualCiteKey      string
	manualTitle        string
	manualAuthors      []string
	manualOrganisation string
	manualAddress      string
	manualYear         string
)

// manualCmd represents the manual command
var manualCmd = &cobra.Command{
	Use:   "manual",
	Short: "Generate @manual bibtex entry",
	Long: `This will generate an @manual bibtex entry e.g.
	
@manual{CitekeyManual,
	title        = "<Title>",
	year		 = <2002>,
	author       = "<Lastname, Firstname>",
	organisation = "<Organisation",
	address      = "<Address>",
	edition		 = "<2nd>",
	month		 = <jul>,
	note		 = "<Note>
}

Required:
	- title
	- year
Optional:
	- author
	- organisation
	- address
	- edition
	- month
	- note
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatManualBibtex(manualCiteKey, manualTitle, manualAuthors, manualOrganisation, manualAddress, manualYear, braces)

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
	rootCmd.AddCommand(manualCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manualCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manualCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	manualCmd.Flags().StringVarP(&manualUrl, "url", "u", "", "url for online manual to auto-generate entry")
	manualCmd.Flags().StringVarP(&manualCiteKey, "key", "k", "", "citation key for bibtex entry")
	manualCmd.Flags().StringVarP(&manualTitle, "title", "t", "", "title for manual")
	manualCmd.Flags().StringArrayVarP(&manualAuthors, "author", "a", []string{}, "author name(s) for manual")
	manualCmd.Flags().StringVarP(&manualOrganisation, "organisation", "o", "", "organisation that released the manual")
	manualCmd.Flags().StringVarP(&manualAddress, "address", "l", "", "address of organisation")
	manualCmd.Flags().StringVarP(&manualYear, "year", "y", "", "year the manual was released")
}

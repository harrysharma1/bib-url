/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
  		title        = "{R}: A Language and Environment for Statistical Computing",
  		author       = "{R Core Team}",
  		organisation = "R Foundation for Statistical Computing",
  		address      = "Vienna, Austria",
  		year         = 2018
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("manual called")
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

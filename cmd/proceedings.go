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
	proceedingsEditors   []string
	proceedingsTitle     string
	proceedingsSeries    string
	proceedingsVolume    string
	proceedingsPublisher string
	proceedingsAddress   string
	proceedingsYear      string
)

// proceedingsCmd represents the proceedings command
var proceedingsCmd = &cobra.Command{
	Use:   "proceedings",
	Short: "Generate @proceedings bibtex entry",
	Long: `This will generate an @proceedings entry e.g.
	
	@proceedings{CitekeyProceedings,
  		editor    = "Susan Stepney and Sergey Verlan",
  		title     = "Proceedings of the 17th International Conference on Computation and Natural Computation, Fontainebleau, France",
  		series    = "Lecture Notes in Computer Science",
  		volume    = "10867",
  		publisher = "Springer",
  		address   = "Cham, Switzerland",
  		year      = 2018
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatProceedingsBibtex()

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
	proceedingsCmd.Flags().StringArrayVarP(&proceedingsEditors, "editors", "e", []string{}, "editor name(s) for proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsTitle, "title", "t", "", "title of proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsSeries, "series", "s", "", "series where proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedingsVolume, "volume", "v", "", "volume where proceedings was published")
	proceedingsCmd.Flags().StringVarP(&proceedingsPublisher, "publisher", "p", "", "the group that published proceedings")
	proceedingsCmd.Flags().StringVarP(&proceedingsAddress, "address", "l", "", "location of publisher")
	proceedingsCmd.Flags().StringVarP(&proceedingsYear, "year", "y", "", "year proceedings was published")
}

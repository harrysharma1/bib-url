/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inproceedingsCmd represents the inproceedings command
var inproceedingsCmd = &cobra.Command{
	Use:   "inproceedings",
	Short: "Generate @inproceedings bibtex entry",
	Long: `This will generate an @inproceedings bibtex entry e.g.
	
	@inproceedings{CitekeyInproceedings,
  		author    = "Holleis, Paul and Wagner, Matthias and Koolwaaij, Johan",
  		title     = "Studying mobile context-aware social services in the wild",
  		booktitle = "Proc. of the 6th Nordic Conf. on Human-Computer Interaction",
  		series    = "NordiCHI",
  		year      = 2010,
  		pages     = "207--216",
  		publisher = "ACM",
  		address   = "New York, NY"
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inproceedings called")
	},
}

func init() {
	rootCmd.AddCommand(inproceedingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inproceedingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inproceedingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

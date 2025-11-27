/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	proceedingsYear      int
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("proceedings called")
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
}

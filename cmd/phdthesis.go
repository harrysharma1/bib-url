/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	phdthesisUrl     string
	phdthesisCiteKey string
	phdthesisAuthors []string
	phdthesisTitle   string
	phdthesisSchool  string
	phdthesisAddress string
	phdthesisYear    int
	phdthesisMonth   string
)

// phdthesisCmd represents the phdthesis command
var phdthesisCmd = &cobra.Command{
	Use:   "phdthesis",
	Short: "Generate @phdthesis bibtex entry",
	Long: `This will generate an @phdthesis bibtex entry e.g.
	
	@phdthesis{CitekeyPhdthesis,
  		author  = "Rempel, Robert Charles",
  		title   = "Relaxation Effects for Coupled Nuclear Spins",
  		school  = "Stanford University",
  		address = "Stanford, CA",
  		year    = 1956,
  		month   = jun
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("phdthesis called")
	},
}

func init() {
	rootCmd.AddCommand(phdthesisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// phdthesisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// phdthesisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

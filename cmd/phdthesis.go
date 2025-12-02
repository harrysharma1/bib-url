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
	phdthesisYear    string
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
	phdthesisCmd.Flags().StringVarP(&phdthesisUrl, "url", "u", "", "url for online PhD thesis to auto-generate entry")
	phdthesisCmd.Flags().StringVarP(&phdthesisCiteKey, "key", "k", "", "citation key for bibtex entry")
	phdthesisCmd.Flags().StringArrayVarP(&phdthesisAuthors, "author", "a", []string{}, "author name(s) for PhD thesis")
	phdthesisCmd.Flags().StringVarP(&phdthesisTitle, "title", "t", "", "title of PhD thesis")
	phdthesisCmd.Flags().StringVarP(&phdthesisSchool, "school", "s", "", "school where PhD thesis was published")
	phdthesisCmd.Flags().StringVarP(&phdthesisAddress, "address", "l", "", "location of school")
	phdthesisCmd.Flags().StringVarP(&phdthesisYear, "year", "y", "", "year PhD thesis was published")
	phdthesisCmd.Flags().StringVarP(&phdthesisMonth, "month", "m", "", "month PhD thesis was published")
}

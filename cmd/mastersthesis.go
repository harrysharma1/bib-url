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
	mastersthesisUrl     string
	mastersthesisCiteKey string
	mastersthesisAuthors []string
	mastersthesisTitle   string
	mastersthesisSchool  string
	mastersthesisYear    string
	mastersthesisAddress string
	mastersthesisMonth   string
)

// mastersthesisCmd represents the mastersthesis command
var mastersthesisCmd = &cobra.Command{
	Use:   "mastersthesis",
	Short: "Generate @mastersthesis bibtex entry",
	Long: `This will generate an @mastersthesis bibtex entry e.g.
	
	@mastersthesis{CitekeyMastersthesis,
  		author  = "Jian Tang",
  		title   = "Spin structure of the nucleon in the asymptotic limit",
  		school  = "Massachusetts Institute of Technology",
  		year    = 1996,
  		address = "Cambridge, MA",
  		month   = sep
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatMastersThesisBibtex(mastersthesisCiteKey, mastersthesisAuthors, mastersthesisTitle, mastersthesisSchool, mastersthesisYear, mastersthesisAddress, mastersthesisMonth, braces)

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
	rootCmd.AddCommand(mastersthesisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mastersthesisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mastersthesisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisUrl, "url", "u", "", "url for online masters thesis to auto-generate entry")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisCiteKey, "key", "k", "", "citation key for bibtex entry")
	mastersthesisCmd.Flags().StringArrayVarP(&mastersthesisAuthors, "author", "a", []string{}, "author name(s) for masters thesis")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisTitle, "title", "t", "", "title of masters thesis")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisSchool, "school", "s", "", "school where the masters thesis came from")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisYear, "year", "y", "", "year the masters thesis was released")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisAddress, "address", "l", "", "address of academic institution")
	mastersthesisCmd.Flags().StringVarP(&mastersthesisMonth, "month", "m", "", "month the masters thesis was released")
}

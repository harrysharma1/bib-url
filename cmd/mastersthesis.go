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

var mastersthesis models.Mastersthesis

// mastersthesisCmd represents the mastersthesis command
var mastersthesisCmd = &cobra.Command{
	Use:   "mastersthesis",
	Short: "Generate @mastersthesis bibtex entry",
	Long: `This will generate an @mastersthesis bibtex entry e.g.
	
@mastersthesis{CitekeyMastersthesis,
	author  = "<Lastname, Firstname>",
	title   = "<Title>",
	school  = "<School>",
	year    = <2002>,
	type	= "<Type>",
	address = "<Address>",
	month   = <jul>,
	note	= "<Note>"
}

Required:
	- author
	- title
	- school
	- year
Optional:
	- type
	- address
	- month
	- note
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatMastersThesisBibtex(
			mastersthesis.CiteKey,
			mastersthesis.Authors,
			mastersthesis.Title,
			mastersthesis.School,
			mastersthesis.Year,
			mastersthesis.Type,
			mastersthesis.Address,
			mastersthesis.Month,
			mastersthesis.Note,
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
	rootCmd.AddCommand(mastersthesisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mastersthesisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mastersthesisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.CiteKey, "key", "k", "", "citation key for bibtex entry")
	mastersthesisCmd.Flags().StringArrayVarP(&mastersthesis.Authors, "author", "a", []string{}, "author name(s) for masters thesis")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Title, "title", "t", "", "title of masters thesis")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.School, "school", "s", "", "school where the masters thesis came from")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Year, "year", "y", "", "year the masters thesis was released")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Type, "type", "g", "", "type of masters thesis")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Address, "address", "l", "", "address of academic institution")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Month, "month", "m", "", "month the masters thesis was released")
	mastersthesisCmd.Flags().StringVarP(&mastersthesis.Note, "note", "n", "", "additional notes for masters thesis")
}

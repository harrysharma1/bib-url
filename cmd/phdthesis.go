/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/helper"
	"bibcli/models"
	"fmt"

	"github.com/spf13/cobra"
)

var phdthesis models.Phdthesis

// phdthesisCmd represents the phdthesis command
var phdthesisCmd = &cobra.Command{
	Use:   "phdthesis",
	Short: "Generate @phdthesis bibtex entry",
	Long: `This will generate an @phdthesis bibtex entry e.g.
	
@phdthesis{CitekeyPhdthesis,
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
Optional
	- type
	- address
	- month
	- note
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatPhDThesisBibtex(
			phdthesis.CiteKey,
			phdthesis.Authors,
			phdthesis.Title,
			phdthesis.School,
			phdthesis.Year,
			phdthesis.Type,
			phdthesis.Address,
			phdthesis.Month,
			phdthesis.Note,
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
	rootCmd.AddCommand(phdthesisCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// phdthesisCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// phdthesisCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	phdthesisCmd.Flags().StringVarP(&phdthesis.CiteKey, "key", "k", "", "citation key for bibtex entry")
	phdthesisCmd.Flags().StringArrayVarP(&phdthesis.Authors, "author", "a", []string{}, "author name(s) for PhD thesis")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Title, "title", "t", "", "title of PhD thesis")
	phdthesisCmd.Flags().StringVarP(&phdthesis.School, "school", "s", "", "school where PhD thesis was published")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Year, "year", "y", "", "year PhD thesis was published")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Type, "type", "g", "", "type of PhD thesis")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Address, "address", "l", "", "location of school")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Month, "month", "m", "", "month PhD thesis was published")
	phdthesisCmd.Flags().StringVarP(&phdthesis.Note, "note", "n", "", "additional notes for PhD thesis")
}

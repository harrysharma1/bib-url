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

var techreport models.Techreport

// techreportCmd represents the techreport command
var techreportCmd = &cobra.Command{
	Use:   "techreport",
	Short: "Generate @techreport bibtex entry",
	Long: `This will generate an @techreport bibtex entry e.g.
	
@techreport{CitekeyTechreport,
	title       = "<Title>",
	author      = "<Lastname, Firstname>",
	institution = "<Institution>",
	address     = "<Address>",
	number      = "<50>",
	year        = <2002>,
	month       = <jul>
}

Required:
	- title
	- author
	- institution
	- address
	- number
	- year
	- month
Optional:
	- !*
<>: indicates that it is a example value and should be changed.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = format.FormatTechReportBibtex(
			techreport.CiteKey,
			techreport.Title,
			techreport.Authors,
			techreport.Institution,
			techreport.Address,
			techreport.Number,
			techreport.Year,
			techreport.Month,
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
	rootCmd.AddCommand(techreportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// techreportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// techreportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	techreportCmd.Flags().StringVarP(&techreport.CiteKey, "key", "k", "", "citation key for bibtex entry")
	techreportCmd.Flags().StringVarP(&techreport.Title, "title", "t", "", "title of tech report")
	techreportCmd.Flags().StringArrayVarP(&techreport.Authors, "author", "a", []string{}, "author name(s) for tech report")
	techreportCmd.Flags().StringVarP(&techreport.Institution, "institution", "i", "", "institution that publishes tech report")
	techreportCmd.Flags().StringVarP(&techreport.Address, "address", "l", "", "address of institution")
	techreportCmd.Flags().StringVarP(&techreport.Number, "number", "n", "", "tech report identifier")
	techreportCmd.Flags().StringVarP(&techreport.Year, "year", "y", "", "year the tech report was released")
	techreportCmd.Flags().StringVarP(&techreport.Month, "month", "m", "", "month the tech report was released")
}

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
	techreportUrl         string
	techreportCiteKey     string
	techreportTitle       string
	techreportAuthors     []string
	techreportInstitution string
	techreportAddress     string
	techreportNumber      string
	techreportYear        string
	techreportMonth       string
)

// techreportCmd represents the techreport command
var techreportCmd = &cobra.Command{
	Use:   "techreport",
	Short: "Generate @techreport bibtex entry",
	Long: `This will generate an @techreport bibtex entry e.g.
	
@techreport{CitekeyTechreport,
	title       = "<Title>",
	author      = "<Lastname, Firstname>",
	institution = "<Institution",
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
		var bibtex = helper.FormatTechReportBibtex(
			techreportCiteKey,
			techreportTitle,
			techreportAuthors,
			techreportInstitution,
			techreportAddress,
			techreportNumber,
			techreportYear,
			techreportMonth,
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
	techreportCmd.Flags().StringVarP(&techreportUrl, "url", "u", "", "url for tech report to auto-generate entry")
	techreportCmd.Flags().StringVarP(&techreportCiteKey, "key", "k", "", "citation key for bibtex entry")
	techreportCmd.Flags().StringVarP(&techreportTitle, "title", "t", "", "title of tech report")
	techreportCmd.Flags().StringArrayVarP(&techreportAuthors, "author", "a", []string{}, "author name(s) for tech report")
	techreportCmd.Flags().StringVarP(&techreportInstitution, "institution", "i", "", "institution that publishes tech report")
	techreportCmd.Flags().StringVarP(&techreportAddress, "address", "l", "", "address of institution")
	techreportCmd.Flags().StringVarP(&techreportNumber, "number", "n", "", "tech report identifier")
	techreportCmd.Flags().StringVarP(&techreportYear, "year", "y", "", "year the tech report was released")
	techreportCmd.Flags().StringVarP(&techreportMonth, "month", "m", "", "month the tech report was released")
}

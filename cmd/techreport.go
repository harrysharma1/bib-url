/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	techreportYear        int
	techreportMonth       string
)

// techreportCmd represents the techreport command
var techreportCmd = &cobra.Command{
	Use:   "techreport",
	Short: "Generate @techreport bibtex entry",
	Long: `This will generate an @techreport bibtex entry e.g.
	
	@techreport{CitekeyTechreport,
  		title       = "{W}asatch {S}olar {P}roject Final Report",
  		author      = "Bennett, Vicki and Bowman, Kate and Wright, Sarah",
  		institution = "Salt Lake City Corporation",
  		address     = "Salt Lake City, UT",
  		number      = "DOE-SLC-6903-1",
  		year        = 2018,
  		month       = sep
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("techreport called")
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
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	incollectionUrl       string
	incollectionCiteKey   string
	incollectionAuthors   []string
	incollectionEditor    string
	incollectionTitle     string
	incollectionBookTitle string
	incollectionYear      int
	incollectionPublisher string
	incollectionAddress   string
	incollectionPages     string
)

// incollectionCmd represents the incollection command
var incollectionCmd = &cobra.Command{
	Use:   "incollection",
	Short: "Generate @incollection bibtex entry",
	Long: `This will generate an @incollection bibtex entry e.g.
	
	@incollection{CitekeyIncollection,
  		author    = "Shapiro, Howard M.",
  		editor    = "Hawley, Teresa S. and Hawley, Robert G.",
  		title     = "Flow Cytometry: The Glass Is Half Full",
  		booktitle = "Flow Cytometry Protocols",
  		year      = 2018,
 	 	publisher = "Springer",
  		address   = "New York, NY",
  		pages     = "1--10"
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incollection called")
	},
}

func init() {
	rootCmd.AddCommand(incollectionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incollectionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incollectionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

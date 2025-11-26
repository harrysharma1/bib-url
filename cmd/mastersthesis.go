/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mastersthesis called")
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
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// manualCmd represents the manual command
var manualCmd = &cobra.Command{
	Use:   "manual",
	Short: "Generate @manual bibtex entry",
	Long: `This will generate an @manual bibtex entry e.g.
	
	@manual{CitekeyManual,
  		title        = "{R}: A Language and Environment for Statistical Computing",
  		author       = "{R Core Team}",
  		organization = "R Foundation for Statistical Computing",
  		address      = "Vienna, Austria",
  		year         = 2018
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("manual called")
	},
}

func init() {
	rootCmd.AddCommand(manualCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manualCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manualCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

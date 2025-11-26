/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unpublishedCmd represents the unpublished command
var unpublishedCmd = &cobra.Command{
	Use:   "unpublished",
	Short: "Generate @unpublished bibtex entry",
	Long: `This will generate an @unpublished bibtex entry e.g.
	
	@unpublished{CitekeyUnpublished,
  		author = "Mohinder Suresh",
  		title  = "Evolution: a revised theory",
 		year   = 2006
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unpublished called")
	},
}

func init() {
	rootCmd.AddCommand(unpublishedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unpublishedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unpublishedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

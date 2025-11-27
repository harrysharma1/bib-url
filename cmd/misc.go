/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	miscUrl          string
	miscCiteKey      string
	miscAuthors      []string
	miscHowPublished string
	miscNote         string
)

// miscCmd represents the misc command
var miscCmd = &cobra.Command{
	Use:   "misc",
	Short: "Generate @misc bibtex entry",
	Long: `This will generate an @misc bibtex entry e.g.
	
	@misc{CitekeyMisc,
  		title        = "Pluto: The 'Other' Red Planet",
  		author       = "{NASA}",
  		howpublished = "\url{https://www.nasa.gov/nh/pluto-the-other-red-planet}",
  		year         = 2015,
  		note         = "Accessed: 2018-12-06"
	}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("misc called")
	},
}

func init() {
	rootCmd.AddCommand(miscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// miscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// miscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

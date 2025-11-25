/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bookletUrl          string
	bookletCiteKey      string
	bookletTitle        string
	bookletAuthor       string
	bookletHowPublished string
	bookletMonth        string
	bookletYear         int
)

// bookletCmd represents the booklet command
var bookletCmd = &cobra.Command{
	Use:   "booklet",
	Short: "Generate @booklet bibtex entry",
	Long: `This will generate an @booklet bibtex entry e.g.
	
	@booklet{CitekeyBooklet,
 		title        = "Canoe tours in {S}weden",
  		author       = "Maria Swetla", 
  		howpublished = "Distributed at the Stockholm Tourist Office",
  		month        = jul,
  		year         = 2015
}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("booklet called")
	},
}

func init() {
	rootCmd.AddCommand(bookletCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bookletCmd.Flags().StringVarP(&bookletUrl, "url", "u", "https://example.com", "url for online booklet to auto-generate entry")
	bookletCmd.Flags().StringVarP(&bookletCiteKey, "key", "k", "uuid.uuid4()", "citation key for bibtex entry")
	bookletCmd.Flags().StringVarP(&bookletTitle, "title", "t", "Title", "title of booklet")
	bookletCmd.Flags().StringVarP(&bookletAuthor, "author", "a", "Surname, Forname", "author name(s) for book")
	bookletCmd.Flags().StringVarP(&bookletHowPublished, "published", "p", "Distributed at the Stockholm Tourist Office", "how the booklet was published")
	bookletCmd.Flags().StringVarP(&bookletMonth, "month", "m", "jul", "month the booklet was released")
	bookletCmd.Flags().IntVarP(&bookletYear, "year", "y", 2002, "year the booklet was released")
}

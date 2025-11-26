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
	bookUrl       string
	bookCiteKey   string
	bookAuthors   []string
	bookTitle     string
	bookPublisher string
	bookAddress   string
	bookYear      int
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Generate @book bibtex entry",
	Long: `This will generate an @book bibtex entry e.g.
	
	@book{CitekeyBook,
  		author    = "Leonard Susskind and George Hrabovsky",
  		title     = "Classical mechanics: the theoretical minimum",
  		publisher = "Penguin Random House",
  		address   = "New York, NY",
  		year      = 2014
	}
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var bibtex = helper.FormatBookBibtex(bookCiteKey, bookAuthors, bookTitle, bookPublisher, bookAddress, bookYear)

		if copy {
			helper.Copy(bibtex)
		}

		if save != "" {
			if err := helper.Save(save, bibtex); err != nil {
				return err
			}
		}

		fmt.Println(bibtex)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bookCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bookCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bookCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	bookCmd.Flags().StringVarP(&bookUrl, "url", "u", "https://example.com", "url for online book to auto-generate entry")
	bookCmd.Flags().StringVarP(&bookCiteKey, "key", "k", "uuid.uuid4()", "citation key for bibtex entry")
	bookCmd.Flags().StringArrayVarP(&bookAuthors, "author", "a", []string{"Leonard Susskind", "George Hrabovsky"}, "author name(s) for book")
	bookCmd.Flags().StringVarP(&bookTitle, "title", "t", "Title", "title of book")
	bookCmd.Flags().StringVarP(&bookPublisher, "publisher", "p", "Penguin Random House", "publisher that released the book")
	bookCmd.Flags().StringVarP(&bookAddress, "address", "l", "New York, NY", "address of publisher")
	bookCmd.Flags().IntVarP(&bookYear, "year", "y", 2002, "year the book was published")
}

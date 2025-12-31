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
	bookISBN      string
	bookCiteKey   string
	bookAuthors   []string
	bookTitle     string
	bookPublisher string
	bookAddress   string
	bookYear      string
)

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Generate @book bibtex entry",
	Long: `This will generate an @book bibtex entry e.g.
	
@book{CitekeyBook,
	author    = "<Lastname, Firstname>",
	title     = "<Title>",
	publisher = "<Publisher>",
	address   = "<Address>",
	year      = <2002>
}

Required:
	- author
	- title
	- publisher
	- address
	- year
<>: indicates that it is a example value and should be changed.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if bookISBN != "" {
			isbnAuthors, isbnTitle, isbnPublisher, isbnYear, err := helper.BookFromISBN(bookISBN)
			if err != nil {
				return err
			}

			if len(bookAuthors) == 0 {
				bookAuthors = isbnAuthors
			} else {
				bookAuthors = append(bookAuthors, isbnAuthors...)
			}

			if bookTitle == "" {
				bookTitle = isbnTitle
			}

			if bookPublisher == "" {
				bookPublisher = isbnPublisher
			}

			if bookYear == "" {
				bookYear = isbnYear
			}

		}
		var bibtex = helper.FormatBookBibtex(bookCiteKey, bookAuthors, bookTitle, bookPublisher, bookAddress, bookYear, braces)

		if copy {
			helper.Copy(bibtex)
		}

		if save != "" {
			if err = helper.Save(save, bibtex); err != nil {
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
	bookCmd.Flags().StringVarP(&bookISBN, "isbn", "i", "", "isbn for online book to auto-generate entry")
	bookCmd.Flags().StringVarP(&bookCiteKey, "key", "k", "", "citation key for bibtex entry")
	bookCmd.Flags().StringArrayVarP(&bookAuthors, "author", "a", []string{}, "author name(s) for book")
	bookCmd.Flags().StringVarP(&bookTitle, "title", "t", "", "title of book")
	bookCmd.Flags().StringVarP(&bookPublisher, "publisher", "p", "", "publisher that released the book")
	bookCmd.Flags().StringVarP(&bookAddress, "address", "l", "", "address of publisher")
	bookCmd.Flags().StringVarP(&bookYear, "year", "y", "", "year the book was published")

}

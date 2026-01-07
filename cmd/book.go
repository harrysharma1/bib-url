/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bibcli/format"
	"bibcli/helper"
	"bibcli/models"
	"fmt"

	"github.com/spf13/cobra"
)

var book models.Book

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

		if cmd.Flags().Changed("isbn") {
			isbnAuthors, isbnTitle, isbnPublisher, isbnYear, err := format.BookFromISBN(book.ISBN)
			if err != nil {
				return err
			}

			if len(book.Authors) == 0 {
				book.Authors = isbnAuthors
			} else {
				book.Authors = append(book.Authors, isbnAuthors...)
			}

			if book.Title == "" {
				book.Title = isbnTitle
			}

			if book.Publisher == "" {
				book.Publisher = isbnPublisher
			}

			if book.Year == "" {
				book.Year = isbnYear
			}

		}
		var bibtex = format.FormatBookBibtex(
			book.CiteKey,
			book.Authors,
			book.Title,
			book.Publisher,
			book.Address,
			book.Year,
			braces)

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
	bookCmd.Flags().StringVarP(&book.ISBN, "isbn", "i", "", "isbn for online book to auto-generate entry")
	bookCmd.Flags().StringVarP(&book.CiteKey, "key", "k", "", "citation key for bibtex entry")
	bookCmd.Flags().StringArrayVarP(&book.Authors, "author", "a", []string{}, "author name(s) for book")
	bookCmd.Flags().StringVarP(&book.Title, "title", "t", "", "title of book")
	bookCmd.Flags().StringVarP(&book.Publisher, "publisher", "p", "", "publisher that released the book")
	bookCmd.Flags().StringVarP(&book.Address, "address", "l", "", "address of publisher")
	bookCmd.Flags().StringVarP(&book.Year, "year", "y", "", "year the book was published")

}

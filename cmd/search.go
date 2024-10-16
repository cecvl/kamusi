package cmd

import (
	"fmt"
	"kamusicli/pkg/db"

	"github.com/spf13/cobra"
)

var word string

// Define the search command
var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search for a word's meaning and synonyms",
    Run: func(cmd *cobra.Command, args []string) {
        if word == "" {
            fmt.Println("Please provide a word using the --word flag.")
            return
        }

        meaning, synonyms, err := db.SearchWord(word)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            return
        }

        fmt.Printf("Meaning: %s\n", meaning)
        fmt.Printf("Synonyms: %v\n", synonyms)
    },
}

// Initialize the command flags
func init() {
    searchCmd.Flags().StringVarP(&word, "word", "w", "", "The word to search for in the dictionary")
    rootCmd.AddCommand(searchCmd)
}

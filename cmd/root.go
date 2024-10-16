package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "kms",
    Short: "A command-line tool to search for words in a dictionary",
    Long:  `This tool allows you to search for the meaning and synonyms of a word in an SQLite database.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

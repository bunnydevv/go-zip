package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Display the current version of go-zip.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-zip v%s\n", version)
		fmt.Println("A powerful CLI tool for file compression and decompression")
	},
}
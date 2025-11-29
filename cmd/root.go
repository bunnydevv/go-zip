package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-zip",
	Short: "A powerful CLI tool for file compression and decompression",
	Long: `go-zip is a comprehensive command-line tool for handling various compression formats.
Similar to 7zip, it supports multiple compression algorithms including ZIP, GZIP, TAR, and more.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(compressCmd)
	rootCmd.AddCommand(decompressCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(versionCmd)
}
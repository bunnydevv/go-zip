package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bunnydevv/go-zip/pkg/compression"
	"github.com/spf13/cobra"
)

var (
	extractPath     string
	extractPassword string
)

var decompressCmd = &cobra.Command{
	Use:     "decompress [archive]",
	Short:   "Decompress an archive file",
	Long:    `Extract files from compressed archives.`,
	Aliases: []string{"d", "extract", "x"},
	Args:    cobra.ExactArgs(1),
	RunE:    runDecompress,
}

func init() {
	decompressCmd.Flags().StringVarP(&extractPath, "output", "o", ".", "Output directory for extracted files")
	decompressCmd.Flags().StringVarP(&extractPassword, "password", "p", "", "Password for encrypted archives")
}

func runDecompress(cmd *cobra.Command, args []string) error {
	archivePath := args[0]

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return fmt.Errorf("archive file not found: %s", archivePath)
	}

	fmt.Printf("Extracting %s to %s...\n", archivePath, extractPath)

	archiveType := detectArchiveType(archivePath)
	var err error

	switch archiveType {
	case "zip":
		err = compression.ExtractZip(archivePath, extractPath, extractPassword)
	case "tar":
		err = compression.ExtractTar(archivePath, extractPath)
	case "tar.gz", "tgz":
		err = compression.ExtractTarGz(archivePath, extractPath)
	case "tar.bz2", "tbz2":
		err = compression.ExtractTarBz2(archivePath, extractPath)
	case "gzip", "gz":
		err = compression.ExtractGzip(archivePath, extractPath)
	default:
		return fmt.Errorf("unsupported or unknown archive type: %s", archivePath)
	}

	if err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	fmt.Printf("✓ Successfully extracted to %s\n", extractPath)
	return nil
}

func detectArchiveType(filename string) string {
	lower := strings.ToLower(filename)
	
	if strings.HasSuffix(lower, ".tar.gz") || strings.HasSuffix(lower, ".tgz") {
		return "tar.gz"
	}
	if strings.HasSuffix(lower, ".tar.bz2") || strings.HasSuffix(lower, ".tbz2") {
		return "tar.bz2"
	}
	if strings.HasSuffix(lower, ".tar") {
		return "tar"
	}
	if strings.HasSuffix(lower, ".zip") {
		return "zip"
	}
	if strings.HasSuffix(lower, ".gz") {
		return "gzip"
	}
	
	return "unknown"
}
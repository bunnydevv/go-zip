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
	compressionLevel int
	compressionType  string
	outputFile       string
	password         string
)

var compressCmd = &cobra.Command{
	Use:   "compress [files...]",
	Short: "Compress files or directories",
	Long:  `Compress one or more files or directories into an archive.`,
	Aliases: []string{"c", "add"},
	Args:  cobra.MinimumNArgs(1),
	RunE:  runCompress,
}

func init() {
	compressCmd.Flags().IntVarP(&compressionLevel, "level", "l", 6, "Compression level (0-9)")
	compressCmd.Flags().StringVarP(&compressionType, "type", "t", "zip", "Compression type (zip, tar, tar.gz, tar.bz2, gzip)")
	compressCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output archive file name")
	compressCmd.Flags().StringVarP(&password, "password", "p", "", "Password for encryption (zip only)")
}

func runCompress(cmd *cobra.Command, args []string) error {
	if outputFile == "" {
		outputFile = generateOutputFileName(args, compressionType)
	}

	fmt.Printf("Compressing to %s...\n", outputFile)

	var err error
	switch strings.ToLower(compressionType) {
	case "zip":
		err = compression.CreateZip(args, outputFile, compressionLevel, password)
	case "tar":
		err = compression.CreateTar(args, outputFile)
	case "tar.gz", "tgz":
		err = compression.CreateTarGz(args, outputFile, compressionLevel)
	case "tar.bz2", "tbz2":
		err = compression.CreateTarBz2(args, outputFile, compressionLevel)
	case "gzip", "gz":
		if len(args) > 1 {
			return fmt.Errorf("gzip can only compress a single file")
		}
		err = compression.CreateGzip(args[0], outputFile, compressionLevel)
	default:
		return fmt.Errorf("unsupported compression type: %s", compressionType)
	}

	if err != nil {
		return fmt.Errorf("compression failed: %w", err)
	}

	info, _ := os.Stat(outputFile)
	fmt.Printf("✓ Successfully created %s (%.2f MB)\n", outputFile, float64(info.Size())/(1024*1024))
	return nil
}

func generateOutputFileName(inputs []string, compType string) string {
	if len(inputs) == 1 {
		base := filepath.Base(inputs[0])
		base = strings.TrimSuffix(base, filepath.Ext(base))
		return base + "." + compType
	}
	return "archive." + compType
}
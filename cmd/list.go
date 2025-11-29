package cmd

import (
	"fmt"

	"github.com/bunnydevv/go-zip/pkg/compression"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [archive]",
	Short:   "List contents of an archive",
	Long:    `Display the contents of a compressed archive without extracting it.`,
	Aliases: []string{"l", "ls"},
	Args:    cobra.ExactArgs(1),
	RunE:    runList,
}

func runList(cmd *cobra.Command, args []string) error {
	archivePath := args[0]
	archiveType := detectArchiveType(archivePath)

	fmt.Printf("Contents of %s:\n\n", archivePath)

	var err error
	switch archiveType {
	case "zip":
		err = compression.ListZip(archivePath)
	case "tar", "tar.gz", "tgz", "tar.bz2", "tbz2":
		err = compression.ListTar(archivePath, archiveType)
	default:
		return fmt.Errorf("listing not supported for archive type: %s", archiveType)
	}

	return err
}
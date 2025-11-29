package compression

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateGzip(source, target string, level int) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()


targetFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	gzWriter, err := gzip.NewWriterLevel(targetFile, level)
	if err != nil {
		return err
	}
	defer gzWriter.Close()

	gzWriter.Name = filepath.Base(source)

	_, err = io.Copy(gzWriter, sourceFile)
	return err
}

func ExtractGzip(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	gzReader, err := gzip.NewReader(sourceFile)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	outputName := strings.TrimSuffix(filepath.Base(source), ".gz")
	if gzReader.Name != "" {
		outputName = gzReader.Name
	}

	targetPath := filepath.Join(destination, outputName)
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, gzReader)
	return err
}
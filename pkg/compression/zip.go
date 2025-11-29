package compression

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateZip(sources []string, target string, level int, password string) error {
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	for _, source := range sources {
		err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = strings.TrimPrefix(path, filepath.Dir(source)+string(os.PathSeparator))
			header.Method = zip.Deflate

			if info.IsDir() {
				header.Name += "/"
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()
				_, err = io.Copy(writer, file)
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func ExtractZip(source, destination, password string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		path := filepath.Join(destination, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}

		destFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		fileReader, err := file.Open()
		if err != nil {
			destFile.Close()
			return err
		}

		_, err = io.Copy(destFile, fileReader)
		destFile.Close()
		fileReader.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func ListZip(source string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	fmt.Printf("%-10s %-15s %s\n", "Size", "Modified", "Name")
	fmt.Println(strings.Repeat("-", 70))

	for _, file := range reader.File {
		fmt.Printf("%-10d %-15s %s\n", 
			file.UncompressedSize64, 
			file.Modified.Format("2006-01-02 15:04"),
			file.Name)
	}

	return nil
}
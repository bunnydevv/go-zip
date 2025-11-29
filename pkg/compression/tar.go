package compression

import (
	"archive/tar"
	"compress/bzip2"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateTar(sources []string, target string) error {
	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	return createTarArchive(sources, tarFile)
}

func CreateTarGz(sources []string, target string, level int) error {
	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzWriter, err := gzip.NewWriterLevel(tarFile, level)
	if err != nil {
		return err
	}
	defer gzWriter.Close()

	return createTarArchive(sources, gzWriter)
}

func CreateTarBz2(sources []string, target string, level int) error {
	return fmt.Errorf("bzip2 compression not yet implemented")
}

func createTarArchive(sources []string, writer io.Writer) error {
	tarWriter := tar.NewWriter(writer)
	defer tarWriter.Close()

	for _, source := range sources {
		err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			header.Name = strings.TrimPrefix(path, filepath.Dir(source)+string(os.PathSeparator))

			if err := tarWriter.WriteHeader(header); err != nil {
				return err
			}

			if !info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = io.Copy(tarWriter, file)
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

func ExtractTar(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	return extractTarArchive(file, destination)
}

func ExtractTarGz(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzReader.Close()

	return extractTarArchive(gzReader, destination)
}

func ExtractTarBz2(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	bzReader := bzip2.NewReader(file)
	return extractTarArchive(bzReader, destination)
}

func extractTarArchive(reader io.Reader, destination string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, tarReader)
			outFile.Close()
			if err != nil {
				return err
			}
		}
		}

	return nil
}

func ListTar(source, archiveType string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	var reader io.Reader = file

	if strings.Contains(archiveType, "gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		defer gzReader.Close()
		reader = gzReader
	} else if strings.Contains(archiveType, "bz2") {
		reader = bzip2.NewReader(file)
	}

	tarReader := tar.NewReader(reader)

	fmt.Printf("%-10s %-15s %s\n", "Size", "Modified", "Name")
	fmt.Println(strings.Repeat("-", 70))

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Printf("%-10d %-15s %s\n",
			header.Size,
			header.ModTime.Format("2006-01-02 15:04"),
			header.Name)
	}

	return nil
}
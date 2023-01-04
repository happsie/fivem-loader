package updater

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(resourceTargetDirectory, scriptName, zipName string) error {
	archive, err := zip.OpenReader(zipName)
	if err != nil {
		return err
	}
	defer archive.Close()

	folderName := ""
	for _, f := range archive.File {
		if f.FileInfo().IsDir() && folderName == "" && strings.HasSuffix(f.Name, "-master/") {
			folderName = f.Name
		}
		filePath := filepath.Join(resourceTargetDirectory, f.Name)
		fmt.Printf("unzipping file %s\n", filePath)

		if f.FileInfo().IsDir() {
			fmt.Printf("creating directory %s", filePath)
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	if folderName != "" {
		err = os.Rename(filepath.Join(resourceTargetDirectory, folderName), filepath.Join(resourceTargetDirectory, scriptName))
		if err != nil {
			return err
		}
	}
	return nil
}

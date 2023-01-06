package updater

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/happsie/fivem-loader/internal"
)

type Unzipped struct {
	scripts []string
}

func Unzip(resourceTargetDirectory, scriptName, zipName string) (Unzipped, error) {
	archive, err := zip.OpenReader(zipName)
	if err != nil {
		return Unzipped{}, err
	}
	defer archive.Close()

	folderName := strings.ReplaceAll(zipName, ".zip", "")
	fmt.Printf(internal.InfoColor, fmt.Sprintf("Installing script [%s] to %s\n", scriptName, resourceTargetDirectory))

	var directories []string
	for _, f := range archive.File {
		filePath := filepath.Join(resourceTargetDirectory, f.Name)

		if f.FileInfo().IsDir() {
			directories = append(directories, f.FileInfo().Name())
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return Unzipped{}, err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return Unzipped{}, err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return Unzipped{}, err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return Unzipped{}, err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return Unzipped{}, err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	unzipped := Unzipped{}
	if scriptName != "" && len(directories) == 1 {
		err = os.RemoveAll(filepath.Join(resourceTargetDirectory, scriptName))
		if err != nil {
			return Unzipped{}, err
		}
		err = os.Rename(filepath.Join(resourceTargetDirectory, folderName), filepath.Join(resourceTargetDirectory, scriptName))
		if err != nil {
			return Unzipped{}, err
		}
		unzipped.scripts = append(unzipped.scripts, scriptName)
	}
	// assume multiscript
	if len(directories) > 1 {
		for _, dir := range directories {
			// if it's the parent directory, lets continue to the next directory
			if dir == folderName {
				continue
			}
			err = os.RemoveAll(filepath.Join(resourceTargetDirectory, dir))
			if err != nil {
				return unzipped, err
			}
			err := os.Rename(filepath.Join(resourceTargetDirectory, folderName, dir), filepath.Join(resourceTargetDirectory, dir))
			if err != nil {
				return unzipped, err
			}
			unzipped.scripts = append(unzipped.scripts, dir)
		}
		err = os.RemoveAll(filepath.Join(resourceTargetDirectory, folderName))
		if err != nil {
			return unzipped, err
		}
	}
	return unzipped, nil
}

package updater

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"

	"github.com/happsie/fivem-loader/internal"
)

func DownloadZip(url string) (zipName string, err error) {
	resp, err := downloadFromSource(url)
	if err != nil {
		return "", err
	}
	fmt.Printf(internal.InfoColor, fmt.Sprintf("Download of resource [%s] complete\n", url))
	defer resp.Body.Close()
	_, params, err := mime.ParseMediaType(resp.Header.Get("content-disposition"))
	if err != nil {
		return "", err
	}
	fileName := params["filename"]
	out, err := createZipFile(fileName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	err = copyToFile(out, resp.Body)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func RemoveFile(fileName string) error {
	return os.Remove(fileName)
}

func downloadFromSource(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func createZipFile(fileName string) (file *os.File, err error) {
	return os.Create(fileName)
}

func copyToFile(file *os.File, body io.ReadCloser) error {
	w, err := io.Copy(file, body)
	if err != nil {
		return err
	}
	if w == 0 {
		return fmt.Errorf("no bytes written to disk")
	}
	return nil
}

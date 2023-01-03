package updater

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func DownloadZip(url string) (zipName string, err error) {
	resp, err := downloadFromSource(url)
	if err != nil {
		return "", err
	}
	log.Printf("Download complete [%s]\n", url)
	defer resp.Body.Close()
	fileName := getFileName()
	out, err := createZipFile(fileName)
	if err != nil {
		return "", err
	}
	log.Printf("Zipfile created [%s]\n", fileName)
	defer out.Close()
	err = copyToFile(out, resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("Content written to file [%s]\n", fileName)
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

func getFileName() string {
	randomNames := []string{"panda", "bear", "duck", "lion", "dog", "cat"}
	rand.Seed(time.Now().Unix())
	return fmt.Sprintf("%s.zip", randomNames[rand.Intn(len(randomNames)-0)+0])
}

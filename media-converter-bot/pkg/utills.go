package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		cleanUpTmp()
		panic(err)
	}
}

func readApiToken() (string, error) {
	f, err := os.Open("hidden/token")
	if err != nil {
		return "", err
	}
	defer f.Close()

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(f)
	token := fileBuffer.String()
	return token, nil
}

func downloadFile(url string, dest string) (err error) {
	out, err := os.Create(dest)
	checkError(err)
	defer out.Close()

	res, err := http.Get(url)
	checkError(err)
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	checkError(err)
	return nil
}

func isMimeTypeVideo(mimeType string) bool {
	return strings.Split(mimeType, "/")[0] == "video"
}

func cleanUpTmp() {
	tmpExist, err := exists("tmp/")
	checkError(err)

	if tmpExist {
		err := os.RemoveAll("tmp/")
		checkError(err)
	}

	err = os.Mkdir("tmp/", 0755)
	checkError(err)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

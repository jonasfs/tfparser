package parser

import (
	"crypto/sha256"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func HashFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return h.Sum(nil), err
}

func GetFiles(dirPath string) (filePaths []string, err error) {
	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		return
	}
	files, err2 := ioutil.ReadDir(dirPath)
	if err2 != nil {
		err = err2
		return
	}

	for _, file := range files {
		fullPath := filepath.Join(dirPath, file.Name())
		filePaths = append(filePaths, fullPath)
	}
	return
}

func GetParserVersion() int {
	return TP.version
}

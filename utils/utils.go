package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
)

// CreateOutputFile generate a output file
func CreateOutputFile(outputPath string, outputPrefix string, outputFileExtension string) (outputAbosulePath string) {
	buff := make([]byte, int(math.Round(float64(32)/2)))
	rand.Read(buff)
	randomStr := hex.EncodeToString(buff)
	out := outputPath + string(os.PathSeparator) + outputPrefix + randomStr + outputFileExtension
	os.OpenFile(out, os.O_RDONLY|os.O_CREATE, 0666)
	log.Println("Output file: ", out)
	return out
}

// ReadFile read files from absolute path
func ReadFile(inputFile string) (data []byte) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println("Error reading file: ", err)
	}
	return data
}

// WriteToFile writes a buffer of bytes in a file
func WriteToFile(fileContent []byte, outputFile string) {
	file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s", fileContent)
}

// FilePathWalkDir get absolute path of the files in a folder
func FilePathWalkDir(folder string) ([]string, error) {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Println("Error reading inputpath ", err)
	}
	return files, err
}

// FileHasContent check if a file has content or is empty
func FileHasContent(filePath string) (r bool) {
	f, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
	}
	size := f.Size()
	if size != 0 {
		return true
	}
	return false
}

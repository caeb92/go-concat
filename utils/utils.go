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

// CreateOutputFile genera un archivo y retorna el path absoluto
func CreateOutputFile(outputPath string, outputPrefix string) (outputAbosulePath string) {
	buff := make([]byte, int(math.Round(float64(32)/2)))
	rand.Read(buff)
	randomStr := hex.EncodeToString(buff)
	out := outputPath + "/" + outputPrefix + randomStr + ".txt"
	os.OpenFile(out, os.O_RDONLY|os.O_CREATE, 0666)
	log.Println("output file: ", out)
	return out
}

// ReadFile lee un archivo del filesystem, el input es el path absoluto
func ReadFile(inputFile string) (data []byte) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Println("Error reading file: ", err)
	}
	return data
}

// WriteToFile hace append de texto a un archivo existente
func WriteToFile(fileContent []byte, outputFile string) {
	file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s", fileContent)
}

// FilePathWalkDir obtiene los path absolutos de los archivos almacenados en el path root
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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

// FileHasContent revisa si el peso del archivo es diferente de 0
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

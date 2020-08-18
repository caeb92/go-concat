package main

import (
	"bytes"
	"go-concat/utils"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// Read app CLI arguments
	var inputFolder, outputPath, outputPrefix, fileExtesion string
	var docsCount, c int

	inputFolder = os.Args[1]
	outputPath = os.Args[2]
	docsCount, _ = strconv.Atoi(os.Args[3])
	outputPrefix = os.Args[4]
	fileExtesion = os.Args[4]

	// Config logs
	logFile, err := os.OpenFile("go-concat.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.Println("PROCESS START")

	if len(os.Args[1:]) != 4 {
		log.Println("Some CLI arguments are missging")
	}

	log.Println("CLI arguments: ", os.Args[1:])

	// Read files in inputpath
	files, err := utils.FilePathWalkDir(inputFolder)
	log.Println("Docs count inputpath: ", len(files))

	// Create the first output file
	var outputAbosulePath string
	outputAbosulePath = utils.CreateOutputFile(outputPath, outputPrefix, fileExtesion)

	// Bytes buffer
	b := new(bytes.Buffer)
	for _, f := range files {
		if c == docsCount {
			c = 0
			utils.WriteToFile(b.Bytes(), outputAbosulePath)
			outputAbosulePath = utils.CreateOutputFile(outputPath, outputPrefix, fileExtesion)
			b.Reset()
		}
		c++
		b.Write(utils.ReadFile(f))
	}

	if !utils.FileHasContent(outputAbosulePath) {
		utils.WriteToFile(b.Bytes(), outputAbosulePath)
	}

	log.Println("PROCESS END")
}

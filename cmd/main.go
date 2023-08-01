package main

import (
	"flag"
	"log"
	"os"

	"FDVersion/pkg/FileSystemWalker"
)

func main() {
	pathPtr := flag.String("path", "", "Full path for data input")
	resultPathPtr := flag.String("output", "", "Full path for directory and file info output")
	flag.Parse()
	println("Input: " + *pathPtr)
	checkPath(pathPtr, "input")
	//checkPath(resultPathPtr, "result")

	fsw := FileSystemWalker.New()
	res := fsw.Visit(*pathPtr)
	fsw.CalculateHashes(res)

	println(*resultPathPtr)
	outputXML := res.XML()
	println(outputXML)
	SaveStringToFile(*resultPathPtr, outputXML)
}

func checkPath(pathPtr *string, name string) {
	exists, err := CheckIfPathExists(*pathPtr)

	LogFatalIfError(err)

	if !exists {
		log.Fatalf("Error: %s path does not exist", name)
	}
}

// CheckIfPathExists returns is file or directory exists or an error
func CheckIfPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	} else {
		return true, err
	}
}

func SaveStringToFile(path string, content string) {
	f, err := os.Create(path)
	defer f.Close()
	LogFatalIfError(err)

	_, err_w := f.WriteString(content)
	LogFatalIfError(err_w)
}

func LogFatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

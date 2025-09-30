package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"time"
)

//go:embed quotes/*
var quotesFS embed.FS

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	files, shouldReturn := getQuoteFiles()
	if shouldReturn {
		return
	}
	content := readRandomFile(files, r)
	fmt.Println(string(content))
}

func readRandomFile(files []fs.DirEntry, r *rand.Rand) []byte {
	randomFile := files[r.Intn(len(files))].Name()
	filePath := "quotes/" + randomFile
	content, err := quotesFS.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read embedded file %s: %v", filePath, err)
	}
	return content
}

func getQuoteFiles() ([]fs.DirEntry, bool) {
	files, err := quotesFS.ReadDir("quotes")
	if err != nil {
		log.Fatalf("Failed to read embedded directory: %v", err)
	}

	if len(files) == 0 {
		fmt.Println("No ASCII art files found in embedded quotes directory")
		return nil, true
	}
	return files, false
}

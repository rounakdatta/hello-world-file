package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func readAllFilesToStringFromDirectory(directoryPath string) string {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return "DIRECTORY INVALID"
	}

	directoryContents := ""
	for _, file := range files {
		if !file.IsDir() {
			fileContent, err := os.ReadFile(filepath.Join(directoryPath, file.Name()))
			if err == nil {
				directoryContents += fmt.Sprintf("%s\n---\n%s\n\n", file.Name(), fileContent)
			}
		}
	}
	return directoryContents
}

func handler(w http.ResponseWriter, r *http.Request) {
	directoryPath := os.Getenv("DIRECTORY")
	displayText := readAllFilesToStringFromDirectory(directoryPath)
	fmt.Fprintf(w, displayText)
}

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = "6060"
	}
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

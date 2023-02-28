package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func applyJsonIndentation(contents map[string]interface{}) string {
	output, err := json.MarshalIndent(contents, "", "\t")
	if err != nil {
		return ""
	}
	return string(output)
}

func formatFileContent(fileContent string) map[string]interface{} {
	contentJson := make(map[string]interface{})
	_ = json.Unmarshal([]byte(fileContent), &contentJson)

	// specific use-case, wouldn't panic on key absence
	delete(contentJson, "valueTypeUrn")
	delete(contentJson, "hidden")
	return contentJson
}

func readAllFilesToStringFromDirectory(directoryPath string, isJsonMode bool) string {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return err.Error()
	}

	directoryContents := make(map[string]interface{})
	for _, file := range files {
		if !file.IsDir() {
			fileContent, err := ioutil.ReadFile(filepath.Join(directoryPath, file.Name()))
			if err == nil {
				if !isJsonMode {
					directoryContents[file.Name()] = string(fileContent)
				} else {
					directoryContents[file.Name()] = formatFileContent(string(fileContent))
				}
			}
		}
	}
	return applyJsonIndentation(directoryContents)
}

func handler(w http.ResponseWriter, r *http.Request) {
	directoryPath := os.Getenv("DIRECTORY")
	// json mode flag controls whether the environment variable value would have to be deserialized from a JSON or not
	isJsonMode := false
	if os.Getenv("JSON_MODE") != "" {
		isJsonMode = true
	}

	displayText := readAllFilesToStringFromDirectory(directoryPath, isJsonMode)
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

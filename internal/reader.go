package internal

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
)

/**
 * GetFileContent reads an XML file and returns its content as JSON string
 * @param filepath string
 * @return string
 * @return error
 */
func getFileContent(filepath string) (string, error) {
	//Get the file pathcp
	f, err := os.Open(filepath)
	if err != nil {
		return "", throwError(OpenFileError, err)
	}
	//Close the file
	defer f.Close()
	//Read the file content
	content, err := io.ReadAll(f)
	if err != nil {
		return "", throwError(ReadFileError, err)
	}
	// Parse the file content
	invoice := InvoiceFile{}
	if err := xml.Unmarshal(content, &invoice); err != nil {
		return "", throwError(ParseFileError, err)
	}
	// Encode the file content as json
	json, err := json.Marshal(invoice)
	if err != nil {
		return "", throwError(DecodeFileError, err)
	}
	// Return the json string
	return string(json), nil
}

func listFiles(directory string) {
	// Read the files in the directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
		content, err := getFileContent(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(content)
	}
}

func processFiles() {
	if dir := os.Getenv("INPUT_DIR"); dir == "" {
		log.Fatal(OpenInputDirError)
	} else {
		listFiles(dir)
	}
}

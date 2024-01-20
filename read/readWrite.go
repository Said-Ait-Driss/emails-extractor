package read_write

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func ReadJSON(urls *[]string) error {
	// Open the JSON file
	file, err := os.Open("urls_to_scrap.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// Read the file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(data, urls)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return err
	}
	return err
}

func SaveContent(content []string) error {

	currentTime := time.Now()

	filename := fmt.Sprintf("file_%s.json", currentTime.Format("2006-01-02"))

	quotedStrings := make([]string, len(content))
	for i, s := range content {
		quotedStrings[i] = fmt.Sprintf(`"%s"`, s)
	}

	stringToSaved := "[" + strings.Join(quotedStrings, ", ") + "]"

	toSave := []byte(stringToSaved)

	return ioutil.WriteFile(filename, toSave, 0644)

}

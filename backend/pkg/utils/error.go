package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func WriteError(w http.ResponseWriter, statusCode int, message string, err error) {
	http.Error(w, fmt.Sprintf("%s: %v", message, err), statusCode)
}

var messages map[string]map[string]string

func LoadMessages(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&messages); err != nil {
		return err
	}
	return nil
}

func GetErrorMessage(code, lang string) string {
	if langMessages, ok := messages[code]; ok {
		if message, ok := langMessages[lang]; ok {
			return message
		}
	}
	return "Unknown error"
}

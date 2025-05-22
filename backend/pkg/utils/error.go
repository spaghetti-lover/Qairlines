package utils

import (
	"fmt"
	"net/http"
)

func WriteError(w http.ResponseWriter, statusCode int, message string, err error) {
	http.Error(w, fmt.Sprintf("%s: %v", message, err), statusCode)
}

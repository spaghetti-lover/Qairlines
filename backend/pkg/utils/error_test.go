package utils

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteError(t *testing.T) {
	rec := httptest.NewRecorder()
	err := errors.New("something went wrong")
	WriteError(rec, http.StatusBadRequest, "Bad Request", err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Bad Request: something went wrong")
}

func TestGetErrorMessage(t *testing.T) {
	// Setup messages map manually
	messages = map[string]map[string]string{
		"ERR001": {
			"en": "Invalid input",
			"vi": "Dữ liệu không hợp lệ",
		},
	}

	assert.Equal(t, "Invalid input", GetErrorMessage("ERR001", "en"))
	assert.Equal(t, "Dữ liệu không hợp lệ", GetErrorMessage("ERR001", "vi"))
	assert.Equal(t, "Unknown error", GetErrorMessage("ERR002", "en"))
	assert.Equal(t, "Unknown error", GetErrorMessage("ERR001", "jp"))
}

func TestLoadMessages_Success(t *testing.T) {
	// Tạo file tạm chứa JSON hợp lệ
	tmpFile, err := os.CreateTemp("", "messages.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	content := `{
        "ERR001": {"en": "Invalid input", "vi": "Dữ liệu không hợp lệ"},
        "ERR002": {"en": "Not found", "vi": "Không tìm thấy"}
    }`
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()

	err = LoadMessages(tmpFile.Name())
	assert.NoError(t, err)
	assert.Equal(t, "Invalid input", GetErrorMessage("ERR001", "en"))
	assert.Equal(t, "Không tìm thấy", GetErrorMessage("ERR002", "vi"))
}

func TestLoadMessages_FileNotFound(t *testing.T) {
	err := LoadMessages("not_exist_file.json")
	assert.Error(t, err)
}

func TestLoadMessages_InvalidJSON(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "messages_invalid.json")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(`{invalid json}`)
	assert.NoError(t, err)
	tmpFile.Close()

	err = LoadMessages(tmpFile.Name())
	assert.Error(t, err)
}

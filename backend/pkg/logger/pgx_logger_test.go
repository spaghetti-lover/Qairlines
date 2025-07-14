package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSQL_RemoveQuotesAndComments(t *testing.T) {
	sql := `-- name: ListNews :many
SELECT id, title FROM \"news\" -- comment
WHERE id = $1`
	info := parseSQL(sql)
	assert.Equal(t, "ListNews", info.QueryName)
	assert.Equal(t, "MANY", info.OperationType)
	assert.NotContains(t, info.CleanSQL, "\"")
	assert.NotContains(t, info.CleanSQL, "\\")
	assert.NotContains(t, info.CleanSQL, "--")
	assert.Contains(t, info.CleanSQL, "SELECT id, title FROM news WHERE id = $1")
}

func TestFormatArg(t *testing.T) {
	assert.Equal(t, "'abc'", formatArg("abc"))
	assert.Equal(t, "true", formatArg(true))
	assert.Equal(t, "123", formatArg(123))
	assert.Equal(t, "NULL", formatArg(nil))
	assert.Equal(t, "'2023-01-01T00:00:00Z'", formatArg(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, "''", formatArg(""))
	assert.Equal(t, "'test with spaces'", formatArg("test with spaces"))
	assert.Equal(t, "'test with ''quotes'''", formatArg("test with 'quotes'"))
	assert.Equal(t, "0", formatArg(0))
	assert.Equal(t, "false", formatArg(false))
}

func TestReplacePlaceHolders_NoPlaceholders(t *testing.T) {
	sql := "SELECT * FROM users"
	args := []any{1, "John"}
	result := replacePlaceHolders(sql, args)
	assert.Equal(t, "SELECT * FROM users", result)
}

func TestReplacePlaceHolders_SinglePlaceholder(t *testing.T) {
	sql := "SELECT * FROM users WHERE id = $1"
	args := []any{42}
	result := replacePlaceHolders(sql, args)
	assert.Equal(t, "SELECT * FROM users WHERE id = 42", result)
}

func TestReplacePlaceHolders_MorePlaceholdersThanArgs(t *testing.T) {
	sql := "SELECT * FROM users WHERE id = $1 AND name = $2 AND age = $3"
	args := []any{7, "Alice"}
	result := replacePlaceHolders(sql, args)
	// $3 không có arg, nên sẽ giữ nguyên $3
	assert.Equal(t, "SELECT * FROM users WHERE id = 7 AND name = 'Alice' AND age = $3", result)
}

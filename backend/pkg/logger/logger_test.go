package logger

import (
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPrettyJSONWriter_Write_ValidJSON(t *testing.T) {
	var buf bytes.Buffer
	writer := PrettyJSONWriter{Writer: &buf}
	input := []byte(`{"foo":"bar","baz":1}`)

	n, err := writer.Write(input)
	assert.NoError(t, err)
	assert.Greater(t, n, 0)
	// Result should have newlines and spaces (pretty formatted)
	assert.Contains(t, buf.String(), "\n")
	assert.Contains(t, buf.String(), "  \"foo\": \"bar\"")
}

func TestPrettyJSONWriter_Write_InvalidJSON(t *testing.T) {
	var buf bytes.Buffer
	writer := PrettyJSONWriter{Writer: &buf}
	input := []byte(`not a json`)

	n, err := writer.Write(input)
	assert.NoError(t, err)
	assert.Greater(t, n, 0)
	// Result should be same as input since it cannot be formatted
	assert.Equal(t, "not a json", buf.String())
}

func TestGetTraceID(t *testing.T) {
	ctx := context.WithValue(context.Background(), TraceIdKey, "abc-123")
	traceID := GetTraceID(ctx)
	assert.Equal(t, "abc-123", traceID)

	emptyCtx := context.Background()
	traceID = GetTraceID(emptyCtx)
	assert.Equal(t, "", traceID)
}

func TestNewLoggerWithPath(t *testing.T) {
	// Only test initialization, not actual file writing
	logger := NewLoggerWithPath("test.log", "info")
	assert.IsType(t, &zerolog.Logger{}, logger)
}

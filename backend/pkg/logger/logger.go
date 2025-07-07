package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spaghetti-lover/qairlines/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type contextKey string

const TraceIdKey contextKey = "trace_id"

type LoggerConfig struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	IsDev      string
}

var writer io.Writer

func NewLogger(config LoggerConfig) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	lvl, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)

	switch config.IsDev {
	case "development":
		writer = PrettyJSONWriter{Writer: os.Stdout}
	case "production":
		writer = &lumberjack.Logger{
			Filename:   config.Filename,
			MaxSize:    config.MaxSize, // megabytes
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,   // days
			Compress:   config.Compress, // disabled by default
		}
	default:
		log.Fatal(`IsDev only has 2 options: "development" and "production"`)
	}

	logger := zerolog.New(writer).With().Timestamp().Logger()
	return &logger
}

// Overwrite Writer allows to set a custom writer for the logger.
type PrettyJSONWriter struct {
	Writer io.Writer
}

func (w PrettyJSONWriter) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, p, "", "  ")
	if err != nil {
		return w.Writer.Write(p)
	}

	return w.Writer.Write(prettyJSON.Bytes())
}

func NewLoggerWithPath(path string, level string) *zerolog.Logger {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	settings := LoggerConfig{
		Level:      level,
		Filename:   path,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     5,    // days
		Compress:   true, // disabled by default
		IsDev:      config.AppEnv,
	}
	return NewLogger(settings)
}

func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(TraceIdKey).(string); ok {
		return traceID
	}
	return ""
}

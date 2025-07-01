package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

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

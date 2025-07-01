package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func RecoveryMiddleware(recoveryLogger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				stack_at := ExtractFirstAppStackLine(stack)

				recoveryLogger.Error().
					Str("path", ctx.Request.URL.Path).
					Str("method", ctx.Request.Method).
					Str("client_ip", ctx.ClientIP()).
					Str("stack_at", stack_at).
					Str("stack", string(stack)).
					Str("panic", fmt.Sprintf("%v", err)).
					Msg("Panic occured")
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":    "Internal Server Error",
					"msessage": "An unexpected error occurred. Please try again later.",
				})
			}
		}()
		ctx.Next()
	}
}

var stackLineRegex = regexp.MustCompile(`(.+\.go:\d+)`)

func ExtractFirstAppStackLine(stack []byte) string {
	lines := bytes.Split(stack, []byte("\n"))

	for _, line := range lines {
		if bytes.Contains(line, []byte(".go")) &&
			!bytes.Contains(line, []byte("/runtime/")) &&
			!bytes.Contains(line, []byte("/debug")) &&
			!bytes.Contains(line, []byte("recovery_middleware.go")) {
			cleanLine := strings.TrimSpace(string(line))
			match := stackLineRegex.FindStringSubmatch(cleanLine)

			if len(match) > 1 {
				return match[1]
			}
		}
	}
	return ""
}

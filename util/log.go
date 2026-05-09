package util

import (
	"fmt"
	"log/slog"
	"os"
)

// Info: logs an informational message
func Info(format string, args ...any) {
	if len(args) == 0 {
		slog.Info(format)
		return
	}
	slog.Info(fmt.Sprintf(format, args...))
}

// // Json: logs a JSON message formatted
// func Json(data []byte, format string, args ...any) {
// 	slog.Info(fmt.Sprintf(format, args...), slog.String("json", string(data)))
// }

// Warn: logs a warning message if the error is not nil.
// It's generally more idiomatic to pass a non-nil error directly to logging functions.
// Consider renaming to WarnOnError or removing the nil check if callers are expected to handle it.
func Warn(err error, format string, args ...any) {
	if err != nil {
		// Use slog's structured logging capabilities for better error handling.
		slog.Warn(fmt.Sprintf(format, args...), "error", err)
	}
}

// Fail: logs an error message and exits the program if the error is not nil.
// In a utility package, returning an error or panicking might be preferred over os.Exit(1)
// to allow the caller to control the application's exit behavior.
func Fail(err error, format string, args ...any) {
	if err != nil {
		// Use slog's structured logging capabilities for better error handling.
		slog.Error(fmt.Sprintf(format, args...), "error", err)
		os.Exit(1)
	}
}

package util

import (
	"fmt"
	"log/slog"
	"os"
)

// Info: logs an informational message
func Info(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}

// // Json: logs a JSON message formatted
// func Json(data []byte, format string, args ...any) {
// 	slog.Info(fmt.Sprintf(format, args...), slog.String("json", string(data))) 
// }

// Warn: logs a warning message if the error is not nil.
func Warn(err error, format string, args ...any) {
	if err != nil {
		slog.Warn(fmt.Sprintf("%s: %s", fmt.Sprintf(format, args...), err))
	}
}

// Fail: logs an error message and exits the program if the error is not nil.
func Fail(err error, format string, args ...any) {
	if err != nil {
		slog.Error(fmt.Sprintf("%s: %s", fmt.Sprintf(format, args...), err))
		os.Exit(1)
	}
}

package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Log levels
const (
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

// Color codes for terminal
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Yellow  = "\033[33m"
	Green   = "\033[32m"
	Magenta = "\033[35m"
)

var (
	infoLogger   *log.Logger
	warnLogger   *log.Logger
	errorLogger  *log.Logger
	currentLevel string
)

// Initialize sets up the loggers with the specified log level
func Initialize(logLevel string) {
	// Set log output format and destination
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)

	// Set the current log level
	currentLevel = strings.ToUpper(logLevel)

	// Initialize loggers for each level with prefixes
	infoLogger = log.New(os.Stdout, fmt.Sprintf("%sINFO: %s", Green, Reset), log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(os.Stdout, fmt.Sprintf("%sWARN: %s", Yellow, Reset), log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, fmt.Sprintf("%sERROR: %s", Red, Reset), log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs an informational message with green color
func Info(message string) {
	if shouldLog(LevelInfo) {
		infoLogger.Println(message)
	}
}

// Warn logs a warning message with yellow color
func Warn(message string) {
	if shouldLog(LevelWarn) {
		warnLogger.Println(message)
	}
}

// Error logs an error message with red color
func Error(message string, err error) {
	if shouldLog(LevelError) {
		if err != nil {
			errorLogger.Println(fmt.Sprintf("%s: %v", message, err))
		} else {
			errorLogger.Println(message)
		}
	}
}

// shouldLog checks if a message should be logged based on the current log level
func shouldLog(level string) bool {
	switch currentLevel {
	case LevelInfo:
		return true // Log all levels if log level is INFO
	case LevelWarn:
		return level == LevelWarn || level == LevelError // Log warnings and errors
	case LevelError:
		return level == LevelError // Only log errors
	default:
		return false
	}
}

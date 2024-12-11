// Package clog provides an enhanced logging system with colored output, multiple log levels,
// and debug features for Go applications.
//
// Basic usage:
//
//	clog.Info("Starting application...")
//	clog.Success("Server started on port %d", 8080)
//	clog.Error("Failed to connect: %v", err)
//
// Log Levels (from lowest to highest):
//   - PanicLevel: System is unusable, halts execution
//   - ErrorLevel: Error events that might still allow the application to continue running
//   - WarningLevel: Warning messages for potentially harmful situations
//   - InfoLevel: General informational messages about system operation
//   - DebugLevel: Detailed information for debugging
//   - TraceLevel: Extremely detailed debugging information
//
// Configuration:
//
//	clog.SetLogLevel(clog.DebugLevel)
//	clog.SetShowFileLine(true)
//	clog.SetShowGoroutineID(true)
//
// Features:
//   - Colored output using emoji prefixes
//   - UTC timestamp with millisecond precision
//   - File and line number tracking
//   - Goroutine ID tracking
//   - Performance metrics logging
//   - Multiple log levels with filtering
//   - Panic handling with stack traces
package clog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	// Existing prefixes
	infoPrefix    = color.New(color.FgCyan).Sprint("ℹ️  INFO     ")
	successPrefix = color.New(color.FgGreen).Sprint("✅ SUCCESS  ")
	initPrefix    = color.New(color.FgBlue).Sprint("🚀 INIT     ")
	configPrefix  = color.New(color.FgMagenta).Sprint("⚙️  CONFIG   ")
	warningPrefix = color.New(color.FgYellow).Sprint("⚠️  WARNING  ")
	errorPrefix   = color.New(color.FgRed).Sprint("❌ ERROR    ")

	// New debug prefixes
	debugPrefix  = color.New(color.FgWhite).Sprint("🔍 DEBUG    ")
	tracePrefix  = color.New(color.FgHiBlue).Sprint("📍 TRACE    ")
	panicPrefix  = color.New(color.FgHiRed, color.Bold).Sprint("💥 PANIC    ")
	metricPrefix = color.New(color.FgHiMagenta).Sprint("📊 METRIC   ")
)

// LogLevel type for controlling log output
type LogLevel int

const (
	// PanicLevel logs and then calls panic()
	PanicLevel LogLevel = iota

	// ErrorLevel indicates error conditions
	ErrorLevel

	// WarningLevel indicates potentially harmful situations
	WarningLevel

	// InfoLevel indicates general operational information
	InfoLevel

	// DebugLevel indicates detailed debug information
	DebugLevel

	// TraceLevel indicates the most detailed debugging information
	TraceLevel
)

var (
	currentLogLevel = InfoLevel
	showFileLine    = true
	showGoroutineID = true
)

// SetLogLevel sets the current logging level
func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

// SetShowFileLine enables/disables file and line number in logs
func SetShowFileLine(show bool) {
	showFileLine = show
}

// SetShowGoroutineID enables/disables goroutine ID in logs
func SetShowGoroutineID(show bool) {
	showGoroutineID = show
}

func getFileInfo() string {
	if !showFileLine {
		return ""
	}
	_, file, line, ok := runtime.Caller(3) // Skip 3 frames to get to the caller
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", filepath.Base(file), line)
}

func getGoroutineID() string {
	if !showGoroutineID {
		return ""
	}
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	id := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	return fmt.Sprintf("(goroutine %s)", id)
}

func logWithTimestamp(prefix, msg string, level LogLevel) {
	if level > currentLogLevel {
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05.000 UTC")
	fileInfo := getFileInfo()
	goroutineInfo := getGoroutineID()

	// Construct the log message
	logMsg := fmt.Sprintf("%s [%s]", prefix, timestamp)
	if fileInfo != "" {
		logMsg += fmt.Sprintf(" [%s]", fileInfo)
	}
	if goroutineInfo != "" {
		logMsg += fmt.Sprintf(" %s", goroutineInfo)
	}
	logMsg += fmt.Sprintf(" %s", msg)

	fmt.Println(logMsg)
}

// Existing methods with log level support
func Info(format string, args ...interface{}) {
	logWithTimestamp(infoPrefix, fmt.Sprintf(format, args...), InfoLevel)
}

func Success(format string, args ...interface{}) {
	logWithTimestamp(successPrefix, fmt.Sprintf(format, args...), InfoLevel)
}

func Init(format string, args ...interface{}) {
	logWithTimestamp(initPrefix, fmt.Sprintf(format, args...), InfoLevel)
}

func Config(format string, args ...interface{}) {
	logWithTimestamp(configPrefix, fmt.Sprintf(format, args...), InfoLevel)
}

func Warning(format string, args ...interface{}) {
	logWithTimestamp(warningPrefix, fmt.Sprintf(format, args...), WarningLevel)
}

func Error(format string, args ...interface{}) {
	logWithTimestamp(errorPrefix, fmt.Sprintf(format, args...), ErrorLevel)
}

// New debug methods
func Debug(format string, args ...interface{}) {
	logWithTimestamp(debugPrefix, fmt.Sprintf(format, args...), DebugLevel)
}

func Trace(format string, args ...interface{}) {
	logWithTimestamp(tracePrefix, fmt.Sprintf(format, args...), TraceLevel)
}

func Panic(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logWithTimestamp(panicPrefix, msg, PanicLevel)
	panic(msg)
}

func Metric(name string, value interface{}, tags ...string) {
	tagStr := ""
	if len(tags) > 0 {
		tagStr = fmt.Sprintf(" [%s]", strings.Join(tags, ", "))
	}
	logWithTimestamp(metricPrefix, fmt.Sprintf("%s: %v%s", name, value, tagStr), InfoLevel)
}

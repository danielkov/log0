/*
Package log0 is a simple, pluggable logging library.

A more detailed usage guide can be found at https://github.com/danielkov/log0.

Simple example:

	var log = log0.Default()

	func main() {
		defer log.Close()
		log.Info("This message will log at INFO level.")
	}
*/
package log0

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type level int

// Pseudo-enum of all logging levels
const (
	FINEST level = iota
	FINE
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
	OFF
)

const timeFormat = "2006-01-02 15:04:05.000"

// Logging levels to use by default
var (
	levelStr      = []string{"FNST", "FINE", "DEBUG", "TRACE", "INFO", "WARN", "ERROR", "FATAL"}
	DefaultLevels = []level{INFO, WARNING, ERROR, FATAL}
)

// Logger is a struct that takes in a *bufio.Writer which it writes to with the corresponding logging methods
type Logger struct {
	Writer    *bufio.Writer
	Includes  []level
	ShowTrace bool
	Formatter Formatter
}

// Log is a struct that contains meta data about the log that is created by the methods on Logger.
// Not every Log will be written, only the ones that have a level, which is included in the Loggers levels.
type Log struct {
	Level   level
	Time    time.Time
	Message string
	Frame   runtime.Frame
}

func (l level) String() string {
	return levelStr[int(l)]
}

// Formatter is a function type that is used on the Logger to create a string representation of the Log object
type Formatter func(Log) string

// Level is a method that takes a single level, above which it will include all levels in the logs with the supplied level included
func (l *Logger) Level(lv level) {
	var nl []level
	for i := int(lv); i <= int(FATAL); i++ {
		nl = append(nl, level(i))
	}
	l.Includes = nl
}

// Include is a method that takes any number of levels and adds them to the list of log levels to be displayed
func (l *Logger) Include(levels ...level) {
	for _, lv := range levels {
		if !contains(l.Includes, lv) {
			l.Includes = append(l.Includes, lv)
		}
	}
}

// Exclude is a method that takes anu number of levels and removes them from the list of log levels to be displayed
func (l *Logger) Exclude(levels ...level) {
	l.Includes = diff(l.Includes, levels)
}

// Finest is a logger method that logs the message with the Loggers formatting at FNST level
func (l *Logger) Finest(message string) error {
	return l.writeLog(FINEST, message)
}

// Fine is a logger method that logs the message with the Loggers formatting at FINE level
func (l *Logger) Fine(message string) error {
	return l.writeLog(FINE, message)
}

// Debug is a logger method that logs the message with the Loggers formatting at DEBUG level
func (l *Logger) Debug(message string) error {
	return l.writeLog(DEBUG, message)
}

// Trace is a logger method that logs the message with the Loggers formatting at TRACE level
func (l *Logger) Trace(message string) error {
	return l.writeLog(TRACE, message)
}

// Info is a logger method that logs the message with the Loggers formatting at INFO level
func (l *Logger) Info(message string) error {
	return l.writeLog(INFO, message)
}

// Warning is a logger method that logs the message with the Loggers formatting at WARN level
func (l *Logger) Warning(message string) error {
	return l.writeLog(WARNING, message)
}

// Error is a logger method that logs the message with the Loggers formatting at ERROR level
func (l *Logger) Error(message string) error {
	return l.writeLog(ERROR, message)
}

// Fatal is a logger method that logs the message with the Loggers formatting at FATAL level
func (l *Logger) Fatal(message string) error {
	return l.writeLog(FATAL, message)
}

func (l *Logger) writeLog(lv level, message string) error {
	if !contains(l.Includes, lv) {
		return nil
	}
	var fr runtime.Frame
	if l.ShowTrace {
		fr = trace(4)
	}
	log := Log{lv, time.Now(), message, fr}
	_, err := l.Writer.WriteString(l.Formatter(log))
	return err
}

// Close is a method that cleans up after the Logger is no longer needed. Calling Close() is important because it flushes the buffered writer
func (l *Logger) Close() error {
	return l.Writer.Flush()
}

// DefaultFormatter is a Formatter function that returns the log in the format:
//     <time(UNIXDate)> [<level>] - <message> :: <filename>:<line> -> <function>
func DefaultFormatter(l Log) string {
	return fmt.Sprintf("%s [%s] - %s :: %s:%d -> %s\n", l.Time.Format(time.UnixDate), l.Level, l.Message, filepath.Base(l.Frame.File), l.Frame.Line, l.Frame.Function)
}

// SimpleFormatter is a Formatter function that returns the log in the format:
//     <time> [<level>] <function> - <message>
func SimpleFormatter(l Log) string {
	return fmt.Sprintf("%s [%s] %s - %s\n", l.Time.Format(timeFormat), l.Level, l.Frame.Function, l.Message)
}

// New is a function that returns an instance of Logger
func New(w *bufio.Writer, l []level, t bool, f Formatter) Logger {
	return Logger{w, l, t, f}
}

// Default is a function that returns an instance of Logger with the default config: os.Stdout as output, INFO level, trace generation and DefaultFormatter
func Default() Logger {
	return New(bufio.NewWriter(os.Stdout), DefaultLevels, true, DefaultFormatter)
}

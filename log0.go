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

var (
	STRINGS       []string = []string{"FNST", "FINE", "DEBUG", "TRACE", "INFO", "WARN", "ERROR", "FATAL"}
	DefaultLevels          = []level{INFO, WARNING, ERROR, FATAL}
)

type Logger struct {
	Writer    *bufio.Writer
	Includes  []level
	ShowTrace bool
	Formatter Formatter
}

type Log struct {
	Level   level
	Time    time.Time
	Message string
	Frame   runtime.Frame
}

func (l level) String() string {
	return STRINGS[int(l)]
}

type Formatter func(Log) string

func (l *Logger) Level(lv level) {
	var nl []level
	for i := int(lv); i <= int(FATAL); i++ {
		nl = append(nl, level(i))
	}
	l.Includes = nl
}

func (l *Logger) Include(levels ...level) {
	for _, lv := range levels {
		if !contains(l.Includes, lv) {
			l.Includes = append(l.Includes, lv)
		}
	}
}

func (l *Logger) Exclude(levels ...level) {
	l.Includes = diff(l.Includes, levels)
}

func (l *Logger) Finest(message string) {
	l.writeLog(FINEST, message)
}

func (l *Logger) Fine(message string) {
	l.writeLog(FINE, message)
}

func (l *Logger) Debug(message string) {
	l.writeLog(DEBUG, message)
}

func (l *Logger) Trace(message string) {
	l.writeLog(TRACE, message)
}

func (l *Logger) Info(message string) {
	l.writeLog(INFO, message)
}

func (l *Logger) Warning(message string) {
	l.writeLog(WARNING, message)
}

func (l *Logger) Error(message string) {
	l.writeLog(ERROR, message)
}

func (l *Logger) Fatal(message string) {
	l.writeLog(FATAL, message)
}

func (l *Logger) writeLog(lv level, message string) {
	if !contains(l.Includes, lv) {
		return
	}
	var fr runtime.Frame
	if l.ShowTrace {
		fr = trace(4)
	}
	log := Log{lv, time.Now(), message, fr}
	l.Writer.WriteString(l.Formatter(log))
}

func (l *Logger) Close() {
	l.Writer.Flush()
}

func DefaultFormatter(l Log) string {
	return fmt.Sprintf("%s [%s] - %s :: %s:%d -> %s\n", l.Time.Format(time.UnixDate), l.Level, l.Message, filepath.Base(l.Frame.File), l.Frame.Line, l.Frame.Function)
}

func SimpleFormatter(l Log) string {
	return fmt.Sprintf("%s [%s] %s - %s\n", l.Time.Format(timeFormat), l.Level, l.Frame.Function, l.Message)
}

func New(w *bufio.Writer, l []level, t bool, f Formatter) Logger {
	return Logger{w, l, t, f}
}
func Default() Logger {
	return New(bufio.NewWriter(os.Stdout), DefaultLevels, true, DefaultFormatter)
}

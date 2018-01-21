package log0_test

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/danielkov/log0"
)

func createBufferLogger() (log0.Logger, *bytes.Buffer) {
	var b bytes.Buffer
	return log0.New(bufio.NewWriter(&b), log0.DefaultLevels, false, log0.DefaultFormatter), &b
}

func TestNew(t *testing.T) {
	log, b := createBufferLogger()
	log.Error("testing")
	log.Close()
	got := b.String()
	contains("Has the word testing.", got, "testing", t)
	contains("Shows that it's ERROR level", got, "[ERROR]", t)
}

func TestLevel(t *testing.T) {
	log, b := createBufferLogger()
	log.Level(log0.FATAL)
	log.Error("error")
	log.Fatal("fatal")
	log.Close()
	got := b.String()
	assert("Should not contain error.", strings.Contains(got, "error"), false, t)
	contains("Contains fatal.", got, "fatal", t)
}

func TestInclude(t *testing.T) {
	log, b := createBufferLogger()
	log.Include(log0.FINE)
	log.Fine("fine")
	log.Fine("another")
	log.Debug("debug")
	log.Error("error")
	log.Close()
	got := b.String()
	assert("Should not contain debug.", strings.Contains(got, "debug"), false, t)
	contains("Contains fine.", got, "fine", t)
	contains("Contains error by default.", got, "error", t)
}

func TestExclude(t *testing.T) {
	log, b := createBufferLogger()
	log.Exclude(log0.ERROR)
	log.Error("error")
	log.Fatal("fatal")
	log.Close()
	got := b.String()
	assert("Should not contain error.", strings.Contains(got, "error"), false, t)
	contains("Contains fatal.", got, "fatal", t)
}

func testFormatter(l log0.Log) string {
	return fmt.Sprintf("%s - %s\n", l.Level, l.Message)
}

func TestAllMethods(t *testing.T) {
	log, b := createBufferLogger()
	log.Level(log0.FINEST)
	log.Formatter = testFormatter
	log.Finest("finest")
	log.Fine("fine")
	log.Debug("debug")
	log.Trace("trace")
	log.Info("info")
	log.Warning("warning")
	log.Error("error")
	log.Fatal("fatal")
	log.Close()
	got := b.String()
	contains("Should contain finest.", got, "FNST - finest", t)
	contains("Should contain fine.", got, "FINE - fine", t)
	contains("Should contain debug.", got, "DEBUG - debug", t)
	contains("Should contain trace.", got, "TRACE - trace", t)
	contains("Should contain info.", got, "INFO - info", t)
	contains("Should contain warning.", got, "WARN - warning", t)
	contains("Should contain error.", got, "ERROR - error", t)
	contains("Should contain fatal.", got, "FATAL - fatal", t)
}

func TestTrace(t *testing.T) {
	log, b := createBufferLogger()
	log.ShowTrace = true
	log.Formatter = func(l log0.Log) string {
		return fmt.Sprintf("%s - %s", l.Message, l.Frame.Function)
	}
	log.Fatal("name")
	log.Close()
	got := b.String()
	contains("Contains function name.", got, "name - github.com/danielkov/log0_test.TestTrace", t)
}

func TestDefault(t *testing.T) {
	log := log0.Default()
	n := funcName(log.Formatter)
	assert("Uses default formatter.", n, "github.com/danielkov/log0.DefaultFormatter", t)
}

func TestSimpleFormatter(t *testing.T) {
	log, b := createBufferLogger()
	log.Formatter = log0.SimpleFormatter
	log.Fatal("fatal")
	log.Close()
	got := b.String()
	contains("Contains FATAL.", got, "[FATAL]", t)
	contains("Contains fatal.", got, "- fatal", t)
}

func assert(message string, f, s interface{}, t *testing.T) {
	if f != s {
		t.Errorf("%s\nAssertion failed. Expected: %v, got: %v", message, s, f)
	}
}

func contains(msg, f, s string, t *testing.T) {
	if !strings.Contains(f, s) {
		t.Errorf("%s\nComparison failed. Expected: %s to contain: %s", msg, f, s)
	}
}

func funcName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

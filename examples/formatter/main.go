package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/danielkov/log0"
)

func Formatter(l log0.Log) string {
	return fmt.Sprintf("[%s]: %s at %s\n", l.Level, l.Message, l.Time.Format(time.RFC822))
}

var log = log0.New(bufio.NewWriter(os.Stdout), log0.DefaultLevels, false, Formatter)

func main() {
	defer log.Close()
	log.Fine("This should not print.")
	log.Info("This should print.")
	privateFunc()
	PublicFunc()
}

func privateFunc() {
	log.Warning("Calling log from private function.")
}

func PublicFunc() {
	log.Error("Calling log from public function.")
}

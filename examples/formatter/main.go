package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/danielkov/log0"
)

func formatter(l log0.Log) string {
	return fmt.Sprintf("[%s]: %s at %s\n", l.Level, l.Message, l.Time.Format(time.RFC822))
}

var log = log0.New(bufio.NewWriter(os.Stdout), log0.DefaultLevels, false, formatter)

func main() {
	defer h(log.Close())
	h(log.Fine("This should not print."))
	h(log.Info("This should print."))
	privateFunc()
	PublicFunc()
}

func privateFunc() {
	h(log.Warning("Calling log from private function."))
}

// PublicFunc is a function
func PublicFunc() {
	h(log.Error("Calling log from public function."))
}

func h(e error) {
	if e != nil {
		fmt.Printf("Error: %v", e)
	}
}

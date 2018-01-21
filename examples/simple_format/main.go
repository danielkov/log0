package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danielkov/log0"
)

var log = log0.New(bufio.NewWriter(os.Stdout), log0.DefaultLevels, true, log0.SimpleFormatter)

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

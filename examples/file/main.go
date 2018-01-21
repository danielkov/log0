package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danielkov/log0"
)

var log = log0.New(nil, log0.DefaultLevels, true, log0.DefaultFormatter)

func main() {
	f, err := os.Create("./logfile")
	if err != nil {
		panic("Could not create log file.")
	}
	log.Writer = bufio.NewWriter(f)
	defer func() {
		h(log.Close())
		h(f.Close())
	}()
	h(log.Warning("This should show up in the log file."))
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

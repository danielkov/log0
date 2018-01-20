package main

import (
	"bufio"
	"os"

	"github.com/danielkov/log0"
)

var log = log0.Logger{nil, log0.DefaultLevels, true, log0.DefaultFormatter}

func main() {
	f, err := os.Create("./logfile")
	if err != nil {
		panic("Could not create log file.")
	}
	log.Writer = bufio.NewWriter(f)
	defer func() {
		log.Close()
		f.Close()
	}()
	log.Warning("This should show up in the log file.")
	privateFunc()
	PublicFunc()
}

func privateFunc() {
	log.Warning("Calling log from private function.")
}

func PublicFunc() {
	log.Error("Calling log from public function.")
}

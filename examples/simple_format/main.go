package main

import (
	"bufio"
	"os"

	"github.com/danielkov/log0"
)

var log = log0.New(bufio.NewWriter(os.Stdout), log0.DefaultLevels, true, log0.SimpleFormatter)

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

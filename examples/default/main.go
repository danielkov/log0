package main

import "github.com/danielkov/log0"

var log = log0.Default()

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

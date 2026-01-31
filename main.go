package main

import (
	"log"

	"github.com/softwarespot/pausefy/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err, 1)
	}
}

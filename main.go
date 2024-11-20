package main

import (
	"log"

	"github.com/softwarespot/serve/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err, 1)
	}
}

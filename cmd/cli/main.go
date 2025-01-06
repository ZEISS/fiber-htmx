package main

import (
	"log"
	"os"

	"github.com/zeiss/fiber-htmx/cmd/cli/cmd"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

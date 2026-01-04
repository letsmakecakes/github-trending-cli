package main

import (
	"fmt"
	"log"
	"os"

	"github.com/letsmakecakes/github-trending-cli/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

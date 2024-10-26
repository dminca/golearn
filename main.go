package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("No arguments provided")
	}

	composeContent := strings.Join(args, " \\\n")
	if err := os.WriteFile("outfile.txt", []byte(composeContent), 0644); err != nil {
		log.Fatalf("Error writing args to file: %v", err)
	}
}

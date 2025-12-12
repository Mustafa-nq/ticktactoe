package main

import (
	"bubbletea-cli/internal/ui"
	"log"
	"os"
)

func main() {
	p := ui.NewProgram()
	if err := p.Start(); err != nil {
		log.Printf("Could not start program: %v\n", err)
		os.Exit(1)
	}

}

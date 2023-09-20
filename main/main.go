package main

import (
	"ghostls"
	"log"
	"os"
)

func main() {
	terminalArgs := os.Args[1:]
	if len(terminalArgs) == 0 {
		log.Println("[USAGE]: go run . [OPTIONS] [FILES]")
	}
	mainargs := ghostls.ParseFlags(terminalArgs)
	for _, terminalArgument := range mainargs {
		if ghostls.IsFlag(terminalArgument) {
			continue
		}
		if ghostls.RecursiveSearch {
			ghostls.RecursiveSearchDir(terminalArgument)
		} else {
			ghostls.NormalSearchDir(terminalArgument)
		}
	}
}

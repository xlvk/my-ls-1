package main

import (
	"ghostls"
	"os"
)

func main() {
	terminalArgs := os.Args[1:]
	mainargs := ghostls.ParseFlags(terminalArgs)
	if ghostls.FlagCounter == len(terminalArgs) || len(terminalArgs) == 0 {
		mainargs = append(mainargs, ".")
	}
	for _, terminalArgument := range mainargs {
		if ghostls.IsSingleFlag(terminalArgument) ||ghostls.IsMultiFlag(terminalArgument) {
			continue
		}

		if ghostls.RecursiveSearch {
			ghostls.RecursiveSearchDir(terminalArgument)
		} else {
			ghostls.NormalSearchDir(terminalArgument)
		}
	}
}

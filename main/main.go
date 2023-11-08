package main

import (
	"fmt"
	"ghostls"
	"log"
	"os"
)

//TODO: Fix Block count func
//TODO: handle colors
//TODO: optimize search algorithm

func main() {
	terminalArgs := os.Args[1:]
	mainargs := ghostls.ParseFlags(terminalArgs)
	if ghostls.FlagCounter == len(terminalArgs) || len(terminalArgs) == 0 {
		mainargs = append(mainargs, ".")
	}
	for _, terminalArgument := range mainargs {
		checkisfile, err := os.Lstat(terminalArgument)

		if ghostls.IsSingleFlag(terminalArgument) || ghostls.IsMultiFlag(terminalArgument) {
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		if !checkisfile.IsDir() {
			if ghostls.LongFormat {
				ghostls.LongFormatDisplay(terminalArgument)
			} else {
				fmt.Println(terminalArgument)
			}
		} else {
			ghostls.DirSearcher(terminalArgument)
			fmt.Println()
		}
	}
}

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
		checkisfile, err := os.Stat(terminalArgument)

		if ghostls.IsSingleFlag(terminalArgument) || ghostls.IsMultiFlag(terminalArgument) {
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		if !checkisfile.IsDir() {
			fmt.Println(terminalArgument)
			fmt.Println()
		} else {
			fmt.Println(terminalArgument + ":")
			BlockCount, e := ghostls.GetBlockCount(terminalArgument)
			if ghostls.LongFormat {
				fmt.Println("total", BlockCount)
			}
			if e != nil {
				log.Fatal(e)
			}
			if ghostls.RecursiveSearch {
				ghostls.RecursiveSearchDir(terminalArgument)
			} else {
				ghostls.NormalSearchDir(terminalArgument)
			}
		}

	}
}

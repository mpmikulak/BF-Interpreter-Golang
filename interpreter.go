// Main intitializes and coordinates a primitive brainfuck interpreter
// The binary accepts one argument; this is the filename of the interpreted
// sourcecode.
//
// Contact Mike Mikulak github.com/mpmikulak or mpmikulak@gmail.com

package main

import (
	"flag"
	"fmt"

	"./interpret"
	"./tools"
)

var (
	source  string
	help    bool
	Verbose bool
)

// Init is used to parse command line arguments
func init() {
	flag.StringVar(&source, "f", "", "path to the source code file") // Done
	flag.BoolVar(&Verbose, "v", false, "enables verbose mode")       // In-progress
	flag.BoolVar(&help, "h", false, "display the help page")
	flag.BoolVar(&help, "help", false, "display the help page")

	flag.Parse()
}

func main() {

	switch {
	case help:
		fmt.Println("Place holder for the help page.")
		return
	case flag.NFlag() < 1:
		fmt.Println("Too few arguments. Use -h or --help for help.")
		return
	}

	if Verbose {
		tools.Message("Compiling source.")
	}
	tokens, err := tools.Compile(source)
	if err != nil {
		fmt.Println(err)
	}
	if Verbose {
		tools.Message("Source compiled; checking for errors.")
	}

	if tools.Check(tokens) {
		if Verbose {
			tools.Message("No errors found, running code.")
			interpret.VRun(tokens)
			tools.Message("Program complete.")
		} else {
			interpret.Run(tokens)
		}
	}
}

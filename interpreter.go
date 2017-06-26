// Main intitializes and coordinates a primitive brainfuck interpreter
// The binary accepts one argument; this is the filename of the interpreted
// sourcecode.
//
// Contact Mike Mikulak github.com/mpmikulak or mpmikulak@gmail.com

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mpmikulak/Brainfuck-Interpreter-Golang/interpret"
	"github.com/mpmikulak/Brainfuck-Interpreter-Golang/tools"
)

var (
	source  string
	help    bool
	Verbose bool
	bench   bool
	begin   time.Time
)

// Init is used to parse command line arguments
func init() {
	flag.StringVar(&source, "f", "", "path to the source code file") // Done
	flag.BoolVar(&Verbose, "v", false, "enables verbose mode")       // Done
	flag.BoolVar(&bench, "b", false, "enables benchmarking")         // Done
	flag.BoolVar(&help, "h", false, "display the help page")
	flag.BoolVar(&help, "help", false, "display the help page")

	flag.Parse()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case *os.PathError:
				fmt.Println("Path error.")
			}
		}
	}()

	switch {
	case help:
		flag.PrintDefaults()
		return
	case bench:
		begin = time.Now()
	case flag.NFlag() < 1:
		fmt.Println("Too few arguments. Use -h or --help for help.")
		return
	}

	if Verbose {
		tools.Message("Compiling source.")
	}

	tokens, err := tools.Compile(source)
	if err != nil {
		panic(err)
	}

	if Verbose {
		tools.Message("Source compiled; checking for errors.")
	}

	if tools.Check(tokens) {
		if Verbose {
			tools.Message("No errors found, running code.")
			interpret.VRun(tokens)
			if bench {
				tools.Message(fmt.Sprintf("Program complete. Process took %v.", time.Now().Sub(begin)))
			} else {
				tools.Message("Program complete.")
			}
		} else if bench {
			interpret.Run(tokens)
			tools.Message(fmt.Sprintf("Program complete. Process took %v.", time.Now().Sub(begin)))
		} else {
			interpret.Run(tokens)
		}
	}
}

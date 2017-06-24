// Package main intitializes and coordinates a primitive brainfuck interpreter
// The binary accepts one argument; this is the filename of the interpreted
// sourcecode.
//
// Eight operations exist:
// ">" move the pointer to the cell to the right
// "<" move the pointer to the cell to the left
// "+" increment the byte at the current position
// "-" decrement the byte at the current position
// "." output the byte at the current pointer (os.Stdout)
// "," accept one byte of input, storing it at the current position (os.Stdin)
// "[" if the byte at the data pointer is zero, then instead of moving the
//      instruction pointer forward to the next command, jump it forward to the
//      command after the matching ] command.
// "]" if the byte at the data pointer is nonzero, then instead of moving the
//      instruction pointer forward to the next command, jump it back to the
//      command after the matching [ command.
package main

import (
	"fmt"
	"os"

	"./compile"
	"./interpret"
)

var (
	source string
)

func init() {
	if len(os.Args) < 1 {
		os.Exit(1)
	}
	source = os.Args[1]
}

func main() {
	tokens, err := compile.Compile(source)
	if err != nil {
		fmt.Println(err)
	}
	if interpret.Check(tokens) {
		interpret.Run(tokens)
	}
}

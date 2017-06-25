// Tools is a utility package containing usefull functions for loading and vetting
// the program and custom printing.
package tools

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Compile takes as an argument a string representing the filepath and filename
// of the sourcecode file.  Returned is the whole file as a slice of bytes.
func Compile(s string) ([]byte, error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(f)

	onToken := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return 0, data, nil
	}
	sc.Split(onToken)
	sc.Scan()
	return sc.Bytes(), nil
}

// Message takes a string and prints it with a custom formatted timestamp
func Message(s string) {
	fmt.Printf("%v:%v:%v-%s\n", time.Now().Hour(),
		time.Now().Minute(), time.Now().Second(), s)
}

// Check performs basic bracket counting and mismatch verification
func Check(s []byte) bool {
	bracket := 0
	open := 0
	close := 0
	for _, v := range s {
		switch v {
		case 91:
			bracket++
			open++
		case 93:
			if bracket > 0 {
				bracket--
			}
			close++
		}
	}
	if bracket > 0 {
		fmt.Printf("Unmatched brackets. %v open and %v close.\n", open, close)
		return false
	}
	return true
}

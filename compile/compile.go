// Package compile is used to provide functionality for opening and digesting
// source code.
package compile

import (
	"bufio"
	"os"
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

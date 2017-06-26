// Interpret is used to intitialize the data array and defines its methods.
package interpret

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mpmikulak/Brainfuck-Interpreter-Golang/tools"
)

var reader = bufio.NewReader(os.Stdin)

// Tape is a struct type that contains the fields necessay to emulate a memory
// array tape.
type tape struct {
	array []byte // Contains the data that is manipulated
	index int    // Index is the position of the data pointer
	loopF int    // loopF is used by the loop logic to count open brackets
	loopB int    // loopB is used by the loop logic to count close brackets
	step  int    // Step records the number of passes are made during execution
}

// newTape returns a pointer to a new and intitialized tape type
func newTape() *tape {
	return &tape{[]byte{0}, 0, 0, 0, 0}
}

// String provides customized printing for troubleshooting
func (t *tape) String() string {
	return fmt.Sprintf("%v\t%v: %v", t.array, t.index, t.step)
}

// Right moves the data pointer up by one
func (t *tape) right() {
	t.index++
	if t.index+1 > len(t.array) {
		t.array = append(t.array, 0)
	}
}

// Left moves the data pointer down by one, unless it is already zero, in which
// case it does nothing.
func (t *tape) left() {
	if t.index == 0 {
		return
	}
	t.index--
}

// Up increments the byte that the data pointer is at by one. Allows rollover.
func (t *tape) up() {
	t.array[t.index]++
}

// Down decrements the byte that the data pointer is at by one. Allows rollover.
func (t *tape) down() {
	t.array[t.index]--
}

// Out prints to the console the ASCII encoded value represented by the byte
// pointed to by the data pointer.
func (t *tape) out() {
	fmt.Printf("%v", string(t.array[t.index]))
}

// In accepts input from the console. Only the first character will be used and
// all others are truncated. The value is converted to the byte value and stored
// in the spot pointed to by the data pointer.
func (t *tape) in() {
	for {
		fmt.Println("--------------------")
		fmt.Printf("Enter character: ")
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			t.array[t.index] = byte(char)
			return
		}
	}
}

// Run intitializes a new tape type and passes over the sourcecode, byte by byte.
func Run(s []byte) {
	t := newTape()
	for ix := 0; ix < len(s); ix++ {
		switch s[ix] {
		case 62: // >
			t.right()
		case 60: // <
			t.left()
		case 43: // +
			t.up()
		case 45: // -
			t.down()
		case 46: // .
			t.out()
		case 44: // ,
			t.in()
		case 91: // [
			if t.array[t.index] == 0 {
				for {
					ix++
					if s[ix] == 91 {
						t.loopF++
					} else if s[ix] == 93 {
						if t.loopF == 0 {
							break
						}
						t.loopF--
					}
				}
			}
		case 93: // ]
			if t.array[t.index] > 0 {
				for {
					ix--
					if s[ix] == 93 {
						t.loopB++
					} else if s[ix] == 91 {
						if t.loopB == 0 {
							break
						}
						t.loopB--
					}
				}
			}
		}
		t.step++
	}
}

// VRun is similar to Run, but provides verbose output.
func VRun(s []byte) {
	t := newTape()
	for ix := 0; ix < len(s); ix++ {
		switch s[ix] {
		case 62: // >
			t.right()
			tools.Message(fmt.Sprintf("Move data pointer right to %v", t.index))
		case 60: // <
			t.left()
			tools.Message(fmt.Sprintf("Move data pointer left to %v", t.index))
		case 43: // +
			t.up()
			tools.Message(fmt.Sprintf("Incremented block %v to %v", t.index, t.array[t.index]))
		case 45: // -
			t.down()
			tools.Message(fmt.Sprintf("Decremented block %v to %v", t.index, t.array[t.index]))
		case 46: // .
			tools.Message(fmt.Sprintf("Printing block %v, it is:", t.index))
			t.out()
			fmt.Println()
		case 44: // ,
			t.in()
		case 91: // [
			if t.array[t.index] == 0 {
				for {
					ix++
					if s[ix] == 91 {
						t.loopF++
					} else if s[ix] == 93 {
						if t.loopF == 0 {
							tools.Message(fmt.Sprintf("Jumping ahead in the source to %v", ix))
							break
						}
						t.loopF--
					}
				}
			}
		case 93: // ]
			if t.array[t.index] != 0 {
				for {
					ix--
					if s[ix] == 93 {
						t.loopB++
					} else if s[ix] == 91 {
						if t.loopB == 0 {
							tools.Message(fmt.Sprintf("Jumping back in the source to %v", ix))
							break
						}
						t.loopB--
					}
				}
			}
		}
		t.step++
	}
}

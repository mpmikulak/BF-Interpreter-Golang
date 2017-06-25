// Interpret is used to intitialize the data array and defines its methods.
package interpret

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

type tape struct {
	array []byte
	index int
	loopF int
	loopB int
	step  int
}

func newTape() *tape {
	return &tape{[]byte{0}, 0, 0, 0, 0}
}

func (t *tape) String() string {
	top := fmt.Sprintf("%v\t%v: %v", t.array, t.index, t.step)
	return top
}

func (t *tape) right() {
	t.index++
	if t.index+1 > len(t.array) {
		t.array = append(t.array, 0)
	}
	return
}

func (t *tape) left() {
	if t.index == 0 {
		return
	}
	t.index--
}

func (t *tape) up() {
	t.array[t.index]++
	return
}

func (t *tape) down() {
	t.array[t.index]--
	return
}
func (t *tape) out() {
	fmt.Printf("%v", string(t.array[t.index]))
}
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

func VRun(s []byte) {
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

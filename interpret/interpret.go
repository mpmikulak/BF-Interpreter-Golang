// Interpret is used to intitialize the data array and defines its methods.
package interpret

import "fmt"

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
	// bottom := " "
	// for i := 0; i < t.index; i++ {
	// 	bottom += "  "
	// }
	// bottom += "^"
	return top //+ bottom
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

}

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

func Run(s []byte) {
	t := newTape()
	for ix := 0; ix < len(s); ix++ {
		// fmt.Printf("%s:%v\t%v\n", string(s[ix]), t, ix)
		switch s[ix] {
		case 62: // >
			// fmt.Println("Move right.")
			t.right()
		case 60: // <
			// fmt.Println("Move left.")
			t.left()
		case 43: // +
			// fmt.Println("Increment.")
			t.up()
		case 45: // -
			// fmt.Println("Decrement.")
			t.down()
		case 46: // .
			// fmt.Println("Output.")
			t.out()
		case 44: // ,
			// fmt.Println("Input.")
		case 91: // [
			if t.array[t.index] == 0 {
				for {
					ix++
					// fmt.Println("TEST")
					if s[ix] == 91 {
						// fmt.Println("TEST1")
						t.loopF++
					} else if s[ix] == 93 {
						// fmt.Println("TEST2")
						if t.loopF == 0 {
							// fmt.Println("TEST3")
							break
						}
						t.loopF--
					}
				}
			}
			// fmt.Println("Loop begin.")
		case 93: // ]
			if t.array[t.index] > 0 {
				for {
					ix--
					// fmt.Println("TEST")
					if s[ix] == 93 {
						// fmt.Println("TEST1")
						t.loopB++
					} else if s[ix] == 91 {
						// fmt.Println("TEST2")
						if t.loopB == 0 {
							// fmt.Println("TEST3")
							break
						}
						t.loopB--
					}
				}
			}
			// fmt.Println("Loop end.")
		}
		t.step++
	}
}

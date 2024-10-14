package helper

import "fmt"

type Addr struct {
	Alphabet []string
}

func (a Addr) Addr(excelY, raiseY, x, raiseX int) (topLeft, bottomRight string, merge bool) {
	// excelY starts from 1, not from 0
	// X for alphabet

	//                    col
	//              0   1   2   3   4
	// 	            A   B   C   D   E
	//          1 [A1][  ][  ][  ][  ]
	// excelY   2 [  ][  ][  ][  ][  ]
	//          3 [  ][  ][  ][D3][  ]
	//          4 [  ][  ][  ][  ][  ]

	// excelY = 1
	// x = 0
	// raiseY = 3
	// raiseX = 2

	// topLeft = A1
	// bottomRight = D3

	if excelY <= 0 {
		excelY = 1
	}
	if x < 0 {
		x = 0
	}

	l := len(a.Alphabet)

	// wrong index for excel alphabet
	if x >= l {
		x = l - 1 // last
	}
	x2 := x + raiseX
	if x2 >= l {
		x2 = l - 1 // last
	}

	excelY2 := excelY + raiseY

	// no need to merge cells
	// topLeft and bottomRigth would be equal
	if raiseY <= 0 && raiseX <= 0 {
		addr := fmt.Sprintf("%s%d", a.Alphabet[x], excelY)
		return addr, addr, false
	} else if raiseY < 0 {
		excelY2 = excelY
	} else if raiseX < 0 {
		x2 = x
	}

	// default
	return fmt.Sprintf("%s%d", a.Alphabet[x], excelY), fmt.Sprintf("%s%d", a.Alphabet[x2], excelY2), true
}

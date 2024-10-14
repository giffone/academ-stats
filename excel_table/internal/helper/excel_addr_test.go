package helper_test

import (
	"excel_table/internal/helper"
	"testing"
)

type data struct {
	attr attr
	ans  ans
}

type attr struct {
	row, col, raiseRow, raiseCol int
}

type ans struct {
	topLeft, bottomRight string
	merge                bool
}

func TestAddr(t *testing.T) {
	wantAlph := []string{"A", "B", "C", "D", "E", "F", "G"}

	a := helper.Addr{
		Alphabet: helper.GenExcelColumnAlphabet(len(wantAlph)),
	}

	if len(wantAlph) != len(a.Alphabet) {
		t.Fatalf("length alphabet want \"%d\" but got \"%d\"", len(wantAlph), len(a.Alphabet))
	}

	for i := 0; i < len(wantAlph); i++ {
		if wantAlph[i] != a.Alphabet[i] {
			t.Fatal("alphabet is wrong, can not continue")
		}
	}

	want := []data{
		{
			// 0 - no raise
			attr: attr{
				row:      0,
				col:      0,
				raiseRow: 0,
				raiseCol: 0,
			},
			ans: ans{
				topLeft:     "A1",
				bottomRight: "A1",
				merge:       false,
			},
		},
		{
			// 1 - no raise
			attr: attr{
				row:      -1,
				col:      -1,
				raiseRow: -1,
				raiseCol: -1,
			},
			ans: ans{
				topLeft:     "A1",
				bottomRight: "A1",
				merge:       false,
			},
		},
		{
			// 2 - no raise
			attr: attr{
				row:      0,
				col:      0,
				raiseRow: -10,
				raiseCol: -20,
			},
			ans: ans{
				topLeft:     "A1",
				bottomRight: "A1",
				merge:       false,
			},
		},
		{
			// 3 - no raise
			attr: attr{
				row:      100,
				col:      0,
				raiseRow: 0,
				raiseCol: -100,
			},
			ans: ans{
				topLeft:     "A100",
				bottomRight: "A100",
				merge:       false,
			},
		},
		{
			// 4 - no raise
			attr: attr{
				row:      5,
				col:      5,
				raiseRow: 0,
				raiseCol: 0,
			},
			ans: ans{
				topLeft:     "F5",
				bottomRight: "F5",
				merge:       false,
			},
		},
		{
			// 5 - no raise
			attr: attr{
				row:      100,
				col:      100,
				raiseRow: 0,
				raiseCol: 0,
			},
			ans: ans{
				topLeft:     "G100",
				bottomRight: "G100",
				merge:       false,
			},
		},
		{
			// 6 - raise only on left
			attr: attr{
				row:      2,
				col:      1,
				raiseRow: 0,
				raiseCol: 100,
			},
			ans: ans{
				topLeft:     "B2",
				bottomRight: "G2",
				merge:       true,
			},
		},
		{
			// 7 - raise only bottom
			attr: attr{
				row:      2,
				col:      1,
				raiseRow: 100,
				raiseCol: 0,
			},
			ans: ans{
				topLeft:     "B2",
				bottomRight: "B102",
				merge:       true,
			},
		},
		{
			// 8 - raise on left and bottom
			attr: attr{
				row:      3,
				col:      1,
				raiseRow: 100,
				raiseCol: 100,
			},
			ans: ans{
				topLeft:     "B3",
				bottomRight: "G103",
				merge:       true,
			},
		},
		{
			// 9 - raise only bottom
			attr: attr{
				row:      -5,
				col:      2,
				raiseRow: 100,
				raiseCol: -100,
			},
			ans: ans{
				topLeft:     "C1",
				bottomRight: "C101",
				merge:       true,
			},
		},
	}

	for i, w := range want {
		topLeft, bottomRight, merge := a.Addr(w.attr.row, w.attr.raiseRow, w.attr.col, w.attr.raiseCol)
		if w.ans.topLeft != topLeft || w.ans.bottomRight != bottomRight || w.ans.merge != merge {
			t.Fatalf("on test [%d] not equal answer: want \"%s-%s-%v\" but got \"%s-%s-%v\"",
				i, w.ans.topLeft, w.ans.bottomRight, w.ans.merge, topLeft, bottomRight, merge)
		}
	}
}

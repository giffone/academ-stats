package helper_test

import (
	"excel_table/internal/helper"
	"testing"
)

func TestExcelTopCadets(t *testing.T) {
	want := []struct {
		colN int
		alph []string
	}{
		{colN: 0, alph: nil},
		{colN: 5, alph: []string{"A", "B", "C", "D", "E"}},
		{colN: 10, alph: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}},
		{colN: 26, alph: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}},
		{colN: 30, alph: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD"}},
		{colN: 60, alph: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH"}},
		{colN: 90, alph: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ", "BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ", "CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL"}},
	}

	for i := 0; i < len(want); i++ {
		got := helper.GenExcelColumnAlphabet(want[i].colN)
		if want[i].colN != len(got) {
			t.Fatalf("length is not equal on test %d", i+1)
		}
		for j, v := range got {
			if want[i].alph[j] != v {
				t.Fatalf("letter is not equal on test %d: want \"%s\" got \"%s\"", i+1, want[i].alph[j], v)
			}
		}
	}
}

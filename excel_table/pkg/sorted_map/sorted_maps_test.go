package sorted_map_test

import (
	"excel_table/pkg/sorted_map"
	"testing"
)

// func TestMain(m *testing.M) {
// 	m.Run()
// }

func TestMapsSorted(t *testing.T) {
	repeat := 10
	src := []int{10, 8, 6, 4, 2, 9, 7, 5, 3, 1}
	storage := sorted_map.NewStorage[int, int](len(src) * repeat)
	sm := storage.Storage()

	for repeat > 0 {
		for _, val := range src {
			sm.Set(val, val)
		}
		repeat--
	}

	sm.Sort()

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(want) != sm.Len() {
		t.Fatalf("length is not equal: want %d but got %d", len(want), sm.Len())
	}

	for i, val := range sm.Range() {
		if val.Value != want[i] {
			t.Fatalf("test ordered: value index[%d] not equal", i)
		}
	}
}

package sorted_map_test

import (
	"academ_stats/pkg/sorted_map"
	"testing"
)

func TestRangeHistory(t *testing.T) {
	repeat := 2
	src := []int{10, 8, 6, 4, 2, 9, 7, 5, 3, 1}
	sm := sorted_map.Map[int, byte](len(src) * repeat)

	for repeat > 0 {
		for _, v := range src {
			sm.Set(v, 0)
		}
		sm.ReNew() // clear
		repeat--
	}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if len(want) != sm.History().Len() {
		for _, v := range want {
			if _, ok := sm.History().Get(v); !ok {
				t.Fatalf("did not find %d", v)
			}
		}
	}

	for i, v := range sm.History().Range(true) {
		if v.Key != want[i] {
			t.Fatalf("test ordered: value index[%d] not equal", i)
		}
	}
}

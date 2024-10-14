package top_cadets

import (
	"excel_table/internal/domain"
	"excel_table/pkg/sorted_map"
)

const RGBGray = "E0E0E0"

type result struct {
	head    [][]domain.Head                               //
	data    [][]domain.Cell                               //
	headIDs map[domain.HeadIndex]map[domain.HeadIndex]int //
}

// makeHead recursively walks through the map "head" and copies the values into the matrix
func (r *result) makeHead(storage sorted_map.Maps[domain.HeadIndex, domain.Head], loop int) {
	if loop >= len(r.head) {
		r.head = append(r.head, []domain.Head{})
	}
	// all keys need to sort
	storage.Sort()

	for _, header := range storage.Range() {
		r.head[loop] = append(r.head[loop], header.Value)
		if header.Next.Len() != 0 {
			r.makeHead(header.Next, loop+1)
		}
	}
}

func (r *result) readIDs() {
	for i, v := range r.head[len(r.head)-1] {
		if _, ok := r.headIDs[v.Main.ID]; !ok {
			r.headIDs[v.Main.ID] = make(map[domain.HeadIndex]int)
		}
		r.headIDs[v.Main.ID][v.ID] = i
	}
}

// makeHead recursively walks through the map "cadet" and copies the values into the matrix
func (r *result) makeData(storage sorted_map.Maps[int, map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell]) {
	length := 0
	// length by last head_4
	if l := len(r.head); l > 0 {
		length = len(r.head[l-1])
	} else {
		return
	}

	desc := func(a, b int) int {
		if a > b {
			return -1
		} else if a < b {
			return 1
		} else {
			return 0
		}
	}

	// first key - "level" need to sort
	storage.SortFunc(desc)

	for _, levels := range storage.Range() {
		// second key - "xp" need to sort
		levels.Next.SortFunc(desc)
		for _, xps := range levels.Next.Range() {
			// third key - "user_id" no need to sort
			for _, users := range xps.Next.Range() {
				r.data = append(r.data, make([]domain.Cell, length))
				// store user_data to matrix using indexes that saved in "result.headIndexes"
				r.usefulData(users.Value)
			}
		}
	}
}

// func (r *result) usefulData(userJourney map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell) {
// 	if userJourney == nil {
// 		return
// 	}

// 	for i, head := range r.head[len(r.head)-1] {
// 		if h1, ok := userJourney[head.Main.ID]; ok {
// 			if h2, ok := h1[head.ID]; ok {
// 				r.data[len(r.data)-1][i] = h2
// 			} else {
// 				r.data[len(r.data)-1][i] = domain.Cell{
// 					Style: domain.Style{
// 						FillColor: RGBGray,
// 					},
// 				}
// 			}
// 		}
// 	}
// }

func (r *result) usefulData(userJourney map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell) {
	if userJourney == nil {
		return
	}

	for head1, v := range userJourney {
		for head4, cell := range v {
			if id, ok := r.headIDs[head1][head4]; ok {
				r.data[len(r.data)-1][id] = cell
			}
		}
	}
}

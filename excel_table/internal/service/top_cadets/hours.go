package top_cadets

import (
	"excel_table/internal/domain"
	"excel_table/pkg/sorted_map"

	"github.com/xuri/excelize/v2"
)

// Hours heads {domain.Head}

//  -----------------
//  | Hours         | head 1
//  -----------------
//  |               | head 2
//  -----------------
//  |               | head 3
//  -----------------
//  | Jan24 - May24 | head 4 (unique_id - user data will be entered by this key)
//  -----------------
//   user data {domain.Cell}

func (c *currCadet) hoursParse(hours int) {
	// save last_head and user_data by same id
	var head1 domain.HeadIndex = domain.IndexHours
	var head4 domain.HeadIndex = domain.IndexHours

	// create style for current cell using default template
	cellStyle := c.cellStyle
	cellStyle.Alignment = &excelize.Alignment{
		Horizontal: domain.AlignHorizontalRight, // for digit
		Vertical:   domain.AlignVerticalCenter,
	}

	// set data
	if _, ok := c.data[head1]; ok {
		cell := domain.Cell{
			ID:    head4,
			Style: &cellStyle,
		}
		cell.SetData(hours)
		c.data[head1][head4] = cell
	}
}

func hoursHead(period string, h sorted_map.Maps[domain.HeadIndex, domain.Head], lHead int) {
	if h == nil {
		return
	}

	var head1 domain.HeadIndex = domain.IndexHours
	var head4 domain.HeadIndex = domain.IndexHours

	h1, ok := h.Get(head1)
	if !ok {
		return
	}

	// first title {hours}
	// empty {}
	// empty {}
	// {mar 24 - jun 24}

	// set raiseX for current head_1
	h1.Value.RaiseX += 1 // only current position

	empty := lHead - 2 // minus first head {title} and last title {that need to create} = middle will be empty

	// set empty
	buf := h1
	for empty > 0 {
		buf = buf.Next.GetOrSet(0, domain.Head{
			Main: &h1.Value,
		})
		empty--
	}

	// set head
	buf.Next.GetOrSet(head4, domain.Head{
		ID:    head4,
		Title: period,
		Main:  &h1.Value,
	})
}

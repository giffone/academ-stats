package top_cadets

import (
	"excel_table/internal/domain"
	"excel_table/internal/domain/request"
	"fmt"

	"github.com/xuri/excelize/v2"
)

// Module heads {domain.Head}

//  -----------------------------------------------
//  | Module                                      |  head 1
//  -----------------------------------------------
//  | Go                             | other      |  head 2
//  -----------------------------------------------
//  | 5000xp      | 10000xp          | 7000xp     |  head 3
//  -----------------------------------------------
//  | go-reloaded | ascii-art | a... | math-skils |  head 4 (unique_id - user data will be entered by this key)
//  -----------------------------------------------
//              user data {domain.Cell}

// moduleParse saves module projects and check/save head for unique project names
func (c *currCadet) moduleParse(events []request.Event) {
	if len(events) == 0 || c.head == nil {
		return
	}

	// create style for current cell using default template
	cellStyle := c.cellStyle
	cellStyle.Alignment = &excelize.Alignment{
		Horizontal: domain.AlignHorizontalCenter,
		Vertical:   domain.AlignVerticalCenter,
	}

	var head1 domain.HeadIndex = domain.IndexModule
	var head2 domain.HeadIndex
	var head3 domain.HeadIndex
	var head4 domain.HeadIndex

	h1, ok := c.head.Get(head1)
	if !ok {
		return
	}

	for _, e := range events {
		for _, l := range e.Languages {
			// need check language queue [because language has no own id]
			lang := c.langQueue(l.Name)

			// save head_2: set first key - language_id
			head2 = domain.HeadIndex(lang.Queue)
			h2 := h1.Next.GetOrSet(head2, domain.Head{
				ID:     head2,
				Title:  lang.Title,
				Main:   &h1.Value,
				RaiseX: -1, // minus current position
			})

			for _, t := range l.Tasks {
				// save head_3: set second key - xp (it will as id)
				head3 = domain.HeadIndex(t.XP)
				h3 := h2.Next.GetOrSet(head3, domain.Head{
					ID:     head3,
					Title:  fmt.Sprintf("%d xp", t.XP),
					Main:   &h1.Value,
					RaiseX: -1, // minus current position
				})

				// save last_head and user_data by same id
				head4 = t.ID

				// save head_4: set third key - task_id
				_, ok := h3.Next.Get(head4)
				if !ok {
					h3.Next.Set(head4, domain.Head{
						ID:    head4,
						Title: t.Name,
						Main:  &h1.Value,
					})
					h1.Value.RaiseX++
					h2.Value.RaiseX++
					h3.Value.RaiseX++
				}

				// and save task to user
				if _, ok := c.data[head1]; ok {
					cell := domain.Cell{
						ID:    head4,
						Style: &cellStyle,
					}
					cell.SetData("done")
					c.data[head1][head4] = cell
				}
			}
		}
	}
}

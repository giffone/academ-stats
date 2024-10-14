package top_cadets

import "excel_table/internal/domain"

func (c *currCadet) langQueue(value string) domain.TableQueue {
	v, ok := c.tq[value]
	if !ok {
		if v, ok := c.tq[""]; ok {
			return v
		} else {
			return domain.TableQueue{
				Title: domain.NameLangOther,
				Queue: 1000,
			}
		}
	}

	return v
}

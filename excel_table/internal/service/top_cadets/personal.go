package top_cadets

import (
	"excel_table/internal/domain"
	"excel_table/internal/domain/request"
	"excel_table/pkg/sorted_map"

	"github.com/xuri/excelize/v2"
)

// Personal heads {domain.Head}

//  -----------------------------------------------
//  | Personal                                    |  head 1
//  -----------------------------------------------
//  |                                             |  head 2
//  -----------------------------------------------
//  |                                             |  head 3
//  -----------------------------------------------
//  | ID | Login | Full name | Age | Gender | ... | head 4 (unique_id - user data will be entered by this key)
//  -----------------------------------------------
//              user data {domain.Cell}

type fields interface {
	ID() domain.HeadIndex
	Name() string
	Value() interface{}
	Align() string
}

func (c *currCadet) personalParse(p request.Personal) []fields {
	// read user data and save it for iteration by f
	f := []fields{
		idType(p.ID),
		loginType(p.Login),
		fullNameType(p.FullName),
		ageType(p.Age),
		genderType(p.Gender),
		admissionType(p.Admission.Name),
		levelType(p.Level),
		xpTotalType(p.XPTotal),
	}

	var head1 domain.HeadIndex = domain.IndexPersonal
	var head4 domain.HeadIndex

	for _, field := range f {
		// create style for current cell using default template
		cellStyle := c.cellStyle
		cellStyle.Alignment = &excelize.Alignment{
			Horizontal: field.Align(),
			Vertical:   domain.AlignVerticalCenter,
		}

		// save last_head and user_data by same id
		head4 = field.ID()

		// set data
		if _, ok := c.data[head1]; ok {
			cell := domain.Cell{
				ID:    head4,
				Style: &cellStyle,
			}
			cell.SetData(field.Value())
			c.data[head1][head4] = cell
		}
	}

	return f
}

func personalHead(f []fields, h sorted_map.Maps[domain.HeadIndex, domain.Head], lHead int) {
	if len(f) == 0 || h == nil {
		return
	}

	var head1 domain.HeadIndex = domain.IndexPersonal
	var head4 domain.HeadIndex

	h1, ok := h.Get(head1)
	if !ok {
		return
	}

	// first title {person}
	// empty {}
	// empty {}
	// {id}, {login}, ...

	// check how much need to rise parent cells by X
	raiseX := h1.Value.RaiseX + len(f)

	// set raiseX for current head_1
	h1.Value.RaiseX = raiseX

	empty := lHead - 2 // minus first head {title} and last title {that need to create} = middle will be empty

	// set empty
	buf := h1
	for empty > 0 {
		buf = buf.Next.GetOrSet(0, domain.Head{
			RaiseX: raiseX,
			Main:   &h1.Value,
		})
		empty--
	}

	for _, field := range f {
		// save last_head and user_data by same id
		head4 = field.ID()

		// set head
		buf.Next.GetOrSet(head4, domain.Head{
			ID:    head4,
			Title: field.Name(),
			Main:  &h1.Value,
		})
	}
}

type (
	idType        int
	loginType     string
	fullNameType  string
	ageType       int
	genderType    string
	admissionType string
	levelType     int
	xpTotalType   int
)

func (x idType) ID() domain.HeadIndex        { return domain.IndexID }
func (x loginType) ID() domain.HeadIndex     { return domain.IndexLogin }
func (x fullNameType) ID() domain.HeadIndex  { return domain.IndexFullName }
func (x ageType) ID() domain.HeadIndex       { return domain.IndexAge }
func (x genderType) ID() domain.HeadIndex    { return domain.IndexGender }
func (x admissionType) ID() domain.HeadIndex { return domain.IndexAdmission }
func (x levelType) ID() domain.HeadIndex     { return domain.IndexLevel }
func (x xpTotalType) ID() domain.HeadIndex   { return domain.IndexXPTotal }

func (x idType) Name() string        { return domain.NameID }
func (x loginType) Name() string     { return domain.NameLogin }
func (x fullNameType) Name() string  { return domain.NameFullName }
func (x ageType) Name() string       { return domain.NameAge }
func (x genderType) Name() string    { return domain.NameGender }
func (x admissionType) Name() string { return domain.NameAdmission }
func (x levelType) Name() string     { return domain.NameLevel }
func (x xpTotalType) Name() string   { return domain.NameXPTotal }

func (x idType) Value() interface{}        { return int(x) }
func (x loginType) Value() interface{}     { return string(x) }
func (x fullNameType) Value() interface{}  { return string(x) }
func (x ageType) Value() interface{}       { return int(x) }
func (x genderType) Value() interface{}    { return string(x) }
func (x admissionType) Value() interface{} { return string(x) }
func (x levelType) Value() interface{}     { return int(x) }
func (x xpTotalType) Value() interface{}   { return int(x) }

func (x idType) Align() string        { return domain.AlignHorizontalRight }
func (x loginType) Align() string     { return domain.AlignHorizontalLeft }
func (x fullNameType) Align() string  { return domain.AlignHorizontalLeft }
func (x ageType) Align() string       { return domain.AlignHorizontalCenter }
func (x genderType) Align() string    { return domain.AlignHorizontalLeft }
func (x admissionType) Align() string { return domain.AlignHorizontalLeft }
func (x levelType) Align() string     { return domain.AlignHorizontalRight }
func (x xpTotalType) Align() string   { return domain.AlignHorizontalRight }

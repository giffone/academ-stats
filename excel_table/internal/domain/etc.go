package domain

type (
	HeadIndex int
)

const (
	ColorBlack       = "#000000"
	ColorDarkGray    = "#A9A9A9"
	ColorSlateGray   = "#708090"
	ColorGray        = "#808080"
	ColorAcaBorder   = "#476983"
	ColorAcaHeadFill = "#c7e3d4"
	ColorAcaHeadFont = "#3a5260"
)

// for sorting fields by ID

func (h HeadIndex) Name() string {
	switch h {
	case IndexPersonal:
		return NamePersonal
	case IndexHours:
		return NameHours
	case IndexModule:
		return NameModule
	case IndexCheckpoint:
		return NameCheckpoint
	default:
		return ""
	}
}

const (
	DefaultFontSize  float64 = 12
	SqlFieldLanguage         = "language"
	DefaultQueue             = 1000

	NameCadet      = "Cadet"
	NameModule     = "Module"
	NameCheckpoint = "Checkpoint"
	NamePiscine    = "Piscine"
	NameLangOther  = "other"
	NameJourney    = "Journey"
	NamePersonal   = "Personal"
	NameHours      = "Hours"

	NameID        = "ID"
	NameLogin     = "Login"
	NameFullName  = "Full name"
	NameAge       = "Age"
	NameGender    = "Gender"
	NameAdmission = "Admission"
	NameLevel     = "Level"
	NameXPTotal   = "XP total"

	AlignHorizontalLeft   = "left"
	AlignHorizontalCenter = "center"
	AlignHorizontalRight  = "right"
	AlignVerticalCenter   = "center"

	CellValueNil              = "nil"
	CellValueTypeNotAvailable = "type N/A"
)

const (
	CadetIndex = -DefaultQueue // for creating negative indexs of Cadet struct fields {UserID, Login, ...} [field_index-1000]

	IndexPersonal HeadIndex = iota + CadetIndex
	IndexHours
	IndexModule
	IndexCheckpoint

	// Fields of Personal
	IndexID
	IndexLogin
	IndexFullName
	IndexAge
	IndexGender
	IndexAdmission
	IndexLevel
	IndexXPTotal
)

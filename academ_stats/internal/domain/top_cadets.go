// structs for internal use
package domain

// for using in grpc
type (
	TopCadets struct {
		Period string      `json:"period"`
		Cadets []CadetData `json:"cadets"`
	}

	CadetData struct {
		Personal Personal              `json:"personal"` // personal data
		Hours    HoursDTO              `json:"hours"`    //
		Journey  map[string][]EventDTO `json:"journey"`  //
	}

	Personal struct {
		PersonalDTO              //
		Admission   AdmissionDTO `json:"admission"` //
	}
)

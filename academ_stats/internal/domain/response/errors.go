package response

import "errors"

type ErrBadReq struct{ Message string }

func (e *ErrBadReq) Error() string { return e.Message }

var (
	ErrNotCreated   = errors.New("not created")
	ErrAccessDenied = errors.New("access denied")
	ErrNotFound     = &ErrBadReq{"not found"}
)

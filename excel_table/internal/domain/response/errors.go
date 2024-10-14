package response

import "errors"

type ErrBadReq struct{ Message string }

func (e *ErrBadReq) Error() string { return e.Message }

var (
	ErrReqEmpty          = errors.New("request is empty")
	ErrReqCadetDataEmpty = errors.New("request: cadet data is empty")
	ErrNoData            = errors.New("no data")
	ErrNotCreated        = errors.New("not created")
	ErrAccessDenied      = errors.New("access denied")
	ErrNotFound          = &ErrBadReq{"not found"}
)

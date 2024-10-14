// structs to use when sending responses to requests from other services
package response

import (
	"encoding/json"
	"time"
)

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Computer struct {
	ComputerName string          `db:"comp_name" json:"comp_name"`
	Versions     json.RawMessage `db:"versions" json:"versions"`
	Updated      time.Time       `db:"updated" json:"updated"`
}

type TokenExpireDate struct {
	Expire    time.Time `json:"expire"`
	PreNotify time.Time `json:"pre-notify"`
}

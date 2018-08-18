package idevice

import (
	"fmt"
)

// Error list
var (
	ErrInvalidArg    = fmt.Errorf("invalid argument")
	ErrUnknown       = fmt.Errorf("unknown error")
	ErrNoDevice      = fmt.Errorf("no device")
	ErrNotEnoughData = fmt.Errorf("not enough data")
	ErrBadHeader     = fmt.Errorf("bad header")
	ErrSSL           = fmt.Errorf("ssl error")
)

var errorMap = map[int]error{
	-1: ErrInvalidArg,
	-2: ErrUnknown,
	-3: ErrNoDevice,
	-4: ErrNotEnoughData,
	-5: ErrBadHeader,
	-6: ErrSSL,
}

func handleError(code int) error {
	e, ok := errorMap[code]
	if !ok {
		e = ErrUnknown
	}

	return e
}

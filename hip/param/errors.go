package param

import "errors"

var (
	// ErrTooShortToDecode indicates that given bytes cannot be decoded
	// because the length of bytes is too short.
	ErrTooShortToDecode = errors.New("too short to decode")
	// ErrInvalidLength indicates that some operation has been failed
	// due to the value of Length field is not valid.
	ErrInvalidLength = errors.New("length value is invalid")
)

// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// EchoRequestUnsigned represents a EchoRequestUnsigned parameter.
//
// Spec: 5.2.20.  ECHO_REQUEST_SIGNED
type EchoRequestUnsigned struct {
	*Header
	OpaqueData []byte
}

// NewEchoRequestUnsigned creates a new EchoRequestUnsigned.
func NewEchoRequestUnsigned(data []byte) *EchoRequestUnsigned {
	e := &EchoRequestUnsigned{
		Header:     &Header{Type: ParamTypeEchoRequestUnsigned},
		OpaqueData: data,
	}

	e.Padding = make([]byte, padlen(len(data)))
	e.SetLength()
	return e
}

// DecodeEchoRequestUnsigned decodes the given bytes as a EchoRequestUnsigned.
func DecodeEchoRequestUnsigned(b []byte) (*EchoRequestUnsigned, error) {
	e := &EchoRequestUnsigned{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes the given bytes as a EchoRequestUnsigned.
func (e *EchoRequestUnsigned) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 4 {
		return ErrTooShortToDecode
	}

	var err error
	e.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	e.OpaqueData = e.Header.Contents

	return nil
}

// Serialize serializes a EchoRequestUnsigned into bytes.
func (e *EchoRequestUnsigned) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a EchoRequestUnsigned into bytes.
func (e *EchoRequestUnsigned) SerializeTo(b []byte) error {
	e.Header.Contents = make([]byte, e.Len()-4)
	copy(e.Header.Contents, e.OpaqueData)

	return e.Header.SerializeTo(b)
}

// Len returns the total length of a EchoRequestUnsigned, including Padding.
func (e *EchoRequestUnsigned) Len() int {
	return 4 + len(e.OpaqueData) + len(e.Padding)
}

// SetLength sets the length of Contents in Length field.
func (e *EchoRequestUnsigned) SetLength() {
	e.Length = uint16(len(e.OpaqueData))
}

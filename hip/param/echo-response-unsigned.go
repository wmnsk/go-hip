// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// EchoResponseUnsigned represents a EchoResponseUnsigned parameter.
//
// Spec: 5.2.23.  ECHO_RESPONSE_UNSIGNED
type EchoResponseUnsigned struct {
	*Header
	OpaqueData []byte
}

// NewEchoResponseUnsigned creates a new EchoResponseUnsigned.
func NewEchoResponseUnsigned(data []byte) *EchoResponseUnsigned {
	e := &EchoResponseUnsigned{
		Header:     &Header{Type: ParamTypeEchoResponseUnsigned},
		OpaqueData: data,
	}

	e.Padding = make([]byte, padlen(len(data)))
	e.SetLength()
	return e
}

// DecodeEchoResponseUnsigned decodes the given bytes as a EchoResponseUnsigned.
func DecodeEchoResponseUnsigned(b []byte) (*EchoResponseUnsigned, error) {
	e := &EchoResponseUnsigned{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes the given bytes as a EchoResponseUnsigned.
func (e *EchoResponseUnsigned) DecodeFromBytes(b []byte) error {
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

// Serialize serializes a EchoResponseUnsigned into bytes.
func (e *EchoResponseUnsigned) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a EchoResponseUnsigned into bytes.
func (e *EchoResponseUnsigned) SerializeTo(b []byte) error {
	e.Header.Contents = make([]byte, e.Len()-4)
	copy(e.Header.Contents, e.OpaqueData)

	return e.Header.SerializeTo(b)
}

// Len returns the total length of a EchoResponseUnsigned, including Padding.
func (e *EchoResponseUnsigned) Len() int {
	return 4 + len(e.OpaqueData) + len(e.Padding)
}

// SetLength sets the length of Contents in Length field.
func (e *EchoResponseUnsigned) SetLength() {
	e.Length = uint16(len(e.OpaqueData))
}

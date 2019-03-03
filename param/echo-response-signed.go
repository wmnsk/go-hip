// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// EchoResponseSigned represents a EchoResponseSigned parameter.
//
// Spec: 5.2.22.  ECHO_RESPONSE_SIGNED
type EchoResponseSigned struct {
	*Header
	OpaqueData []byte
}

// NewEchoResponseSigned creates a new EchoResponseSigned.
func NewEchoResponseSigned(data []byte) *EchoResponseSigned {
	e := &EchoResponseSigned{
		Header:     &Header{Type: ParamTypeEchoResponseSigned},
		OpaqueData: data,
	}

	e.Padding = make([]byte, padlen(len(data)))
	e.SetLength()
	return e
}

// DecodeEchoResponseSigned decodes the given bytes as a EchoResponseSigned.
func DecodeEchoResponseSigned(b []byte) (*EchoResponseSigned, error) {
	e := &EchoResponseSigned{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes the given bytes as a EchoResponseSigned.
func (e *EchoResponseSigned) DecodeFromBytes(b []byte) error {
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

// Serialize serializes a EchoResponseSigned into bytes.
func (e *EchoResponseSigned) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a EchoResponseSigned into bytes.
func (e *EchoResponseSigned) SerializeTo(b []byte) error {
	e.Header.Contents = make([]byte, e.Len()-4)
	copy(e.Header.Contents, e.OpaqueData)

	return e.Header.SerializeTo(b)
}

// Len returns the total length of a EchoResponseSigned, including Padding.
func (e *EchoResponseSigned) Len() int {
	return 4 + len(e.OpaqueData) + len(e.Padding)
}

// SetLength sets the length of Contents in Length field.
func (e *EchoResponseSigned) SetLength() {
	e.Length = uint16(len(e.OpaqueData))
}

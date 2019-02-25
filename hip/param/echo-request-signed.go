// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// EchoRequestSigned represents a EchoRequestSigned parameter.
//
// Spec: 5.2.20.  ECHO_REQUEST_SIGNED
type EchoRequestSigned struct {
	*Header
	OpaqueData []byte
}

// NewEchoRequestSigned creates a new EchoRequestSigned.
func NewEchoRequestSigned(data []byte) *EchoRequestSigned {
	e := &EchoRequestSigned{
		Header:     &Header{Type: ParamTypeEchoRequestSigned},
		OpaqueData: data,
	}

	e.Padding = make([]byte, padlen(len(data)))
	e.SetLength()
	return e
}

// DecodeEchoRequestSigned decodes the given bytes as a EchoRequestSigned.
func DecodeEchoRequestSigned(b []byte) (*EchoRequestSigned, error) {
	e := &EchoRequestSigned{}
	if err := e.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return e, nil
}

// DecodeFromBytes decodes the given bytes as a EchoRequestSigned.
func (e *EchoRequestSigned) DecodeFromBytes(b []byte) error {
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

// Serialize serializes a EchoRequestSigned into bytes.
func (e *EchoRequestSigned) Serialize() ([]byte, error) {
	b := make([]byte, e.Len())
	if err := e.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a EchoRequestSigned into bytes.
func (e *EchoRequestSigned) SerializeTo(b []byte) error {
	e.Header.Contents = make([]byte, e.Len()-4)
	copy(e.Header.Contents, e.OpaqueData)

	return e.Header.SerializeTo(b)
}

// Len returns the total length of a EchoRequestSigned, including Padding.
func (e *EchoRequestSigned) Len() int {
	return 4 + len(e.OpaqueData) + len(e.Padding)
}

// SetLength sets the length of Contents in Length field.
func (e *EchoRequestSigned) SetLength() {
	e.Length = uint16(len(e.OpaqueData))
}

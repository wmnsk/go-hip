// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// DiffieHellman represents a header (common fields) in HIP parameter.
//
// Spec: 5.2.7.  DIFFIE_HELLMAN
type DiffieHellman struct {
	*Header
	GroupID           uint8
	PublicValueLength uint16
	PublicValue       []byte
}

// NewDiffieHellman creates a new DiffieHellman.
func NewDiffieHellman(id uint8, pubVal []byte) *DiffieHellman {
	d := &DiffieHellman{
		Header:            &Header{Type: ParamTypeDiffieHellman},
		GroupID:           id,
		PublicValueLength: uint16(len(pubVal)),
		PublicValue:       pubVal,
	}

	d.Padding = make([]byte, padlen(3+len(pubVal)))
	d.SetLength()
	return d
}

// DecodeDiffieHellman decodes the given bytes as a DiffieHellman.
func DecodeDiffieHellman(b []byte) (*DiffieHellman, error) {
	d := &DiffieHellman{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes the given bytes as a DiffieHellman.
func (d *DiffieHellman) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 9 {
		return ErrTooShortToDecode
	}

	var err error
	d.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	d.GroupID = d.Header.Contents[0]
	d.PublicValueLength = binary.BigEndian.Uint16(d.Header.Contents[1:3])
	d.PublicValue = d.Header.Contents[3 : 3+d.PublicValueLength]

	return nil
}

// Serialize serializes a DiffieHellman into bytes.
func (d *DiffieHellman) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a DiffieHellman into bytes.
func (d *DiffieHellman) SerializeTo(b []byte) error {
	d.Header.Contents = make([]byte, d.Len()-4)
	d.Header.Contents[0] = d.GroupID
	binary.BigEndian.PutUint16(d.Header.Contents[1:3], d.PublicValueLength)
	copy(d.Header.Contents[3:], d.PublicValue)

	return d.Header.SerializeTo(b)
}

// Len returns the total length of a DiffieHellman, including Padding.
func (d *DiffieHellman) Len() int {
	return 4 + 3 + len(d.PublicValue) + len(d.Padding)
}

// SetLength sets the length of Contents in Length field.
func (d *DiffieHellman) SetLength() {
	d.Length = uint16(3 + len(d.PublicValue))
}

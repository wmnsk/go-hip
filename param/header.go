// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
	"fmt"
)

// Header represents a header (common fields) in HIP parameter.
type Header struct {
	Type     uint16
	Length   uint16
	Contents []byte
	Padding  []byte
}

// NewHeader creates a new Header.
func NewHeader(typ uint16, contents []byte) *Header {
	h := &Header{Type: typ, Contents: contents}
	// add Padding
	h.Padding = make([]byte, padlen(len(contents)))

	h.SetLength()
	return h
}

// DecodeHeader decodes the given bytes as a Header.
func DecodeHeader(b []byte) (*Header, error) {
	h := &Header{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a Header.
func (h *Header) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 4 {
		return ErrTooShortToDecode
	}
	h.Type = binary.BigEndian.Uint16(b[0:2])
	h.Length = binary.BigEndian.Uint16(b[2:4])
	ll := int(h.Length + 4)
	if ll > l {
		return ErrInvalidLength
	}

	h.Contents = b[4:ll]
	h.Padding = b[ll : ll+padlen(len(h.Contents))]
	return nil
}

// Serialize serializes a Header into bytes.
func (h *Header) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a Header into bytes.
func (h *Header) SerializeTo(b []byte) error {
	binary.BigEndian.PutUint16(b[0:2], h.Type)
	binary.BigEndian.PutUint16(b[2:4], h.Length)
	ll := int(h.Length + 4)
	copy(b[4:ll], h.Contents)
	if len(h.Padding) != 0 {
		copy(b[ll:], h.Padding)
	}
	return nil
}

// padlen returns the length of Padding, calcurated from the Contents.
func padlen(l int) int {
	if rem := (l + 4) % 8; rem != 0 {
		return 8 - rem
	}
	return 0
}

// Len returns the total length of a Header, including Padding.
func (h *Header) Len() int {
	return 4 + len(h.Contents) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *Header) SetLength() {
	h.Length = uint16(len(h.Contents))
}

// SetCritical sets the critical bit in Header.
func (h *Header) SetCritical() {
	h.Type |= 1
}

// UnsetCritical removes the critical bit in Header.
func (h *Header) UnsetCritical() {
	h.Type &= 0xe
}

// IsCritical reports whether a Header has critical bit or not.
func (h *Header) IsCritical() bool {
	return h.Type%2 == 1
}

// ParamType returns the type of parameter.
func (h *Header) ParamType() uint16 {
	return h.Type
}

// String returns Header in formatted string.
func (h *Header) String() string {
	return fmt.Sprintf("Type: %d, Length: %d, Contents: %x, Padding: %d bytes",
		h.Type, h.Length, h.Contents, len(h.Padding),
	)
}

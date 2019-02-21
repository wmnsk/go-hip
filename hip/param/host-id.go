// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// HostID represents a header (common fields) in HIP parameter.
//
// Spec: 5.2.9.  HOST_ID
type HostID struct {
	*Header
	HILength         uint16
	DIType           uint8
	DILength         uint16
	Algorithm        uint16
	HostIdentity     []byte
	DomainIdentifier []byte
}

// NewHostID creates a new HostID.
func NewHostID(alg uint16, diType uint8, hi, di []byte) *HostID {
	h := &HostID{
		Header:           &Header{Type: ParamTypeHostID},
		HILength:         uint16(len(hi)),
		DIType:           diType,
		DILength:         uint16(len(di)),
		Algorithm:        alg,
		HostIdentity:     hi,
		DomainIdentifier: di,
	}

	h.Header.Padding = make([]byte, padlen(6+len(h.HostIdentity)+len(h.DomainIdentifier)))
	h.SetLength()
	return h
}

// DecodeHostID decodes the given bytes as a HostID.
func DecodeHostID(b []byte) (*HostID, error) {
	h := &HostID{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HostID.
func (h *HostID) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 10 {
		return ErrTooShortToDecode
	}

	var err error
	h.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	h.HILength = binary.BigEndian.Uint16(h.Header.Contents[:2])
	if l < int(h.HILength) {
		return ErrInvalidLength
	}
	h.DIType = (h.Header.Contents[2] >> 4) & 0xf
	h.DILength = binary.BigEndian.Uint16(h.Header.Contents[2:4]) & 0x3f
	if l < int(h.DILength) {
		return ErrInvalidLength
	}
	h.Algorithm = binary.BigEndian.Uint16(h.Header.Contents[4:6])
	offset := 6 + h.HILength
	h.HostIdentity = h.Header.Contents[6:offset]
	h.DomainIdentifier = h.Header.Contents[offset : offset+h.DILength]

	return nil
}

// Serialize serializes a HostID into bytes.
func (h *HostID) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HostID into bytes.
func (h *HostID) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)

	binary.BigEndian.PutUint16(h.Header.Contents[0:2], h.HILength)
	binary.BigEndian.PutUint16(h.Header.Contents[2:4], h.DILength|(uint16(h.DIType)<<12))
	binary.BigEndian.PutUint16(h.Header.Contents[4:6], h.Algorithm)
	offset := 6 + h.HILength
	copy(h.Header.Contents[6:offset], h.HostIdentity)
	copy(h.Header.Contents[offset:offset+h.DILength], h.DomainIdentifier)
	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HostID, including Padding.
func (h *HostID) Len() int {
	return 4 + 6 + len(h.HostIdentity) + len(h.DomainIdentifier) + len(h.Header.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HostID) SetLength() {
	h.Length = uint16(6 + len(h.HostIdentity) + len(h.DomainIdentifier))
}

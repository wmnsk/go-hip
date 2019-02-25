// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import "encoding/binary"

// Ack represents a Ack parameter.
//
// Spec: 5.2.17.  ACK
type Ack struct {
	*Header
	UpdateIDs []uint16
}

// NewAck creates a new Ack.
func NewAck(ids ...uint16) *Ack {
	a := &Ack{
		Header:    &Header{Type: ParamTypeAck},
		UpdateIDs: ids,
	}

	a.Padding = make([]byte, padlen(len(ids)*2))
	a.SetLength()
	return a
}

// DecodeAck decodes the given bytes as a Ack.
func DecodeAck(b []byte) (*Ack, error) {
	a := &Ack{}
	if err := a.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return a, nil
}

// DecodeFromBytes decodes the given bytes as a Ack.
func (a *Ack) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToDecode
	}

	var err error
	a.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	offset := 0
	ll := len(a.Header.Contents)
	for {
		if offset >= ll {
			break
		}
		a.UpdateIDs = append(a.UpdateIDs, binary.BigEndian.Uint16(a.Header.Contents[offset:offset+2]))
		offset += 2
	}

	return nil
}

// Serialize serializes a Ack into bytes.
func (a *Ack) Serialize() ([]byte, error) {
	b := make([]byte, a.Len())
	if err := a.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a Ack into bytes.
func (a *Ack) SerializeTo(b []byte) error {
	a.Header.Contents = make([]byte, a.Len()-4)
	offset := 0
	for _, id := range a.UpdateIDs {
		binary.BigEndian.PutUint16(a.Header.Contents[offset:offset+2], id)
		offset += 2
	}

	return a.Header.SerializeTo(b)
}

// Len returns the total length of a Ack, including Padding.
func (a *Ack) Len() int {
	return 4 + (len(a.UpdateIDs) * 2) + len(a.Padding)
}

// SetLength sets the length of Contents in Length field.
func (a *Ack) SetLength() {
	a.Length = uint16(len(a.UpdateIDs) * 2)
}

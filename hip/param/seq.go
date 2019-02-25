// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import "encoding/binary"

// Seq represents a Seq parameter.
//
// Spec: 5.2.16.  SEQ
type Seq struct {
	*Header
	UpdateID uint32
}

// NewSeq creates a new Seq.
func NewSeq(id uint32) *Seq {
	s := &Seq{
		Header:   &Header{Type: ParamTypeSeq},
		UpdateID: id,
	}

	s.SetLength()
	return s
}

// DecodeSeq decodes the given bytes as a Seq.
func DecodeSeq(b []byte) (*Seq, error) {
	s := &Seq{}
	if err := s.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return s, nil
}

// DecodeFromBytes decodes the given bytes as a Seq.
func (s *Seq) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 8 {
		return ErrTooShortToDecode
	}

	var err error
	s.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	s.UpdateID = binary.BigEndian.Uint32(s.Contents[0:4])

	return nil
}

// Serialize serializes a Seq into bytes.
func (s *Seq) Serialize() ([]byte, error) {
	b := make([]byte, s.Len())
	if err := s.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a Seq into bytes.
func (s *Seq) SerializeTo(b []byte) error {
	s.Header.Contents = make([]byte, s.Len()-4)
	binary.BigEndian.PutUint32(s.Header.Contents[0:4], s.UpdateID)

	return s.Header.SerializeTo(b)
}

// Len returns the total length of a Seq, including Padding.
func (s *Seq) Len() int {
	return 4 + 4 // fixed by spec
}

// SetLength sets the length of Contents in Length field.
func (s *Seq) SetLength() {
	s.Length = 4 // fixed by spec
}

// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// Solution represents a header (common fields) in HIP parameter.
//
// Spec: 5.2.5.  SOLUTION
type Solution struct {
	*Header
	NoOfK          uint8
	Reserved       uint8
	Opaque         uint16
	Random         []byte
	PuzzleSolution []byte
}

// NewSolution creates a new Solution.
func NewSolution(bits uint8, opaque uint16, random, solution []byte) *Solution {
	s := &Solution{
		Header:         &Header{Type: ParamTypeSolution},
		NoOfK:          bits,
		Opaque:         opaque,
		Random:         random,
		PuzzleSolution: solution,
	}

	s.Padding = make([]byte, padlen(4+len(random)+len(solution)))
	s.SetLength()
	return s
}

// DecodeSolution decodes the given bytes as a Solution.
func DecodeSolution(b []byte) (*Solution, error) {
	s := &Solution{}
	if err := s.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return s, nil
}

// DecodeFromBytes decodes the given bytes as a Solution.
func (s *Solution) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 10 {
		return ErrTooShortToDecode
	}

	var err error
	s.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	max := s.Header.Length
	ll := (max - 4) / 2
	s.NoOfK = s.Header.Contents[0]
	s.Reserved = s.Header.Contents[1]
	s.Opaque = binary.BigEndian.Uint16(s.Header.Contents[2:4])
	s.Random = s.Header.Contents[4 : 4+ll]
	s.PuzzleSolution = s.Header.Contents[4+ll : max]

	return nil
}

// Serialize serializes a Solution into bytes.
func (s *Solution) Serialize() ([]byte, error) {
	b := make([]byte, s.Len())
	if err := s.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a Solution into bytes.
func (s *Solution) SerializeTo(b []byte) error {
	s.Header.Contents = make([]byte, s.Len()-4)
	s.Header.Contents[0] = s.NoOfK
	s.Header.Contents[1] = s.Reserved
	binary.BigEndian.PutUint16(s.Header.Contents[2:4], s.Opaque)

	max := s.Length
	ll := (max - 4) / 2
	copy(s.Header.Contents[4:4+ll], s.Random)
	copy(s.Header.Contents[4+ll:max], s.PuzzleSolution)

	return s.Header.SerializeTo(b)
}

// Len returns the total length of a Solution, including Padding.
func (s *Solution) Len() int {
	return 4 + 4 + len(s.Random) + len(s.PuzzleSolution) + len(s.Padding)
}

// SetLength sets the length of Contents in Length field.
func (s *Solution) SetLength() {
	s.Length = uint16(4 + len(s.Random) + len(s.PuzzleSolution))
}

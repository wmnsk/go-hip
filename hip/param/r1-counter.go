// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// R1Counter represents a R1Counter parameter.
//
// Spec: 5.2.3.  R1_COUNTER
type R1Counter struct {
	*Header
	Reserved            []byte
	R1GenerationCounter uint64
}

// NewR1Counter creates a new R1Counter.
func NewR1Counter(counter uint64) *R1Counter {
	r := &R1Counter{
		Header:              &Header{Type: ParamTypeR1Counter},
		R1GenerationCounter: counter,
	}

	r.SetLength()
	return r
}

// DecodeR1Counter decodes the given bytes as a R1Counter.
func DecodeR1Counter(b []byte) (*R1Counter, error) {
	r := &R1Counter{}
	if err := r.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return r, nil
}

// DecodeFromBytes decodes the given bytes as a R1Counter.
func (r *R1Counter) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 12 {
		return ErrTooShortToDecode
	}

	var err error
	r.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	r.R1GenerationCounter = binary.BigEndian.Uint64(r.Header.Contents[4:12])

	return nil
}

// Serialize serializes a R1Counter into bytes.
func (r *R1Counter) Serialize() ([]byte, error) {
	b := make([]byte, r.Len())
	if err := r.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a R1Counter into bytes.
func (r *R1Counter) SerializeTo(b []byte) error {
	r.Header.Contents = make([]byte, 12)
	binary.BigEndian.PutUint64(r.Header.Contents[4:12], r.R1GenerationCounter)
	return r.Header.SerializeTo(b)
}

// Len returns the total length of a R1Counter, including Padding.
func (r *R1Counter) Len() int {
	return 4 + 12 // fixed by spec
}

// SetLength sets the length of Contents in Length field.
func (r *R1Counter) SetLength() {
	r.Length = 12 // fixed by spec
}

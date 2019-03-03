// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// HIPSignature represents a HIPSignature parameter.
//
// Spec: 5.2.14.  HIP_SIGNATURE
type HIPSignature struct {
	*Header
	SigAlg    uint16
	Signature []byte
}

// NewHIPSignature creates a new HIPSignature.
func NewHIPSignature(alg uint16, sig []byte) *HIPSignature {
	h := &HIPSignature{
		Header:    &Header{Type: ParamTypeHIPSignature},
		SigAlg:    alg,
		Signature: sig,
	}

	h.Padding = make([]byte, padlen(2+len(sig)))
	h.SetLength()
	return h
}

// DecodeHIPSignature decodes the given bytes as a HIPSignature.
func DecodeHIPSignature(b []byte) (*HIPSignature, error) {
	h := &HIPSignature{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HIPSignature.
func (h *HIPSignature) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 6 {
		return ErrTooShortToDecode
	}

	var err error
	h.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	h.SigAlg = binary.BigEndian.Uint16(h.Header.Contents[0:2])
	h.Signature = h.Header.Contents[2:]

	return nil
}

// Serialize serializes a HIPSignature into bytes.
func (h *HIPSignature) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HIPSignature into bytes.
func (h *HIPSignature) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	binary.BigEndian.PutUint16(h.Header.Contents[0:2], h.SigAlg)
	copy(h.Header.Contents[2:], h.Signature)

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HIPSignature, including Padding.
func (h *HIPSignature) Len() int {
	return 4 + 2 + len(h.Signature) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HIPSignature) SetLength() {
	h.Length = uint16(2 + len(h.Signature))
}

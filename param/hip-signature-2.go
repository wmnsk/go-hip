// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// HIPSignature2 represents a HIPSignature2 parameter.
//
// Spec: 5.2.15.  HIP_SIGNATURE_2
type HIPSignature2 struct {
	*Header
	SigAlg    uint16
	Signature []byte
}

// NewHIPSignature2 creates a new HIPSignature2.
func NewHIPSignature2(alg uint16, sig []byte) *HIPSignature2 {
	h := &HIPSignature2{
		Header:    &Header{Type: ParamTypeHIPSignature2},
		SigAlg:    alg,
		Signature: sig,
	}

	h.Padding = make([]byte, padlen(2+len(sig)))
	h.SetLength()
	return h
}

// DecodeHIPSignature2 decodes the given bytes as a HIPSignature2.
func DecodeHIPSignature2(b []byte) (*HIPSignature2, error) {
	h := &HIPSignature2{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HIPSignature2.
func (h *HIPSignature2) DecodeFromBytes(b []byte) error {
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

// Serialize serializes a HIPSignature2 into bytes.
func (h *HIPSignature2) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HIPSignature2 into bytes.
func (h *HIPSignature2) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	binary.BigEndian.PutUint16(h.Header.Contents[0:2], h.SigAlg)
	copy(h.Header.Contents[2:], h.Signature)

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HIPSignature2, including Padding.
func (h *HIPSignature2) Len() int {
	return 4 + 2 + len(h.Signature) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HIPSignature2) SetLength() {
	h.Length = uint16(2 + len(h.Signature))
}

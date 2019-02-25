// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// HIPMAC represents a HIPMAC parameter.
//
// Spec: 5.2.12.  HIP_MAC
type HIPMAC struct {
	*Header
	HMAC []byte
}

// NewHIPMAC creates a new HIPMAC.
func NewHIPMAC(hmac []byte) *HIPMAC {
	h := &HIPMAC{
		Header: &Header{Type: ParamTypeHIPMAC},
		HMAC:   hmac,
	}

	h.Padding = make([]byte, padlen(len(hmac)))
	h.SetLength()
	return h
}

// DecodeHIPMAC decodes the given bytes as a HIPMAC.
func DecodeHIPMAC(b []byte) (*HIPMAC, error) {
	h := &HIPMAC{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HIPMAC.
func (h *HIPMAC) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 4 {
		return ErrTooShortToDecode
	}

	var err error
	h.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	h.HMAC = h.Header.Contents

	return nil
}

// Serialize serializes a HIPMAC into bytes.
func (h *HIPMAC) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HIPMAC into bytes.
func (h *HIPMAC) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	copy(h.Header.Contents, h.HMAC)

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HIPMAC, including Padding.
func (h *HIPMAC) Len() int {
	return 4 + len(h.HMAC) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HIPMAC) SetLength() {
	h.Length = uint16(len(h.HMAC))
}

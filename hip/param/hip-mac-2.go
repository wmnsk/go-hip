// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// HIPMAC2 represents a HIPMAC2 parameter.
//
// Spec: 5.2.13.  HIP_MAC_2
type HIPMAC2 struct {
	*Header
	HMAC []byte
}

// NewHIPMAC2 creates a new HIPMAC2.
func NewHIPMAC2(hmac []byte) *HIPMAC2 {
	h := &HIPMAC2{
		Header: &Header{Type: ParamTypeHIPMAC2},
		HMAC:   hmac,
	}

	h.Padding = make([]byte, padlen(len(hmac)))
	h.SetLength()
	return h
}

// DecodeHIPMAC2 decodes the given bytes as a HIPMAC2.
func DecodeHIPMAC2(b []byte) (*HIPMAC2, error) {
	h := &HIPMAC2{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HIPMAC2.
func (h *HIPMAC2) DecodeFromBytes(b []byte) error {
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

// Serialize serializes a HIPMAC2 into bytes.
func (h *HIPMAC2) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HIPMAC2 into bytes.
func (h *HIPMAC2) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	copy(h.Header.Contents, h.HMAC)

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HIPMAC2, including Padding.
func (h *HIPMAC2) Len() int {
	return 4 + len(h.HMAC) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HIPMAC2) SetLength() {
	h.Length = uint16(len(h.HMAC))
}

// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import "encoding/binary"

// HIPCipher represents a HIPCipher parameter.
//
// Spec: 5.2.6.  DH_GROUP_LIST
type HIPCipher struct {
	*Header
	CipherIDs []uint16
}

// NewHIPCipher creates a new HIPCipher.
func NewHIPCipher(ids ...uint16) *HIPCipher {
	h := &HIPCipher{
		Header:    &Header{Type: ParamTypeHIPCipher},
		CipherIDs: ids,
	}

	h.Padding = make([]byte, padlen(len(ids)*2))
	h.SetLength()
	return h
}

// DecodeHIPCipher decodes the given bytes as a HIPCipher.
func DecodeHIPCipher(b []byte) (*HIPCipher, error) {
	h := &HIPCipher{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HIPCipher.
func (h *HIPCipher) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToDecode
	}

	var err error
	h.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	offset := 0
	ll := len(h.Header.Contents)
	for {
		if offset >= ll {
			break
		}
		h.CipherIDs = append(h.CipherIDs, binary.BigEndian.Uint16(h.Header.Contents[offset:offset+2]))
		offset += 2
	}

	return nil
}

// Serialize serializes a HIPCipher into bytes.
func (h *HIPCipher) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HIPCipher into bytes.
func (h *HIPCipher) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	offset := 0
	for _, id := range h.CipherIDs {
		binary.BigEndian.PutUint16(h.Header.Contents[offset:offset+2], id)
		offset += 2
	}

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HIPCipher, including Padding.
func (h *HIPCipher) Len() int {
	return 4 + (len(h.CipherIDs) * 2) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HIPCipher) SetLength() {
	h.Length = uint16(len(h.CipherIDs) * 2)
}

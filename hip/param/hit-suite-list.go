// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// HITSuiteList represents a HITSuiteList parameter.
//
// Spec: 5.2.10.  HIT_SUITE_LIST
type HITSuiteList struct {
	*Header
	DHGroupIDs []uint8
}

// NewHITSuiteList creates a new HITSuiteList.
func NewHITSuiteList(ids ...uint8) *HITSuiteList {
	h := &HITSuiteList{
		Header:     &Header{Type: ParamTypeHITSuiteList},
		DHGroupIDs: ids,
	}

	h.Padding = make([]byte, padlen(len(ids)))
	h.SetLength()
	return h
}

// DecodeHITSuiteList decodes the given bytes as a HITSuiteList.
func DecodeHITSuiteList(b []byte) (*HITSuiteList, error) {
	h := &HITSuiteList{}
	if err := h.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return h, nil
}

// DecodeFromBytes decodes the given bytes as a HITSuiteList.
func (h *HITSuiteList) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToDecode
	}

	var err error
	h.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	h.DHGroupIDs = h.Contents

	return nil
}

// Serialize serializes a HITSuiteList into bytes.
func (h *HITSuiteList) Serialize() ([]byte, error) {
	b := make([]byte, h.Len())
	if err := h.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a HITSuiteList into bytes.
func (h *HITSuiteList) SerializeTo(b []byte) error {
	h.Header.Contents = make([]byte, h.Len()-4)
	copy(h.Header.Contents, h.DHGroupIDs)

	return h.Header.SerializeTo(b)
}

// Len returns the total length of a HITSuiteList, including Padding.
func (h *HITSuiteList) Len() int {
	return 4 + len(h.DHGroupIDs) + len(h.Padding)
}

// SetLength sets the length of Contents in Length field.
func (h *HITSuiteList) SetLength() {
	h.Length = uint16(len(h.DHGroupIDs))
}

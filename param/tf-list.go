// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import "encoding/binary"

// TransportFormatList represents a TransportFormatList parameter.
//
// Spec: 5.2.11.  TRANSPORT_FORMAT_LIST
type TransportFormatList struct {
	*Header
	TFTypes []uint16
}

// NewTransportFormatList creates a new TransportFormatList.
func NewTransportFormatList(ids ...uint16) *TransportFormatList {
	t := &TransportFormatList{
		Header:  &Header{Type: ParamTypeTransportFormatList},
		TFTypes: ids,
	}

	t.Padding = make([]byte, padlen(len(ids)*2))
	t.SetLength()
	return t
}

// DecodeTransportFormatList decodes the given bytes as a TransportFormatList.
func DecodeTransportFormatList(b []byte) (*TransportFormatList, error) {
	t := &TransportFormatList{}
	if err := t.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return t, nil
}

// DecodeFromBytes decodes the given bytes as a TransportFormatList.
func (t *TransportFormatList) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToDecode
	}

	var err error
	t.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	offset := 0
	ll := len(t.Header.Contents)
	for {
		if offset >= ll {
			break
		}
		t.TFTypes = append(t.TFTypes, binary.BigEndian.Uint16(t.Header.Contents[offset:offset+2]))
		offset += 2
	}

	return nil
}

// Serialize serializes a TransportFormatList into bytes.
func (t *TransportFormatList) Serialize() ([]byte, error) {
	b := make([]byte, t.Len())
	if err := t.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a TransportFormatList into bytes.
func (t *TransportFormatList) SerializeTo(b []byte) error {
	t.Header.Contents = make([]byte, t.Len()-4)
	offset := 0
	for _, id := range t.TFTypes {
		binary.BigEndian.PutUint16(t.Header.Contents[offset:offset+2], id)
		offset += 2
	}

	return t.Header.SerializeTo(b)
}

// Len returns the total length of a TransportFormatList, including Padding.
func (t *TransportFormatList) Len() int {
	return 4 + (len(t.TFTypes) * 2) + len(t.Padding)
}

// SetLength sets the length of Contents in Length field.
func (t *TransportFormatList) SetLength() {
	t.Length = uint16(len(t.TFTypes) * 2)
}

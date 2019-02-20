// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// DHGroupList represents a header (common fields) in HIP parameter.
//
// Spec: 5.2.6.  DH_GROUP_LIST
type DHGroupList struct {
	*Header
	DHGroupIDs []uint8
}

// NewDHGroupList creates a new DHGroupList.
func NewDHGroupList(ids ...uint8) *DHGroupList {
	d := &DHGroupList{
		Header:     &Header{Type: ParamTypeDHGroupList},
		DHGroupIDs: ids,
	}

	d.Padding = make([]byte, padlen(len(ids)))
	d.SetLength()
	return d
}

// DecodeDHGroupList decodes the given bytes as a DHGroupList.
func DecodeDHGroupList(b []byte) (*DHGroupList, error) {
	d := &DHGroupList{}
	if err := d.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return d, nil
}

// DecodeFromBytes decodes the given bytes as a DHGroupList.
func (d *DHGroupList) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 5 {
		return ErrTooShortToDecode
	}

	var err error
	d.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	d.DHGroupIDs = d.Contents

	return nil
}

// Serialize serializes a DHGroupList into bytes.
func (d *DHGroupList) Serialize() ([]byte, error) {
	b := make([]byte, d.Len())
	if err := d.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a DHGroupList into bytes.
func (d *DHGroupList) SerializeTo(b []byte) error {
	d.Header.Contents = make([]byte, d.Len()-4)
	copy(d.Header.Contents, d.DHGroupIDs)

	return d.Header.SerializeTo(b)
}

// Len returns the total length of a DHGroupList, including Padding.
func (d *DHGroupList) Len() int {
	return 4 + len(d.DHGroupIDs) + len(d.Padding)
}

// SetLength sets the length of Contents in Length field.
func (d *DHGroupList) SetLength() {
	d.Length = uint16(len(d.DHGroupIDs))
}

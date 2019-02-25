// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import (
	"encoding/binary"
)

// Notification represents a Notification parameter.
//
// Spec: 5.2.19.  NOTIFICATION
type Notification struct {
	*Header
	Reserved          uint16
	NotifyMessageType uint16
	NotificationData  []byte
}

// NewNotification creates a new Notification.
func NewNotification(nType uint16, data []byte) *Notification {
	n := &Notification{
		Header:            &Header{Type: ParamTypeNotification},
		NotifyMessageType: nType,
		NotificationData:  data,
	}

	n.Padding = make([]byte, padlen(4+len(data)))
	n.SetLength()
	return n
}

// DecodeNotification decodes the given bytes as a Notification.
func DecodeNotification(b []byte) (*Notification, error) {
	n := &Notification{}
	if err := n.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return n, nil
}

// DecodeFromBytes decodes the given bytes as a Notification.
func (n *Notification) DecodeFromBytes(b []byte) error {
	l := len(b)
	if l < 9 {
		return ErrTooShortToDecode
	}

	var err error
	n.Header, err = DecodeHeader(b)
	if err != nil {
		return err
	}

	n.Reserved = binary.BigEndian.Uint16(n.Header.Contents[0:2])
	n.NotifyMessageType = binary.BigEndian.Uint16(n.Header.Contents[2:4])
	n.NotificationData = n.Header.Contents[4:]

	return nil
}

// Serialize serializes a Notification into bytes.
func (n *Notification) Serialize() ([]byte, error) {
	b := make([]byte, n.Len())
	if err := n.SerializeTo(b); err != nil {
		return nil, err
	}
	return b, nil
}

// SerializeTo serializes a Notification into bytes.
func (n *Notification) SerializeTo(b []byte) error {
	n.Header.Contents = make([]byte, n.Len()-4)
	binary.BigEndian.PutUint16(n.Header.Contents[0:2], n.Reserved)
	binary.BigEndian.PutUint16(n.Header.Contents[2:4], n.NotifyMessageType)
	copy(n.Header.Contents[4:], n.NotificationData)

	return n.Header.SerializeTo(b)
}

// Len returns the total length of a Notification, including Padding.
func (n *Notification) Len() int {
	return 4 + 4 + len(n.NotificationData) + len(n.Padding)
}

// SetLength sets the length of Contents in Length field.
func (n *Notification) SetLength() {
	n.Length = uint16(4 + len(n.NotificationData))
}

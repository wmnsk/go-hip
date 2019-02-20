// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package packet

// Packet Type definitions.
const (
	MsgTypeI1       uint8 = 1
	MsgTypeR1       uint8 = 2
	MsgTypeI2       uint8 = 3
	MsgTypeR2       uint8 = 4
	MsgTypeUpdate   uint8 = 16
	MsgTypeNotify   uint8 = 17
	MsgTypeClose    uint8 = 18
	MsgTypeCloseAck uint8 = 19
)

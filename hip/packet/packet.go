// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package packet

// Packet Type definitions.
const (
	PktTypeI1       uint8 = 1
	PktTypeR1       uint8 = 2
	PktTypeI2       uint8 = 3
	PktTypeR2       uint8 = 4
	PktTypeUpdate   uint8 = 16
	PktTypeNotify   uint8 = 17
	PktTypeClose    uint8 = 18
	PktTypeCloseAck uint8 = 19
)

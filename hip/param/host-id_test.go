// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/hip"
	"github.com/wmnsk/go-hip/hip/param"
	"github.com/wmnsk/go-hip/hip/param/testutils"
)

func TestHostID(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewHostID(
				hip.AlgRSA, hip.DomainFQDN,
				[]byte{0x11, 0x22, 0x33},
				[]byte{0x44, 0x55, 0x66},
			),
			Serialized: []byte{
				0x02, 0xc1, 0x00, 0x0c, 0x00, 0x03, 0x10, 0x03,
				0x00, 0x05, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewHostID(
				hip.AlgRSA, hip.DomainFQDN,
				[]byte{0xde, 0xad, 0xbe, 0xef},
				[]byte{0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0x02, 0xc1, 0x00, 0x0e, 0x00, 0x04, 0x10, 0x04,
				0x00, 0x05, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad,
				0xbe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeHostID(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

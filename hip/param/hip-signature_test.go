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

func TestHIPSignature(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewHIPSignature(
				hip.AlgRSA, []byte{0xca, 0xfe},
			),
			Serialized: []byte{
				0xf1, 0x01, 0x00, 0x04, 0x00, 0x05, 0xca, 0xfe,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewHIPSignature(
				hip.AlgRSA, []byte{0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0xf1, 0x01, 0x00, 0x06, 0x00, 0x05, 0xde, 0xad,
				0xbe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeHIPSignature(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

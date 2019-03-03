// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/param"
	"github.com/wmnsk/go-hip/param/testutils"
)

func TestDiffieHellman(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewDiffieHellman(
				1, []byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0x0ef, 0xff},
			),
			Serialized: []byte{
				0x02, 0x01, 0x00, 0x0c, 0x01, 0x00, 0x09,
				0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xff,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewDiffieHellman(
				1, []byte{0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0x02, 0x01, 0x00, 0x07, 0x01, 0x00, 0x04,
				0xde, 0xad, 0xbe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeDiffieHellman(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

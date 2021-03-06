// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/param"
	"github.com/wmnsk/go-hip/param/testutils"
)

func TestSolution(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewSolution(
				4, 0xffff,
				[]byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef},
				[]byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0x01, 0x41, 0x00, 0x14, 0x04, 0x00, 0xff, 0xff,
				0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
				0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewSolution(
				4, 0xffff,
				[]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe},
				[]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe},
			),
			Serialized: []byte{
				0x01, 0x41, 0x00, 0x10, 0x04, 0x00, 0xff, 0xff,
				0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe,
				0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe,
				0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeSolution(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

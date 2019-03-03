// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/param"
	"github.com/wmnsk/go-hip/param/testutils"
)

func TestPuzzle(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewPuzzle(
				4, 32, 0xffff,
				[]byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0x01, 0x01, 0x00, 0x0c, 0x04, 0x20, 0xff, 0xff,
				0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewPuzzle(
				4, 32, 0xffff,
				[]byte{0xde, 0xad, 0xbe, 0xef},
			),
			Serialized: []byte{
				0x01, 0x01, 0x00, 0x08, 0x04, 0x20, 0xff, 0xff,
				0xde, 0xad, 0xbe, 0xef, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodePuzzle(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/param"
	"github.com/wmnsk/go-hip/param/testutils"
)

func TestEchoResponseUnsigned(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  param.NewEchoResponseUnsigned([]byte{0xde, 0xad, 0xbe, 0xef}),
			Serialized: []byte{
				0xf7, 0xc1, 0x00, 0x04, 0xde, 0xad, 0xbe, 0xef,
			},
		}, {
			Description: "WithPadding",
			Structured:  param.NewEchoResponseUnsigned([]byte{0xca, 0xfe}),
			Serialized: []byte{
				0xf7, 0xc1, 0x00, 0x02, 0xca, 0xfe, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeEchoResponseUnsigned(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

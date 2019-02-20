// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/hip/param"
	"github.com/wmnsk/go-hip/hip/param/testutils"
)

func TestR1Counter(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  param.NewR1Counter(0x1122334455667788),
			Serialized: []byte{
				0x00, 0x81, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00,
				0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeR1Counter(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

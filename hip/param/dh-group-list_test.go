// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip/hip/param"
	"github.com/wmnsk/go-hip/hip/param/testutils"
)

func TestDHGroupList(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured:  param.NewDHGroupList(1, 2, 3, 4),
			Serialized: []byte{
				0x01, 0xff, 0x00, 0x04, 0x01, 0x02, 0x03, 0x04,
			},
		}, {
			Description: "WithPadding",
			Structured:  param.NewDHGroupList(1, 2, 3, 4, 5, 6),
			Serialized: []byte{
				0x01, 0xff, 0x00, 0x06, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeDHGroupList(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

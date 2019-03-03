// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param_test

import (
	"testing"

	"github.com/wmnsk/go-hip"

	"github.com/wmnsk/go-hip/param"
	"github.com/wmnsk/go-hip/param/testutils"
)

func TestHITSuiteList(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal",
			Structured: param.NewHITSuiteList(
				0, hip.HITSuiteRSADSASHA256, hip.HITSuiteECDSASHA384, hip.HITSuiteECDSALOWSHA1,
			),
			Serialized: []byte{
				0x02, 0xcb, 0x00, 0x04, 0x00, 0x10, 0x20, 0x30,
			},
		}, {
			Description: "WithPadding",
			Structured: param.NewHITSuiteList(
				hip.HITSuiteRSADSASHA256, hip.HITSuiteECDSASHA384, hip.HITSuiteECDSALOWSHA1,
			),
			Serialized: []byte{
				0x02, 0xcb, 0x00, 0x03, 0x10, 0x20, 0x30, 0x00,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Serializeable, error) {
		v, err := param.DecodeHITSuiteList(b)
		if err != nil {
			return nil, err
		}
		v.Contents = nil

		return v, nil
	})
}

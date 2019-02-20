// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

// Parameter Type definitions.
const (
	ParamTypeR1Counter            uint16 = 129
	ParamTypePuzzle               uint16 = 257
	ParamTypeSolution             uint16 = 321
	ParamTypeSeq                  uint16 = 385
	ParamTypeAck                  uint16 = 449
	ParamTypeDHGroupList          uint16 = 511
	ParamTypeDiffieHellman        uint16 = 513
	ParamTypeHIPCipher            uint16 = 579
	ParamTypeEncrypted            uint16 = 641
	ParamTypeHostID               uint16 = 705
	ParamTypeHITSuiteList         uint16 = 715
	ParamTypeCert                 uint16 = 768
	ParamTypeNotification         uint16 = 832
	ParamTypeEchoRequestSigned    uint16 = 897
	ParamTypeEchoResponseSigned   uint16 = 961
	ParamTypeTransportFormatList  uint16 = 2049
	ParamTypeHIPMAC               uint16 = 61505
	ParamTypeHIPMAC2              uint16 = 61569
	ParamTypeHIPSignature2        uint16 = 61633
	ParamTypeHIPSignature         uint16 = 61697
	ParamTypeEchoRequestUnSigned  uint16 = 63661
	ParamTypeEchoResponseUNSigned uint16 = 63425
)

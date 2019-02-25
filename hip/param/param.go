// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package param

import "encoding/binary"

// Parameter Type definitions.
//
// Spec: 5.2.  HIP Parameters
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
	ParamTypeEchoRequestUnsigned  uint16 = 63661
	ParamTypeEchoResponseUnsigned uint16 = 63425
)

// Param is an interface that all the parameters
type Param interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
	Len() int
	ParamType() uint16
}

// Decode decodes given bytes as Param.
func Decode(b []byte) (Param, error) {
	var p Param
	typ := binary.BigEndian.Uint16(b[:2])
	switch typ {
	case ParamTypeR1Counter:
		p = &R1Counter{}
	case ParamTypePuzzle:
		p = &Puzzle{}
	case ParamTypeSolution:
		p = &Solution{}
	case ParamTypeSeq:
		p = &Seq{}
	/* XXX - not implemented
	case ParamTypeAck:
		p = &Ack{}
	*/
	case ParamTypeDHGroupList:
		p = &DHGroupList{}
	case ParamTypeDiffieHellman:
		p = &DiffieHellman{}
	case ParamTypeHIPCipher:
		p = &HIPCipher{}
	/* XXX - not implemented
	case ParamTypeEncrypted:
		p = &Encrypted{}
	*/
	case ParamTypeHostID:
		p = &HostID{}
	case ParamTypeHITSuiteList:
		p = &HITSuiteList{}
	/* XXX - not implemented
	case ParamTypeCert:
		p = &Cert{}
	case ParamTypeNotification:
		p = &Notification{}
	case ParamTypeEchoRequestSigned:
		p = &EchoRequestSigned{}
	case ParamTypeEchoResponseSigned:
		p = &EchoResponseSigned{}
	*/
	case ParamTypeTransportFormatList:
		p = &TransportFormatList{}
	case ParamTypeHIPMAC:
		p = &HIPMAC{}
	case ParamTypeHIPMAC2:
		p = &HIPMAC2{}
	case ParamTypeHIPSignature2:
		p = &HIPSignature2{}
	case ParamTypeHIPSignature:
		p = &HIPSignature{}
		/* XXX - not implemented
		case ParamTypeEchoRequestUnSigned:
			p = &EchoRequestUnsigned{}
		case ParamTypeEchoResponseUNSigned:
			p = &EchoResponseUnsigned{}
		default:
			p = &Generic{}
		*/
	}

	if err := p.DecodeFromBytes(b); err != nil {
		return nil, err
	}
	return p, nil
}

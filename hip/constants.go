// Copyright 2019 go-hip authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package hip

// DH Group ID definitions.
//
// Spec: 5.2.7.  DIFFIE_HELLMAN
const (
	_ uint8 = iota
	_
	_
	DHGroupModPGroup1536Bit
	DHGroupModPGroup3072Bit
	_
	_
	DHGroupNISTP256
	DHGroupNISTP384
	DHGroupNISTP521
	DHGroupSecP160R1
	DHGroupModPGroup2048Bit
)

// Suite ID definitions.
//
// Spec: 5.2.8.  HIP_CIPHER
const (
	_ uint16 = iota
	CipherSuiteNullEncrypt
	CipherSuiteAES128CBC
	_
	CipherSuiteAES256CBC
)

// Domain Identifier definitions.
//
// Spec: 5.2.9.  HOST_ID
const (
	DomainNone uint8 = iota
	DomainFQDN
	DomainNAI
)

// HI Algorithm definitions.
//
// Spec: 5.2.9.  HOST_ID
const (
	_ uint16 = iota
	_
	_
	AlgDSA
	_
	AlgRSA
	_
	AlgECDSA
	_
	AlgECDSALow
)

// ECC Curve definitions for ECDSA algorithm.
//
// Spec: 5.2.9.  HOST_ID
const (
	_ uint16 = iota
	ECCNISTP256
	ECCNISTP384
)

// ECC Curve definitions for ECDSA_LOW algorithm.
//
// Spec: 5.2.9.  HOST_ID
const (
	_ uint16 = iota
	ECCSecP160R1
)

// HIT Suite definitions.
//
// Spec: 5.2.10.  HIT_SUITE_LIST
const (
	_                    uint8 = 0x00
	HITSuiteRSADSASHA256 uint8 = 0x10
	HITSuiteECDSASHA384  uint8 = 0x20
	HITSuiteECDSALOWSHA1 uint8 = 0x30
)

// Transport Format definitions.
//
// Spec: 5.2.11.  TRANSPORT_FORMAT_LIST, RFC 7401 and
// 5.1.  New Parameters, RFC 7402.
const (
	TFESP uint16 = 4095
)

// Notify Message Type definitions.
//
// Spec: 5.2.19.  NOTIFICATION
const (
	NotifyUnsupportedCriticalParameterType uint16 = 1
	NotifyInvalidSyntax                    uint16 = 7
	NotifyNoDHProposalChosen               uint16 = 14
	NotifyInvalidDHChosen                  uint16 = 15
	NotifyNoHIPProposalChosen              uint16 = 16
	NotifyInvalidHIPCipherChosen           uint16 = 17
	NotifyUnsupportedHITSuite              uint16 = 20
	NotifyAuthenticationFailed             uint16 = 24
	NotifyChecksumFailed                   uint16 = 26
	NotifyHIPMACFailed                     uint16 = 28
	NotifyEncryptionFailed                 uint16 = 32
	NotifyInvalidHIT                       uint16 = 40
	NotifyBlockedByPolicy                  uint16 = 42
	NotifyResponderBusyPleaseRetry         uint16 = 44
	NotifyI2Acknowledgement                uint16 = 16384
)

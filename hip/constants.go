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

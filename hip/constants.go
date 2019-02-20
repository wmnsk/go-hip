package hip

// DH Group ID definitions.
//
// Spec: 5.2.7.  DIFFIE_HELLMAN
const (
	_ uint8 = iota
	_
	_
	ModPGroup1536Bit
	ModPGroup3072Bit
	_
	_
	NISTP256
	NISTP384
	NISTP521
	SecP160R1
	ModPGroup2048Bit
)

// Suite ID definitions.
//
// Spec: 5.2.8.  HIP_CIPHER
const (
	_ uint16 = iota
	NullEncrypt
	AES128CBC
	_
	AES256CBC
)

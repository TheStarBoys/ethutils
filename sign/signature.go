package sign

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Signature [65]byte

func HexToSignature(str string) Signature {
	var sig Signature
	bts := common.Hex2Bytes(str)
	copy(sig[:], bts[0:65])
	return sig
}

func BytesToSignature(sig []byte) Signature {
	var res Signature
	copy(res[:], sig[:65])
	return res
}

func (s Signature) Hex() string {
	return fmt.Sprintf("0x%x", s)
}

func (s Signature) Bytes() []byte {
	return s[:]
}

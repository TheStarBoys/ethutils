package ethutils

import (
	"crypto/ecdsa"

	"github.com/TheStarBoys/ethutils/common"
	"github.com/TheStarBoys/ethutils/sign"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func PersonalSign(raw []byte, priv *ecdsa.PrivateKey) (common.Signature, error) {
	return sign.PersonalSign(raw, priv)
}

func EcRecover(data []byte, sig common.Signature) (ethcommon.Address, error) {
	return sign.EcRecover(data, sig)
}

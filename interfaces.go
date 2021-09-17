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

func PersonalSignHash(hash ethcommon.Hash, priv *ecdsa.PrivateKey) (common.Signature, error) {
	return sign.PersonalSignHash(hash, priv)
}

func EcRecover(data []byte, sig common.Signature) (ethcommon.Address, error) {
	return sign.EcRecover(data, sig)
}

func EcRecoverHash(hash ethcommon.Hash, sig common.Signature) (ethcommon.Address, error) {
	return sign.EcRecoverHash(hash, sig)
}

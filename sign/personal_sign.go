package sign

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/TheStarBoys/ethutils/common"
	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PersonalSign(raw []byte, priv *ecdsa.PrivateKey) (common.Signature, error) {
	hash := accounts.TextHash(raw)
	return PersonalSignHash(ethcommon.BytesToHash(hash), priv)
}

func PersonalSignHash(hash ethcommon.Hash, priv *ecdsa.PrivateKey) (common.Signature, error) {
	sig, err := crypto.Sign(hash.Bytes(), priv)
	if err != nil {
		return common.Signature{}, err
	}

	sig[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 0/1 to 27/28
	return common.BytesToSignature(sig), nil
}

func EcRecover(data []byte, sig common.Signature) (ethcommon.Address, error) {
	hash := accounts.TextHash(data)
	return EcRecoverHash(ethcommon.BytesToHash(hash), sig)
}

func EcRecoverHash(hash ethcommon.Hash, sig common.Signature) (ethcommon.Address, error) {
	if len(sig) != crypto.SignatureLength {
		return ethcommon.Address{}, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}

	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return ethcommon.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}

	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	rpk, err := crypto.SigToPub(hash.Bytes(), sig.Bytes())
	if err != nil {
		return ethcommon.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}

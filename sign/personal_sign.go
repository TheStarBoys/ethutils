package sign

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PersonalSign(raw []byte, priv *ecdsa.PrivateKey) (Signature, error) {
	hash := accounts.TextHash(raw)
	sig, err := crypto.Sign(hash, priv)
	if err != nil {
		return Signature{}, err
	}

	sig[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 0/1 to 27/28
	return BytesToSignature(sig), nil
}

func EcRecover(data []byte, sig Signature) (common.Address, error) {
	if len(sig) != crypto.SignatureLength {
		return common.Address{}, fmt.Errorf("signature must be %d bytes long", crypto.SignatureLength)
	}

	if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}

	sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

	rpk, err := crypto.SigToPub(accounts.TextHash(data), sig.Bytes())
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}
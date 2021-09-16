package sign

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestPersonalSign(t *testing.T) {
	var tests = []struct {
		privHex   string
		rawText   string
		expectSig string
	}{
		{
			"9a01f5c57e377e0239e6036b7b2d700454b760b2dab51390f1eeb2f64fe98b68",
			"hello",
			"0xa6fb700f874d5a293ce7ca1f7850cc2fcff5a5e79e65ee901428802c1bf9762e01395ce53844bd9b9dfb79d96bb9a558d257d21171f958a00d573d970a2387391c",
		},
		{
			"9a01f5c57e377e0239e6036b7b2d700454b760b2dab51390f1eeb2f64fe98b68",
			"hello, world",
			"0xb86c6f20f5184bdc482a5e8fd32585aba2be021221af8831b893d8186957c58e766692a7d6b10ab8e25b165b6def0550515a087523d3e0c5dda2892ec77a8ed61c",
		},
	}

	for _, tt := range tests {
		priv, err := crypto.HexToECDSA(tt.privHex)
		if err != nil {
			t.Fatal(err)
		}

		sig, err := PersonalSign([]byte(tt.rawText), priv)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tt.expectSig, sig.Hex())

		gotAddr, err := EcRecover([]byte(tt.rawText), sig)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, crypto.PubkeyToAddress(priv.PublicKey).Hex(), gotAddr.Hex())
	}
}

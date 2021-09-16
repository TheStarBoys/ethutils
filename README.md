# ethutils
The collection of ethereum utils.
## Quick Start
```go
package main

import (
	"fmt"

	"github.com/TheStarBoys/ethutils"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	priv, err := crypto.HexToECDSA("9a01f5c57e377e0239e6036b7b2d700454b760b2dab51390f1eeb2f64fe98b68")
	if err != nil {
		panic(err)
	}

	rawText := "hello"
	sig, err := ethutils.PersonalSign([]byte(rawText), priv)
	if err != nil {
		panic(err)
	}

	fmt.Println("got sig:", sig.Hex())

	gotAddr, err := ethutils.EcRecover([]byte(rawText), sig)
	if err != nil {
		panic(err)
	}

	fmt.Println("got addr:", gotAddr.Hex())
}
```
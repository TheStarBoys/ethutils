package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type Wallet struct {
	Mnemonic string
	*hdwallet.Wallet
}

func main() {
	pattern := ".*"

	if len(os.Args) < 2 {
		log.Println("Without any pattern. If you want to add pattern, please use `command <pattern>`")
	} else {
		pattern = os.Args[1]
	}

	count := 0

	walletChan := make(chan *Wallet, 1000)
	go func() {
		for {
			// Generate new wallet
			mnemonic, err := hdwallet.NewMnemonic(128)
			if err != nil {
				log.Fatal(err)
			}

			wallet, err := hdwallet.NewFromMnemonic(mnemonic)
			if err != nil {
				log.Fatal(err)
			}

			walletChan <- &Wallet{
				Mnemonic: mnemonic,
				Wallet:   wallet,
			}
		}
	}()

GetWallet:
	for {
		wallet := <-walletChan
		count++
		if count%10000 == 0 {
			log.Println("Total retry count: ", count)
		}

		var match bool
		// Find the address you expected
		for i := 0; i < 10; i++ {
			path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", i))
			account, err := wallet.Derive(path, false)
			if err != nil {
				log.Fatal(err)
			}

			if i == 0 {
				match, err = isGoodNumber(pattern, account.Address.Hex())
				if err != nil {
					log.Fatal("Good number err: ", err)
				}

				if !match {
					continue GetWallet
				} else {
					log.Println("Match!")
				}
			}

			privateKey, err := wallet.PrivateKeyHex(account)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%d: address: %v, privateKey: %v", i, account.Address.Hex(), privateKey)
		}

		if match {
			log.Println("Mnemonic: ", wallet.Mnemonic)
			break
		}
	}

	log.Println("Total retry count: ", count)
}

func isGoodNumber(pattern, address string) (bool, error) {
	address = strings.ToLower(address)
	return regexp.Match(pattern, []byte(address))
}

#!/bin/bash

geth=../../cmd/geth
$geth account import --datadir data/ --password passwd.txt priv.key
$geth account import --datadir data/ --password passwd.txt priv2.key

$geth init --datadir data genesis.json

miner=0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a
# run
$geth --datadir data --networkid 62888 --allow-insecure-unlock --gcmode=archive --nodiscover \
	--unlock "$miner, 0x06514D014e997bcd4A9381bF0C4Dc21bD32718D4" \
	--password "passwd.txt" \
	--http.vhosts "*" --http.api eth,net,web3,debug,txpool \
	--http --http.addr 0.0.0.0 --http.corsdomain "http://localhost:8000" \
	--http.corsdomain "https://remix.ethereum.org,http://remix.ethereum.org,http://localhost:8000" \
	--graphql \
	--ws --ws.api "eth,net,web3,debug,txpool" --ws.addr 0.0.0.0 --ws.origins "*" \
	--mine --miner.etherbase $miner

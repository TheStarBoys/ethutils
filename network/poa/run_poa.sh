#!/bin/bash
geth account import --datadir data/ --password passwd.txt priv.key
geth account import --datadir data/ --password passwd.txt priv2.key

geth init --datadir data genesis.json

miner=0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a
# run
geth --datadir data --networkid 15 --rpc \
	--unlock "$miner, 0x06514D014e997bcd4A9381bF0C4Dc21bD32718D4" \
	--password "passwd.txt" \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
	--graphql --graphql.addr 0.0.0.0 --graphql.port 8547 \
	--ws --rpcapi eth,net,web3,personal,admin,debug,txpool \
        -wsaddr 0.0.0.0 -rpcaddr 0.0.0.0 \
	--allow-insecure-unlock \
	--mine --miner.etherbase $miner

#!/bin/bash
geth account import --datadir data/ --password passwd.txt priv.key
geth init --datadir data genesis.json

boot_key_path=bootnode.key

# run
geth --datadir data --networkid 15 --rpc \
	--unlock 0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a \
	--password passwd.txt \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
	--ws --rpcapi eth,net,web3,personal,admin,debug,miner,txpool \
        -wsaddr 0.0.0.0 -rpcaddr 0.0.0.0 --verbosity 3 \
	--port 30301 --rpcport 8645 --wsport 8646 \
	--nodekey $boot_key_path \
	--allow-insecure-unlock \
	--miner.gasprice 1000 \
	--mine

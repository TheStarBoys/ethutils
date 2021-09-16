#!/bin/bash

datadir=./data/boot1
boot_key_path=bootnode1.key
networkid=15
discvport=30301
rpcport=8645

geth init --datadir $datadir genesis.json

# run
geth --datadir $datadir --networkid $networkid \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
        --rpc -rpcaddr 0.0.0.0 --verbosity 3 \
	--port $discvport --rpcport $rpcport \
	--nodekey $boot_key_path \
	--allow-insecure-unlock \
	--miner.etherbase 0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a

#!/bin/bash

datadir=./data/boot2

boot1_ip=127.0.0.1
boot1_port=30301
boot1_key_path=bootnode1.key
nodeid1="enode://`bootnode -nodekey ${boot1_key_path} --writeaddress`@${boot1_ip}:${boot1_port}"

echo 'boot1 ' $nodeid1

boot_key_path=bootnode2.key
networkid=15
discvport=30302
rpcport=8745

geth init --datadir $datadir genesis.json

# run
geth --datadir $datadir --networkid $networkid \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
        --rpc -rpcaddr 0.0.0.0 --verbosity 3 \
	--port $discvport --rpcport $rpcport \
	--nodekey $boot_key_path --bootnodes $nodeid1 \
	--allow-insecure-unlock \
	--miner.etherbase 0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a

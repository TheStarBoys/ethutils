#!/bin/bash
datadir=./data/node1
networkid=15
discvport=30303
rpcport=8845

boot1_ip=127.0.0.1
boot1_port=30301
boot1_key_path=bootnode1.key
nodeid1="enode://`bootnode -nodekey ${boot1_key_path} --writeaddress`@${boot1_ip}:${boot1_port}"

boot2_ip=127.0.0.1
boot2_port=30302
boot2_key_path=bootnode2.key
nodeid2="enode://`bootnode -nodekey ${boot2_key_path} --writeaddress`@${boot2_ip}:${boot2_port}"

geth account import --datadir $datadir/ --password passwd.txt priv.key
geth init --datadir $datadir genesis.json

# run
geth --datadir $datadir --networkid $networkid --rpc \
	--unlock 0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a \
	--password passwd.txt \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
	--rpcapi eth,net,web3,personal,admin,debug,txpool \
        -wsaddr 0.0.0.0 -rpcaddr 0.0.0.0 --verbosity 3 \
	--port $discvport --rpcport $rpcport \
	--bootnodes "${nodeid1},${nodeid2}" \
	--miner.gasprice 1000 --mine \
	--allow-insecure-unlock

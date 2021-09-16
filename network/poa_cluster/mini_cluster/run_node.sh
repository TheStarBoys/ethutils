#!/bin/bash
geth account import --datadir node1/ --password passwd.txt priv.key
geth init --datadir node1 genesis.json

boot_ip=127.0.0.1
boot_port=30301
boot_key_path=bootnode.key
nodeid="enode://`bootnode -nodekey ${boot_key_path} --writeaddress`@${boot_ip}:${boot_port}"

# run
geth --datadir node1 --networkid 15 --rpc \
	--unlock 0x11a449eD8eadAcDA0A290ad0ffEE174fdF0B3f7a \
	--password passwd.txt \
	--rpccorsdomain "https://remix.ethereum.org,http://remix.ethereum.org" \
	--ws --rpcapi eth,net,web3,personal,admin,debug,txpool \
        -wsaddr 0.0.0.0 -rpcaddr 0.0.0.0 --verbosity 3 \
	--port 30302 --rpcport 8745 --wsport 8746 \
	--bootnodes ${nodeid} \
	--miner.gasprice 1000 \
	--allow-insecure-unlock

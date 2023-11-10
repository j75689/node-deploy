#!/bin/bash

i=$1
stateScheme="hash"
HTTPPort=8545
WSPort=${HTTPPort}
MetricsPort=6060

for j in /server/clusterNetwork/validator${i}/keystore/*; do
    cons_addr="0x$(cat ${j} | jq -r .address)"
done

nohup /server/clusterNetwork/node${i}/geth${i} --config /server/clusterNetwork/node${i}/config.toml \
    --datadir /server/clusterNetwork/node${i} \
    --password /server/clusterNetwork/password.txt \
    --blspassword /server/clusterNetwork/password.txt \
    --nodekey /server/clusterNetwork/node${i}/geth/nodekey \
    -unlock ${cons_addr} --miner.etherbase ${cons_addr} --rpc.allow-unprotected-txs --allow-insecure-unlock \
    --ws.addr 0.0.0.0 --ws.port ${WSPort} --http.addr 0.0.0.0 --http.port ${HTTPPort} --http.corsdomain "*" \
    --metrics --metrics.addr localhost --metrics.port ${MetricsPort} --metrics.expensive \
    --gcmode archive --state.scheme ${stateScheme} --syncmode full --mine --vote --monitor.maliciousvote \
    >/server/clusterNetwork/node${i}/bsc-node.log 2>&1 &

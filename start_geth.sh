#!/bin/bash

i=$1
cmd=$2
stateScheme="hash"
HTTPPort=8545
WSPort=${HTTPPort}
MetricsPort=6060

if [ "${cmd}" == "reset" ]; then
    rm -rf /server/bsc/validator
    mkdir -p /server/bsc/validator
    cp -f /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/password.txt /server/bsc/validator/
    cp -r /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/clusterNetwork/node${i}/* /server/bsc/validator/
fi

rm -rf /server/bsc/validator/bls
rm -rf /server/bsc/validator/keystore
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/geth /server/bsc/validator/geth${i}
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/clusterNetwork/node${i}/config.toml /server/bsc/validator/
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/clusterNetwork/node${i}/keystore /server/bsc/validator/
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion/bsc_cluster/clusterNetwork/node${i}/bls /server/bsc/validator/


for j in /server/bsc/validator/keystore/*; do
    cons_addr="0x$(cat ${j} | jq -r .address)"
done

nohup /server/bsc/validator/geth${i} --config /server/bsc/validator/config.toml \
    --datadir /server/bsc/validator/ \
    --password /server/bsc/validator/password.txt \
    --blspassword /server/bsc/validator/password.txt \
    --nodekey /server/bsc/validator/geth/nodekey \
    -unlock ${cons_addr} --miner.etherbase ${cons_addr} --rpc.allow-unprotected-txs --allow-insecure-unlock \
    --ws.addr 0.0.0.0 --ws.port ${WSPort} --http.addr 0.0.0.0 --http.port ${HTTPPort} --http.corsdomain "*" \
    --metrics --metrics.addr localhost --metrics.port ${MetricsPort} --metrics.expensive \
    --gcmode archive --state.scheme ${stateScheme} --syncmode full --mine --vote --monitor.maliciousvote \
    > /server/bsc/validator/bsc-node.log 2>&1 &

#!/bin/bash

i=$1
cmd=$2

if [ "${cmd}" == "reset" ]; then
    rm -rf /server/bc/node
    mkdir -p /server/bc/node
    cp -r /mnt/efs/bsc-qa/bc-fusion-gov-env/bc_cluster/node${i}/* /server/bc/node/
fi
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bc_cluster/bnbchaind /server/bc/node/
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bc_cluster/node${i}/config/app.toml /server/bc/node/config/
yes | cp -rf /mnt/efs/bsc-qa/bc-fusion-gov-env/bc_cluster/node${i}/config/config.toml /server/bc/node/config/
nohup /server/bc/node/bnbchaind start --home /server/bc/node > /server/bc/node/bc-node.log 2>&1 &
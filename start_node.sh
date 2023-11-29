#!/bin/bash

i=$1
mkdir -p /server/bc/node
rm -rf /server/bc/node
cp -f /mnt/efs/bsc-qa/bc-fusion/bc_cluster/bnbchaind /server/bc/node/
cp -r /mnt/efs/bsc-qa/bc-fusion/bc_cluster/node${i}/* /server/bc/node/

nohup /server/bc/node/bnbchaind start --home /server/bc/node > /server/bc/node/bc-node.log 2>&1 &
